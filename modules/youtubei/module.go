// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtubei

import (
	_ "embed"

	"github.com/antoniszymanski/innertube-go/internal"
	_ "github.com/antoniszymanski/innertube-go/modules/fetch"
	"github.com/grafana/sobek"
	_ "github.com/ohayocorp/sobek_nodejs/console"
	"github.com/ohayocorp/sobek_nodejs/process"
	"github.com/ohayocorp/sobek_nodejs/require"
	_ "github.com/ohayocorp/sobek_nodejs/url"
)

const ModuleName = "youtubei"

//go:embed module/dist/index.js
var source string

var program = internal.MustCompile(ModuleName, source)

func Require(vm *sobek.Runtime, module *sobek.Object) {
	process.Enable(vm)
	fn, err := vm.RunProgram(program)
	if err != nil {
		panic(err)
	}
	call, ok := sobek.AssertFunction(fn)
	if !ok {
		panic(require.InvalidModuleError)
	}
	_, err = call(nil, module)
	if err != nil {
		panic(err)
	}
}

func Enable(vm *sobek.Runtime) {
	m := require.Require(vm, ModuleName).ToObject(vm)
	vm.Set(ModuleName, m) //nolint:errcheck
}

func init() {
	require.RegisterCoreModule(ModuleName, Require)
}
