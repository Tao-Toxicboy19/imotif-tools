package cli

import (
	"fmt"

	figure "github.com/common-nighthawk/go-figure"
)

type Banner struct{}

func NewBanner() *Banner {
	return &Banner{}
}

func (c *Banner) PrintBanner() {
	fig := figure.NewFigure("IMOITF Tools", "", true)
	fig.Print()
}

func (c *Banner) PrintHelp() {
	fmt.Println(`imotif-tools - interactive git commit helper

Usage:
  imotif-tools                 Show banner
  imotif-tools commit "msg"    Run commit with message
  imotif-tools init            Setup alias
  imotif-tools update          Self-update
  imotif-tools -v              Show version
`)
}
