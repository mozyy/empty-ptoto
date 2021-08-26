package main

import (
	"log"

	"github.com/gogo/protobuf/vanity/command"
	"protoc.yyue.dev/protoc-gen-orm/oorm"
)

func main() {
	o := &oorm.Orm{}
	response := command.GeneratePlugin(command.Read(), o, ".pb.gorm.go")
	if len(response.String()) == 0 {
		log.Printf("empty response")
	}
	command.Write(response)
}
