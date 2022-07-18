package parser

import (
	"github.com/yuin/goldmark"
)

type customExtension struct {
}

var CustomExtension = &customExtension{}

func (c *customExtension) Extend(m goldmark.Markdown) {
	Image.Extend(m)
}
