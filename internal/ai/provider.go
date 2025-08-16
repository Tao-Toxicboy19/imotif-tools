package ai

import (
	"context"

	"github.com/imotif-tools/internal/git"
)

type Provider interface {
	Suggest(ctx context.Context, files []git.FileWithContent) (string, error)
}
