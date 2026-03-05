// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

import (
	"unsafe"

	"github.com/grafana/sobek"
	"github.com/grafana/sobek/parser"
	"github.com/ohayocorp/sobek_nodejs/require"
)

func MustCompile(name, src string) *sobek.Program {
	ast, err := sobek.Parse(name, src, parser.WithDisableSourceMaps)
	if err != nil {
		panic(err)
	}
	program, err := sobek.CompileAST(ast, true)
	if err != nil {
		panic(err)
	}
	return program
}

func NewVM() *sobek.Runtime {
	vm := sobek.New()
	vm.SetFieldNameMapper(sobek.TagFieldNameMapper("js", false))
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
