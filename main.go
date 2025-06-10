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
			// 生成 xxx_myhttp.pb.go 文件
			if err := parse.GenerateFile(gen, f); err != nil {
				return err
			}
			// 基于当前目录，生成handler目录，并生成 xxx_handler.go 文件, 该文件
		}
		return nil
	})
}
