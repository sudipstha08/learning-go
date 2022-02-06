#!/usr/bin/env bash
set -xe

# install packages and dependencies
go get github.com/gin-gonic/gin

go get gopkg.in/go-playground/validator.v9

#build command
go build -o bin/application server.go