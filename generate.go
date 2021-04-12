package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

//go:generate go run $GOFILE

const (
	goOutPath  = "./submodules"
	webOutPath = "./submodules/empty-frontend/src/proto"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}
		// fmt.Println(relPath, relPath, info.IsDir(), filepath.Ext(relPath), filepath.Ext(relPath))
		if filepath.Ext(relPath) == ".proto" {

			args := []string{
				"--proto_path=./",
				fmt.Sprintf("--proto_path=%s/.local/include", os.Getenv("HOME")),
				fmt.Sprintf("--proto_path=%s/src/github.com/googleapis/googleapis", os.Getenv("GOPATH")),
				// 分开写会生成两个文件, 下面的写一起只生成一个文件 ,
				// 下面的import的context是golang.org/x/net/context, 以被内部context取代,
				// 所以用上面的

				fmt.Sprintf("--micro_out=%s", goOutPath),
				fmt.Sprintf("--go_out=%s", goOutPath),
				// fmt.Sprintf("--go_out=plugins=micro:%s", goOutPath),
				// js
				fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", webOutPath),
				fmt.Sprintf("--grpc-web_out=import_style=typescript,mode=grpcwebtext:%s", webOutPath),
				relPath,
			}
			cmd := exec.Command("protoc", args...)
			cmd.Dir = dir
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("protoc [%s] error: %s, out: \n%s\n%s\n", relPath, err, out, args)
			} else {
				fmt.Printf("protoc success: %s\n", p)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
