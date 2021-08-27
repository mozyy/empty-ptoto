package oorm

import (
	"log"
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"protoc.yyue.dev/protoc-gen-orm/orm"
)

type Oorm struct {
	*generator.Generator
	generator.PluginImports
}

// Name identifies the plugin.
func (o *Oorm) Name() string {
	return "gorm"
}

// Init is called once after data structures are built but before
// code generation begins.
func (o *Oorm) Init(g *generator.Generator) {
	o.Generator = g
}

// Generate produces the code generated by the plugin for this file,
// except for the imports, by calling the generator's methods P, In, and Out.
func (o *Oorm) Generate(file *generator.FileDescriptor) {
	o.PluginImports = generator.NewPluginImports(o.Generator)

	timePkg := o.NewImport("time")
	for _, message := range file.Messages() {
		if message.GetOptions().GetMapEntry() {
			continue
		}
		if message.Options != nil {
			v, err := proto.GetExtension(message.Options, orm.E_Ormable)
			if err == nil {
				opts, ok := v.(*bool)
				if ok && *opts {
					log.Println("install: ", message.TypeName())
					// pase message
					o.P(`type `, message.TypeName()[0], `ORM struct {`)
					fields := fieldSort(message.GetField())
					sort.Sort(fields)
					for _, field := range fields {
						o.OneOfTypeName(message, field)
						o.Out()
						fieldType := o.getFieldType(message, field)
						fieldName := generator.CamelCase(field.GetName())
						log.Println("install: ", field)
						if fieldType != "" {
							o.P(fieldName, ` `, fieldType, o.renderGormTag(fieldName))
						}
					}
					o.P(`}`)
				}
			}
		}
	}

}

// GenerateImports produces the import declarations for this file.
// It is called after Generate.
func (o *Oorm) GenerateImports(file *generator.FileDescriptor) {
	panic("not implemented") // TODO: Implement
}
