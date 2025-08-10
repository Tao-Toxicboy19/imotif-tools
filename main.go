package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	figure "github.com/common-nighthawk/go-figure"
)

const version = "v1.0.5"

func main() {
	args := os.Args

	// --- Handle: --version
	if len(args) > 1 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Println("imotif-tools version:", version)
		return
	}

	// --- Handle: commit command
	if len(args) > 1 && args[1] == "commit" {
		if len(args) > 2 {
			// ใช้ args หลังจาก commit เป็น commit message เลย
			commitMsg := strings.Join(args[2:], " ")
			runCommitPrompt(commitMsg)
			return
		}
        fmt.Println("Commit message is empty.")
		return
	}

	// --- Handle: init command
	if len(args) > 1 && args[1] == "init" {
		runInitAlias()
		return
	}

	// --- Handle: Update Version
	if len(args) > 1 && args[1] == "update" {
		fmt.Println("Updating imotif-tools...")
		runUpdateVersion()
		return
	}

	// --- Handle: --version
	if len(args) > 1 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Println("imotif-tools version:", version)
		printHelp()
		return
	}

	// --- Default (no args)
	if len(args) == 1 {
		myFigure := figure.NewFigure("IMOITF Tools", "", true)
		myFigure.Print()
		return
	}
}

func printHelp() {
	fmt.Println(`imotif-tools - interactive git commit helper

Usage:
  imotif-tools                 Show banner & usage
  imotif-tools commit          Run interactive commit prompt
  imotif-tools init            Add alias 'itcm' to your shell config
  imotif-tools update          Self-update to the latest release
  imotif-tools --version       Show version
  imotif-tools -h | --help     Show this help

Examples:
  imotif-tools commit
  imotif-tools init
  imotif-tools update
`)
}

func runCommitPrompt(msg string) {
	if strings.TrimSpace(msg) == "" {
        fmt.Println("Commit message is empty.")
		return
	}

	// Step 1: Task ID (support multiple)
	var taskInput string
	err := survey.AskOne(&survey.Input{
		Message: "Enter Task NO. (e.g. Task-1 or Task-1,Task-2,Task-3):",
	}, &taskInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	taskInput = strings.TrimSpace(taskInput)
	if taskInput == "" {
		fmt.Println("You must enter a Task NO.")
		return
	}

	// Split and format task IDs
	tasks := strings.Split(taskInput, ",")
	for i, t := range tasks {
		tasks[i] = strings.TrimSpace(t)
	}
	formattedTask := "[" + strings.Join(tasks, ",") + "]"

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
		"ISS":   "Work related to a reported issue or bug ticket",
	}

	fmt.Println("Available Commit Types:")
	for k, v := range typeMap {
		fmt.Printf("  [%s] %s\n", k, v)
	}

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
	finalMessage := fmt.Sprintf("%s [%s] %s", formattedTask, commitType, msg)

	// Step 5: Run git commit
	args := []string{"commit", "-m", finalMessage}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Commit failed:", err)
		return
	}
}

func runInitAlias() {
	var shellPath string
	prompt := &survey.Input{
		Message: "Enter the path to your shell config file (e.g. ~/.zshrc, ~/.bashrc, ~/.config/fish/config.fish):",
	}
	err := survey.AskOne(prompt, &shellPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	shellPath = strings.TrimSpace(shellPath)
	if shellPath == "" {
		fmt.Println("Shell path is required.")
		return
	}

	// Expand ~ to home dir
	if strings.HasPrefix(shellPath, "~") {
		home, _ := os.UserHomeDir()
		shellPath = filepath.Join(home, shellPath[1:])
	}

	// Prepare alias line
	var aliasLine string
	if strings.Contains(shellPath, "fish") {
		aliasLine = "alias itcm 'imotif-tools commit'"
	} else {
		aliasLine = "alias itcm='imotif-tools commit'"
	}

	// Check if alias already exists
	content, err := os.ReadFile(shellPath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	if strings.Contains(string(content), aliasLine) {
		fmt.Println("Alias already exists in", shellPath)
	} else {
		f, err := os.OpenFile(shellPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Cannot write to file:", err)
			return
		}
		defer f.Close()

		if _, err := f.WriteString("\n" + aliasLine + "\n"); err != nil {
			fmt.Println("Failed to write alias:", err)
			return
		}
		fmt.Println("Alias added to", shellPath)
	}

	fmt.Println("Please restart your terminal or run:")
	if strings.Contains(shellPath, "fish") {
		fmt.Println("source", shellPath)
	} else {
		fmt.Println("source", shellPath)
	}
}

func runUpdateVersion() {
	installPath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to determine current binary path:", err)
		return
	}

	OS := runtime.GOOS
	urlMap := map[string]string{
		"darwin":  "https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/imotif-tools-macos",
		"linux":   "https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/imotif-tools-linux",
		"windows": "https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/imotif-tools.exe",
	}

	binaryURL, ok := urlMap[OS]
	if !ok {
		fmt.Println("Unsupported OS:", OS)
		return
	}

	fmt.Println("Downloading latest version from:", binaryURL)

	// Download the file
	resp, err := http.Get(binaryURL)
	if err != nil {
		fmt.Println("Failed to download latest version:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error: received non-200 response:", resp.Status)
		return
	}

	// Write to temp file first
	tmpFile := installPath + ".tmp"
	out, err := os.Create(tmpFile)
	if err != nil {
		fmt.Println("Failed to create temp file:", err)
		return
	}
	_, err = io.Copy(out, resp.Body)
	out.Close()
	if err != nil {
		fmt.Println("Failed to write file:", err)
		return
	}

	// Replace old binary
	err = os.Rename(tmpFile, installPath)
	if err != nil {
		fmt.Println("Failed to replace old binary:", err)
		return
	}

	// Set execute permission (non-Windows)
	if OS != "windows" {
		if err := os.Chmod(installPath, 0755); err != nil {
			fmt.Println("Failed to set execute permission:", err)
			return
		}
	}

	fmt.Println("imotif-tools updated successfully!")
}
