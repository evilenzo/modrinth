package main

import (
	"log/slog"
	"os"
	"path"

	"github.com/hashicorp/go-getter"
)

//go:generate go run generate.go
//go:generate go tool ogen --target ../../pkg/client -package client --clean openapi.yaml

const openapiURL = "https://raw.githubusercontent.com/modrinth/code/refs/heads/main/apps/docs/public/openapi.yaml"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	pwd, err := os.Getwd()
	if err != nil {
		logger.Error("get pwd", "error", err)

		return
	}

	if err := getter.GetFile(path.Join(pwd, "openapi.yaml"), openapiURL); err != nil {
		logger.Error("get openapi", "error", err)
	}
}
