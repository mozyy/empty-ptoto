git submodule --quiet update --init --recursiv


go install ./protoc-gen-orm/protoc-gen-orm.go 
protoc -I=. -I=/root/.local/include --orm_out=. ./protoc-gen-orm/test.proto