#!/usr/bin/env nu

const root = path self .
cd $root

rm -rf youtubei
mkdir youtubei
let temp = mktemp --dry
http get https://github.com/SuspiciousLookingOwl/youtubei/archive/refs/heads/development.zip | save $temp
bsdtar -xf $temp -C youtubei --strip-components 1 youtubei-development
rm -f $temp

pnpm -C youtubei install
pnpm install
pnpm webpack
rm -rf youtubei

ls dist/index.js | get 0.size | to json
