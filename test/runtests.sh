#!/bin/zsh

go test ./... -v -coverprofile=test/coverage.out && go tool cover -html=test/coverage.out