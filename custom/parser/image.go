package parser

import (
	customAst "github.com/fukata/study-goldmark-extra-tag/custom/ast"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"regexp"
	"strconv"
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

var imageAttributeRe = regexp.MustCompile(`\s+([^=]+)="([^"]+)"`)
var imageTagRe = regexp.MustCompile(`^\[image +.+](?:\r?\n)?`)

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

	if !imageTagRe.Match(line) {
		return nil, parser.NoChildren
	}

	var itemId int64
	var caption string
	matches := imageAttributeRe.FindAllSubmatch(line, -1)
	for _, m := range matches {
		key := string(m[1])
		value := string(m[2])

		switch key {
		case "id":
			itemId, _ = strconv.ParseInt(value, 10, 64)
		case "caption":
			caption = value
		}
	}

	node := customAst.NewImage(itemId, caption)
	return node, parser.NoChildren
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
