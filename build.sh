#!/usr/bin/env bash

rm -rf bin

mkdir bin

go build -o bin/zipbackpack main.go

cd web

zip ../bin/main.go.zip -q -r *

cd ..

cat bin/main.go.zip >> bin/zipbackpack

zip -q -A bin/zipbackpack

bin/zipbackpack
