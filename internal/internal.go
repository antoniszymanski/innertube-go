// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

import (
	"unsafe"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/require"
)

func MustCompile(name, src string) *goja.Program {
	ast, err := goja.Parse(name, src, parser.WithDisableSourceMaps)
	if err != nil {
		panic(err)
	}
	program, err := goja.CompileAST(ast, true)
	if err != nil {
		panic(err)
	}
	return program
}

func NewVM() *goja.Runtime {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("js", false))
	vm.SetParserOptions(parser.WithDisableSourceMaps)
	registry.Enable(vm)
	return vm
}

var registry = require.NewRegistryWithLoader(
	func(string) ([]byte, error) {
		return nil, require.ModuleFileDoesNotExistError
	},
)

func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
