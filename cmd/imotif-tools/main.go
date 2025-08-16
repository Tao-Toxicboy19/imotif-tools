package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/imotif-tools/internal/ai"
	"github.com/imotif-tools/internal/cli"
	"github.com/imotif-tools/internal/config"
	"github.com/imotif-tools/internal/git"
	"github.com/imotif-tools/internal/update"
)

const version = "v1.0.6"

func main() {
	args := os.Args

	if len(args) == 1 {
		// Init Banner
		banner := cli.NewBanner()
		fmt.Println("imotif-tools version:", version)
		banner.PrintBanner()
		return
	}

	switch args[1] {
	case "--version", "-v":
		fmt.Println("imotif-tools version:", version)
		return
	case "--help", "-h":
		// Init Banner
		banner := cli.NewBanner()
		banner.PrintHelp()
		return
	case "commit":
		{
			if len(args) > 2 {
				// Init Commit Prompter
				prompter := cli.NewCommitPrompter()
				msg := strings.Join(args[2:], " ")
				if strings.TrimSpace(msg) == "" {
					fmt.Println("Commit message is empty.")
					return
				}
				msg, err := prompter.Run(msg)
				if err != nil {
					log.Fatal("Failed to run commit prompt:", err)
				}
				fmt.Println(msg)
			} else {
				fmt.Println("Commit message is empty.")
			}
			return
		}
	case "init":
		{
			// Init AliasInitializer
			initializer := cli.NewAliasInitializer()
			initializer.Run()
			return
		}
	case "update":
		{
			// Init SelfUpdater
			selfUpdater, err := update.NewSelfUpdater()
			if err != nil {
				log.Fatal("Failed to create self updater:", err)
			}
			fmt.Println("Updating imotif-tools...")
			selfUpdater.Run()
			return
		}
	case "magic":
		{
			// Init GitExec
			exec := git.NewGitExec()
			// Get staged files with content
			files, err := exec.GetStagedFilesWithContent()
			if err != nil {
				log.Fatal("Failed to get staged files:", err)
			}

			if len(files) == 0 {
				log.Fatal("No staged files found.")
			}
			// Load ENV config (MODEL, API_KEY)
			cfg := config.Load()
			// Init AI provider (Gemini)
			provider := ai.NewGeminiProvider(cfg.Model, cfg.APIKey)
			fmt.Println("Generating commit message...")
			// Generate commit message from file content
			genMsg, err := provider.Suggest(context.Background(), files)
			if err != nil {
				log.Fatal("AI failed to generate commit message:", err)
			}
			fmt.Println("Generated commit message:", genMsg)
			// Run commit prompt
			prompter := cli.NewCommitPrompter()
			// Verify commit message
			verifiedMsg, err := prompter.ConfirmOrEditMessage(genMsg)
			if err != nil {
				log.Fatal("Failed to verify commit message:", err)
			}
			// Run commit
			_, err = prompter.Run(verifiedMsg)
			if err != nil {
				log.Fatal("Failed to run commit prompt:", err)
			}
			return
		}
	default:
		fmt.Println("Invalid command.")
		return
	}
}
