package parser

import (
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

/**
独自画像の記法

Input:

[image id="${image_id}" caption="${caption}"]

Output:

<figure>
	<figcaption>${caption}</figcaption>
	<a href="${url}" target="_blank">
		<img src="${image_url}" alt="${image_caption}">
	</a>
</figure>
*/

type imageParser struct {
}

func NewImageParser() parser.BlockParser {
	return &imageParser{}
}

func (p *imageParser) Trigger() []byte {
	return []byte{'['}
}

func (p *imageParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	fmt.Println(string(line))

	return nil, parser.NoChildren
}

func (p *imageParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.NoChildren
}

func (p *imageParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
}

func (p *imageParser) CanInterruptParagraph() bool {
	return false
}

func (p *imageParser) CanAcceptIndentedLine() bool {
	return false
}

type image struct {
}

var Image = &image{}

func (e *image) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(util.Prioritized(NewImageParser(), 999)),
	)
}
