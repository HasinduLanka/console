#!/bin/sh

go mod tidy

echo "Git tag: $1" >> .gittag

git add .
git commit -m "Version $1"
git push origin

git tag v$1 -m "Release $1"
git push origin v$1

GOPROXY=proxy.golang.org go list -m github.com/HasinduLanka/console@v$1
