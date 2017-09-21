package gendoc_test

import (
	"github.com/meteor/protoc-gen-doc/parser"
	"github.com/meteor/protoc-gen-doc/test"
	"testing"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	codeGenRequest, _ := test.MakeCodeGeneratorRequest()

	for i := 0; i < b.N; i++ {
		parser.ParseCodeRequest(codeGenRequest)
	}
}
