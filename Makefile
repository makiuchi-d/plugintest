srcs := $(wildcard *.go */*.go)

all: plugintest doubleid.so doublename.so

plugintest: $(srcs)
	go build -o "$@" cmd/main.go

%.so: plugins/%/main.go records/record.go
	go build -buildmode=plugin -o "$@" "./$<"

%.so: plugins/% records/record.go
	go build -buildmode=plugin -o "$@" "./$<"

run: plugintest
	./plugintest

run-id: plugintest doubleid.so
	./plugintest doubleid.so

run-name: plugintest doublename.so
	./plugintest doublename.so
