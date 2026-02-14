#!/usr/bin/env nu

const root = path self .
cd $root

rm -rf $"($root)/youtubei"
git clone https://github.com/SuspiciousLookingOwl/youtubei $"($root)/youtubei"
pnpm -C $"($root)/youtubei" install

pnpm install
pnpm webpack
rm -rf $"($root)/youtubei"

ls $"($root)/dist/index.js" | get 0.size | to json
