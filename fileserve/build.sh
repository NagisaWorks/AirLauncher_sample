#!/bin/sh

GOOS=darwin GOARCH=amd64 go build -o mac/fileserve fileserve.go
cp list.xml mac
cp meta.json mac
zip -r ../fileserve_mac mac
rm -r mac

GOOS=windows GOARCH=386 go build -o windows/fileserve.exe fileserve.go
cp list.xml windows
cp meta.json windows
zip -r ../fileserve_win windows
rm -r windows
