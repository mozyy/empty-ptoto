git submodule --quiet update --init --recursiv


// 要用这个生成 gogo版的orm.ptoto 文件
protoc -I=. -I=/root/.local/include --gogo_out=paths=source_relative,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. ./protoc-gen-orm/orm/orm.proto

go install ./protoc-gen-orm/protoc-gen-orm.go && \
protoc -I=. -I=/root/.local/include --orm_out=. ./protoc-gen-orm/test.proto