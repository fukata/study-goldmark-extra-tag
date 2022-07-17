package main

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type customImageParser struct {
}

func (p *customImageParser) SetParserOption(config *parser.Config) {
}

type CustomImageOption interface {
	parser.Option
}

func WithCustomImage() CustomImageOption {
	return &customImageParser{}
}

func (p *customImageParser) Trigger() []byte {
	return nil
}

func (p *customImageParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	return parent, parser.NoChildren
}

func (p *customImageParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.NoChildren
}

func (p *customImageParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
}

func (p *customImageParser) CanInterruptParagraph() bool {
	return false
}

func (p *customImageParser) CanAcceptIndentedLine() bool {
	return false
}
