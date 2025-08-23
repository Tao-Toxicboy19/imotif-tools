package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/imotif-tools/internal/cli"
	"github.com/imotif-tools/internal/git"
	"google.golang.org/api/option"
)

type GeminiProvider struct {
	ModelName string
	APIKey    string
}

func NewGeminiProvider(model, apiKey string) *GeminiProvider {
	return &GeminiProvider{
		ModelName: model,
		APIKey:    apiKey,
	}
}

func (g *GeminiProvider) RunCommand() error {
	ctx := context.Background()
	exec := git.NewGitExec()
	files, err := exec.GetStagedFilesWithContent()
	if err != nil {
		return err
	}
	msg, err := g.genMessage(ctx, files)
	if err != nil {
		return err
	}
	fmt.Println("Generated commit message:", msg)
	// Run commit prompt
	prompter := cli.NewCommitPrompter([]string{})
	// Verify commit message
	verifiedMsg, err := prompter.ConfirmOrEditMessage(msg)
	fmt.Println("Verified commit message:", verifiedMsg)
	if err != nil {
		return err
	}
	// Run commit
	if err := prompter.RunCommit(verifiedMsg); err != nil {
		return err
	}

	return nil
}

func (g *GeminiProvider) genMessage(ctx context.Context, files []git.FileWithContent) (string, error) {
	client, err := g.newClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	prompt := g.buildPrompt(files)

	resp, err := g.generate(ctx, client, prompt)
	if err != nil {
		return "", err
	}

	msg, err := g.extractContent(resp)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (g *GeminiProvider) newClient(ctx context.Context) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(g.APIKey))
	if err != nil {
		return nil, fmt.Errorf("failed to init Gemini client: %w", err)
	}
	return client, nil
}

func (g *GeminiProvider) buildPrompt(files []git.FileWithContent) string {
	var sb strings.Builder
	sb.WriteString("Summarize the following code changes into a Git commit message.\n")
	sb.WriteString("Return only the message, no prefix or issue number.\n")
	sb.WriteString("Limit to 10 words or fewer. Be clear and concise.\n\n")

	for _, f := range files {
		sb.WriteString(fmt.Sprintf("Filename: %s\n", f.Filename))
		sb.WriteString(fmt.Sprintf("Content:\n%s\n\n", f.Content))
	}
	return sb.String()
}

func (g *GeminiProvider) generate(ctx context.Context, client *genai.Client, prompt string) (*genai.GenerateContentResponse, error) {
	model := client.GenerativeModel(g.ModelName)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("failed to generate from Gemini: %w", err)
	}
	return resp, nil
}

func (g *GeminiProvider) extractContent(resp *genai.GenerateContentResponse) (string, error) {
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			if text, ok := part.(genai.Text); ok {
				trimmed := strings.TrimSpace(string(text))
				if trimmed == "" {
					return "", fmt.Errorf("Gemini returned empty commit message")
				}
				return trimmed, nil
			}
		}
	}
	return "", fmt.Errorf("no usable content returned from Gemini")
}
