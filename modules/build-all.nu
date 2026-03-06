#!/usr/bin/env nu

const root = path self .
for path in (glob -Dl $'($root)/*/module/build.nu') {
  nu $path
}
