package html

import (
	"fmt"
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
	if entering {
		_, _ = w.WriteString("<figure>")
		image := node.(*customAst.Image)

		// caption
		if len(image.Caption) > 0 {
			_, _ = w.WriteString("<figcaption>")
			_, _ = w.WriteString(image.Caption)
			_, _ = w.WriteString("</figcaption>")

		}

		// image with link
		if image.ItemId > 0 {
			imageUrl := fmt.Sprintf("https://dummyimage.com/%d/", image.ItemId)
			_, _ = w.WriteString(`<a href="" target="_blank">`)
			_, _ = w.WriteString(`<img`)
			_, _ = w.WriteString(fmt.Sprintf(` src="%s"`, imageUrl))
			_, _ = w.WriteString(` />`)
			_, _ = w.WriteString("</a>")
		}
	} else {
		_, _ = w.WriteString("</figure>\n")
	}
	return ast.WalkContinue, nil
}
