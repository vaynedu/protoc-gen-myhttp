package main

import (
	"github.com/vaynedu/protoc-gen-myhttp/parse"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		// parse all files
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err := parse.GenerateFile(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}
