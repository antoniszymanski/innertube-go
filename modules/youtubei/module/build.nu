#!/usr/bin/env nu

cd $env.FILE_PWD

rm -rf youtubei
git clone https://github.com/antoniszymanski/youtubei.git
cd youtubei
pnpm i
cd ..

pnpm webpack
rm -rf youtubei

ls dist/index.js | get 0.size | to json
