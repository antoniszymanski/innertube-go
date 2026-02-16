// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtubei

import (
	_ "embed"

	"github.com/antoniszymanski/innertube-go/internal"
	_ "github.com/antoniszymanski/innertube-go/modules/fetch"
	"github.com/dop251/goja"
	_ "github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/process"
	"github.com/dop251/goja_nodejs/require"
	_ "github.com/dop251/goja_nodejs/url"
)

const ModuleName = "youtubei"

//go:embed module/dist/index.js
var source string

var program = internal.MustCompile(ModuleName, source)

func Require(vm *goja.Runtime, module *goja.Object) {
	process.Enable(vm)
	fn, err := vm.RunProgram(program)
	if err != nil {
		panic(err)
	}
	call, ok := goja.AssertFunction(fn)
	if !ok {
		panic(require.InvalidModuleError)
	}
	_, err = call(nil, module)
	if err != nil {
		panic(err)
	}
}

func Enable(vm *goja.Runtime) {
	m := require.Require(vm, ModuleName).ToObject(vm)
	vm.Set(ModuleName, m) //nolint:errcheck
}

func init() {
	require.RegisterCoreModule(ModuleName, Require)
}
