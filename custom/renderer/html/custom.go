package html

import (
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

func NewCustomRenderer() renderer.Renderer {
	return renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(html.NewRenderer(), 1000),
			util.Prioritized(NewRenderer(), 1100),
		),
	)
}
