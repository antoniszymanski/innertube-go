#!/usr/bin/env nu

const root = path self .
cd $root

pnpm install
pnpm webpack

ls dist/index.js | get 0.size | to json
