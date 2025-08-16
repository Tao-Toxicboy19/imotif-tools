package update

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

type SelfUpdater struct {
	OS         string
	InstallDir string
	URLMap     map[string]string
}

func NewSelfUpdater() (*SelfUpdater, error) {
	installPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to determine binary path: %w", err)
	}

	return &SelfUpdater{
		OS:         runtime.GOOS,
		InstallDir: installPath,
		URLMap: map[string]string{
			"darwin":  "https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/imotif-tools-macos",
			"linux":   "https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/imotif-tools-linux",
			"windows": "https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/imotif-tools.exe",
		},
	}, nil
}

func (s *SelfUpdater) Run() error {
	binaryURL, ok := s.URLMap[s.OS]
	if !ok {
		return fmt.Errorf("unsupported OS: %s", s.OS)
	}

	fmt.Println("Downloading latest version from:", binaryURL)

	// 1. Download
	resp, err := http.Get(binaryURL)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status %s", resp.Status)
	}

	// 2. Save to temp
	tmpFile := s.InstallDir + ".tmp"
	out, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	if _, err := io.Copy(out, resp.Body); err != nil {
		out.Close()
		return fmt.Errorf("failed to write temp file: %w", err)
	}
	out.Close()

	// 3. Replace old binary
	if err := os.Rename(tmpFile, s.InstallDir); err != nil {
		return fmt.Errorf("failed to replace binary: %w", err)
	}

	// 4. Set permissions (non-Windows)
	if s.OS != "windows" {
		if err := os.Chmod(s.InstallDir, 0755); err != nil {
			return fmt.Errorf("failed to set execute permission: %w", err)
		}
	}

	fmt.Println("imotif-tools updated successfully!")
	return nil
}
