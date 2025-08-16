package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

type FileWithContent struct {
	Filename string
	Content  string
}

type GitExec struct {
	Dir string
}

func NewGitExec() *GitExec {
	return &GitExec{}
}

func (g *GitExec) GetStagedFilesWithContent() ([]FileWithContent, error) {
	statusCmd := exec.Command("git", "diff", "--cached", "--name-status")
	if g.Dir != "" {
		statusCmd.Dir = g.Dir
	}

	var statusOut bytes.Buffer
	statusCmd.Stdout = &statusOut
	if err := statusCmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to get file status: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(statusOut.String()), "\n")
	var result []FileWithContent

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		status, file := parts[0], parts[1]
		content := ""

		if status != "D" {
			contentCmd := exec.Command("git", "show", ":"+file)
			if g.Dir != "" {
				contentCmd.Dir = g.Dir
			}

			var contentBuf bytes.Buffer
			contentCmd.Stdout = &contentBuf
			if err := contentCmd.Run(); err != nil {
				return nil, fmt.Errorf("failed to read file %s: %w", file, err)
			}
			content = contentBuf.String()
		}

		if g.isSkippableFile(file, content) {
			continue
		}

		result = append(result, FileWithContent{
			Filename: file,
			Content:  content,
		})
	}

	return result, nil
}

func (g *GitExec) isSkippableFile(filename, content string) bool {
	bannedExts := map[string]bool{
		".exe": true, ".bin": true, ".so": true, ".dll": true, ".zip": true,
		".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".ico": true, ".pdf": true,
	}

	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" || bannedExts[ext] {
		return true
	}

	if !utf8.ValidString(content) {
		return true
	}

	return false
}
