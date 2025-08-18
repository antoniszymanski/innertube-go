#!/usr/bin/env nu

cd $env.FILE_PWD
rm -rf src
git clone https://github.com/SuspiciousLookingOwl/youtubei.git src
cd src
pnpm i
cd ..

pnpm webpack
rm -rf src

ls dist/index.js | get 0.size | to json
