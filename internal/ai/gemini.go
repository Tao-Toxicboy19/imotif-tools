package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
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

func (g *GeminiProvider) Suggest(ctx context.Context, files []git.FileWithContent) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(g.APIKey))
	if err != nil {
		return "", fmt.Errorf("failed to init Gemini client: %w", err)
	}
	defer client.Close()

	// Compose input prompt
	var sb strings.Builder
	sb.WriteString("You are an AI assistant that writes clean and concise Git commit messages.\n")
	sb.WriteString("Your task is to summarize the following file changes into a single-line commit message.\n")
	sb.WriteString("Do NOT include any prefix like 'feat:', 'fix:', or issue numbers.\n")
	sb.WriteString("Only return the commit message itself, nothing else.\n\n")

	for _, f := range files {
		sb.WriteString(fmt.Sprintf("Filename: %s\n", f.Filename))
		sb.WriteString(fmt.Sprintf("Content:\n%s\n\n", f.Content))
	}

	// Call Gemini model
	model := client.GenerativeModel(g.ModelName)
	resp, err := model.GenerateContent(ctx, genai.Text(sb.String()))
	if err != nil {
		return "", fmt.Errorf("failed to generate from Gemini: %w", err)
	}

	// Extract content
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
