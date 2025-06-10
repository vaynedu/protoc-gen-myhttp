package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"protoc-gen-myhttp/parse"
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
