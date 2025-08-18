// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package fetch

import (
	_ "embed"

	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const ModuleName = "fetch"

//go:embed module/dist/index.js
var source string

var program = internal.MustCompile(
	ModuleName,
	"(function(exports,require,module,__filename,__dirname,__fetch){"+source+"})",
)

func Require(vm *goja.Runtime, module *goja.Object) {
	fn, err := vm.RunProgram(program)
	if err != nil {
		panic(err)
	}
	call, ok := goja.AssertFunction(fn)
	if !ok {
		panic(require.InvalidModuleError)
	}
	exports := module.Get("exports")
	require := vm.Get("require")
	__filename := vm.ToValue(ModuleName + "/index.js")
	__dirname := vm.ToValue(ModuleName)
	__fetch := vm.ToValue(fetch)
	_, err = call(exports, exports, require, module, __filename, __dirname, __fetch)
	if err != nil {
		panic(err)
	}
}

func Enable(vm *goja.Runtime) {
	m := require.Require(vm, ModuleName).ToObject(vm)
	for _, key := range m.Keys() {
		if key != "default" {
			vm.Set(key, m.Get(key)) //nolint:errcheck
		}
	}
}

func init() {
	require.RegisterCoreModule(ModuleName, Require)
}
