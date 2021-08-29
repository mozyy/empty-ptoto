package main

import (
	"log"

	"github.com/gogo/protobuf/vanity/command"
	"protoc.yyue.dev/protoc-gen-orm/oorm"
)

func main() {
	// response := command.GeneratePlugin(command.Read(), &oorm.Orm{}, ".pb.orm.go")
	// command.Write(response)
	// TODO: 使用 https://github.com/golang/protobuf 开发
	response := command.GeneratePlugin(command.Read(), &oorm.Oorm{}, ".pb.orm.go")
	if len(response.String()) == 0 {
		log.Printf("empty response")
	}
	command.Write(response)
}
