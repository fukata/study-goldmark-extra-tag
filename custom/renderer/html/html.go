package html

import (
	customAst "github.com/fukata/study-goldmark-extra-tag/custom/ast"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type Renderer struct {
	Config html.Config
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(customAst.KindImage, r.renderImage)
}

func NewRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &Renderer{
		Config: html.NewConfig(),
	}

	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}
