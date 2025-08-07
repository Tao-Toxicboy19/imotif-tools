package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

const version = "v1.0.0"

func main() {
	args := os.Args

	// --- Handle: --version
	if len(args) > 1 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Println("imotif-tools version:", version)
		return
	}

	// --- Handle: commit (explicit) or via alias (e.g., itgc)
	if len(args) == 1 || (len(args) > 1 && args[1] == "commit") {
		runCommitPrompt()
		return
	}

	// --- Unknown command
	fmt.Println("Unknown command:", args[1])
	fmt.Println("Usage:")
	fmt.Println("  imotif-tools commit        Run interactive commit prompt")
	fmt.Println("  imotif-tools --version     Show version")
	os.Exit(1)
}

func runCommitPrompt() {
	// Step 1: Task ID
	var taskID string
	err := survey.AskOne(&survey.Input{
		Message: "Enter Task NO.:",
	}, &taskID)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด:", err)
		return
	}

	// Step 2: Commit type
	typeMap := map[string]string{
		"FIX":   "Fix a bug or incorrect behavior",
		"ADD":   "Add new code, files, or features",
		"REF":   "Refactor code without changing behavior",
		"MIG":   "Migration related changes (e.g., DB schema)",
		"DOC":   "Documentation updates or changes",
		"STYLE": "Code style changes (no logic impact)",
		"TEST":  "Add or update tests",
		"CHORE": "Maintenance tasks (e.g., config, build)",
		"PERF":  "Performance improvements",
		"FEAT":  "Introduce a new feature",
	}

	fmt.Println("Available Commit Types:")
	for k, v := range typeMap {
		fmt.Printf("  [%s] %s\n", k, v)
	}
	fmt.Println()

	// Step 3: Commit type input
	var inputType string
	survey.AskOne(&survey.Input{
		Message: "Enter commit type (e.g. fix, add, ref):",
	}, &inputType)

	inputType = strings.TrimSpace(inputType)
	if inputType == "" {
		fmt.Println("You must enter a commit type.")
		return
	}
	commitType := strings.ToUpper(inputType)

	// Step 4: Commit message
	var commitMsg string
	err = survey.AskOne(&survey.Input{
		Message: "Enter commit message:",
	}, &commitMsg)
	if err != nil {
		fmt.Println(err)
		return
	}
	commitMsg = strings.TrimSpace(commitMsg)
	if commitMsg == "" {
		fmt.Println("You must enter a commit message.")
		return
	}

	// Step 5: Confirm verify
	var doVerify bool
	err = survey.AskOne(&survey.Confirm{
		Message: "Do you want to verify before committing?",
		Default: true,
	}, &doVerify)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 6: Commit message
	finalMessage := fmt.Sprintf("[%s] [%s] %s", taskID, commitType, commitMsg)

	// Step 7: Run git commit
	args := []string{"commit", "-m", finalMessage}
	if !doVerify {
		args = append(args, "--no-verify")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Commit failed:", err)
		return
	}

	fmt.Println("---------------------")
	fmt.Println("!!!Commit Success!!!")
	fmt.Println("---------------------")
}
