#!/usr/bin/env sh

# container version
VERSION=v0.1.3

docker build -t weismax/todo-ssr:$VERSION ..
docker push weismax/todo-ssr:$VERSION
