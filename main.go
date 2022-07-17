package main

import (
	"bytes"
	"fmt"
	"github.com/fukata/study-goldmark-extra-tag/custom"
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

# Custom Image

[image id="1" caption="TEST"]

`)

	md := goldmark.New(
		goldmark.WithExtensions(custom.Custom),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
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
