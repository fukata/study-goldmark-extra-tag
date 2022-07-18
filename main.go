package main

import (
	"bytes"
	"fmt"
	customParser "github.com/fukata/study-goldmark-extra-tag/custom/parser"
	customHtml "github.com/fukata/study-goldmark-extra-tag/custom/renderer/html"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func main() {
	source := []byte(`
# Hoge

message1

## Foo

message2

message3

# CustomExtension Image

[image id="1" caption="TEST"]

`)

	md := goldmark.New(
		goldmark.WithExtensions(customParser.CustomExtension),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRenderer(customHtml.NewCustomRenderer()),
	)

	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	fmt.Println("Input:")
	fmt.Println(string(source))

	fmt.Println("Output:")
	fmt.Println(buf.String())
}
