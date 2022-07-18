package ast

import "github.com/yuin/goldmark/ast"

var KindImage = ast.NewNodeKind("CustomImage")

type Image struct {
	ast.BaseBlock
	ItemId  int64
	Caption string
}

func (n *Image) Kind() ast.NodeKind {
	return KindImage
}

func (n *Image) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

func NewImage(itemId int64, caption string) *Image {
	return &Image{
		BaseBlock: ast.BaseBlock{},
		ItemId:    itemId,
		Caption:   caption,
	}
}
