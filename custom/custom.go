package custom

import "github.com/yuin/goldmark"

type custom struct {
}

var Custom = &custom{}

func (c *custom) Extend(m goldmark.Markdown) {
	Image.Extend(m)
}
