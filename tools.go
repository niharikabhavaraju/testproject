// +build tools

package main

import (
    _ "github.com/fullstorydev/grpcui/cmd/grpcui"
    _ "github.com/golang/protobuf/protoc-gen-go"
    _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
    _ "github.com/goreleaser/goreleaser"
    _ "github.com/spf13/cobra/cobra"
    _ "github.com/tebeka/go2xunit"
    _ "golang.org/x/lint/golint"
    _ "golang.org/x/perf/cmd/benchstat"
    _ "golang.org/x/tools/cmd/stringer"
)

