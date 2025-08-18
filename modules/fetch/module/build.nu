#!/usr/bin/env nu

cd $env.FILE_PWD
pnpm webpack

ls dist/index.js | get 0.size | to json
