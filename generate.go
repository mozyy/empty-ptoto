package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

//go:generate go run $GOFILE

// const modulesPath = "submodules"
const modulesPath = "../"

func main() {
	// var useGorm = flag.Bool("gorm", false, "use gorm gen")
	flag.Parse()
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	goOutPath := path.Join(dir, modulesPath)
	webOutPath := path.Join(dir, modulesPath, "empty-frontend/src")
	webNgOutPath := path.Join(dir, modulesPath, "empty-frontend-ng/src")
	// openApiOutPath := path.Join(dir, "submodules/empty-frontend/src/openapi")
	os.RemoveAll(path.Join(goOutPath, "empty-news/proto"))
	os.MkdirAll(path.Join(goOutPath, "empty-news/proto"), 0551)
	os.RemoveAll(path.Join(webOutPath, "proto"))
	os.MkdirAll(path.Join(webOutPath, "proto"), 0551)
	os.RemoveAll(path.Join(webNgOutPath, "proto"))
	os.MkdirAll(path.Join(webNgOutPath, "proto"), 0551)

	protoPath := path.Join(dir, "proto")

	err = filepath.Walk(protoPath, func(p string, info os.FileInfo, err error) error {
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
				// fmt.Sprintf("--proto_path=%s/src/github.com/googleapis/googleapis", os.Getenv("GOPATH")),
				// 分开写会生成两个文件, 下面的写一起只生成一个文件 ,
				// 下面的import的context是golang.org/x/net/context, 以被内部context取代,
				// 所以用上面的

				// fmt.Sprintf("--micro_out=%s", goOutPath),
				fmt.Sprintf("--go_out=module=github.com/mozyy:%s", goOutPath),
				fmt.Sprintf("--go-grpc_out=module=github.com/mozyy:%s", goOutPath),
				fmt.Sprintf("--grpc-gateway_out=module=github.com/mozyy,logtostderr=true,generate_unbound_methods=true:%s", goOutPath),
				fmt.Sprintf("--gorm_out=module=github.com/mozyy:%s", goOutPath),
				// fmt.Sprintf("--orm_out=module=github.com/mozyy:%s", goOutPath),
				// fmt.Sprintf("--gogo_out=:%s", "."),
				// fmt.Sprintf("--go_out=plugins=micro:%s", goOutPath),
				// js
				fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", webOutPath),
				fmt.Sprintf("--grpc-web_out=import_style=typescript,mode=grpcwebtext:%s", webOutPath),
				fmt.Sprintf("--openapiv2_out=logtostderr=true:%s", webOutPath),
				fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", webNgOutPath),
				fmt.Sprintf("--grpc-web_out=import_style=typescript,mode=grpcwebtext:%s", webNgOutPath),
				fmt.Sprintf("--openapiv2_out=logtostderr=true,generate_unbound_methods=true:%s", webNgOutPath),
				// openapi
				// fmt.Sprintf("--openapi_out=%s", openApiOutPath),

				// relPath,
			}
			// if *useGorm {
			// 	args = append(args,
			// 		// TODO: 先生成到 ./github.com/mozyy/empty-news, 再用脚本复制过去
			// 		// fmt.Sprintf("--gorm_out=:%s", "."))
			// 		fmt.Sprintf("--orm_out=:%s", "."))
			// }
			args = append(args, relPath)
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
	// if *useGorm {
	// 	fmt.Println("2222")
	// 	err := copy.DirCopy(path.Join(dir, "./github.com/mozyy/empty-news/proto"),
	// 		path.Join(dir, "../empty-news/proto"), copy.Content, false)
	// 	os.RemoveAll(path.Join(dir, "github.com"))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}
