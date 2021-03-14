#!/bin/bash

set -xe


# Go API
oto -v -template ./server/api/server.go.plush \
    -out ./server/tasks/tasks.gen.go \
    -ignore Ignorer -pkg tasks ./server/api/...
gofmt -w ./server/tasks/tasks.gen.go

# Typescript API
oto -v -template tasks-web-client/src/app/client.ts.plush \
    -out tasks-web-client/src/app/oto.gen.ts \
    -ignore Ignorer -pkg tasks ./server/api/...
