#!/usr/bin/env sh

# container version
VERSION=v0.1.0

docker build -t weismax/todo-ssr:$VERSION ..
docker push weismax/todo-ssr:$VERSION
