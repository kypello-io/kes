// Copyright 2022 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package sys

import (
	"errors"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

// Populated by goreleaser ldflags at build time.
// When empty, falls back to VCS build info.
var (
	version  string
	commitID string
)

// BinaryInfo contains build information about a Go binary.
type BinaryInfo struct {
	Version  string // The version of this binary
	CommitID string // The git commit hash
	Runtime  string // The Go runtime version, e.g. go1.21.0
	Compiler string // The Go compiler used to build this binary
}

// ReadBinaryInfo returns the ReadBinaryInfo about this program.
func ReadBinaryInfo() (BinaryInfo, error) { return readBinaryInfo() }

var readBinaryInfo = sync.OnceValues[BinaryInfo, error](func() (BinaryInfo, error) {
	const (
		DefaultVersion  = "<unknown>"
		DefaultCommitID = "<unknown>"
		DefaultCompiler = "<unknown>"
	)
	binaryInfo := BinaryInfo{
		Version:  DefaultVersion,
		CommitID: DefaultCommitID,
		Runtime:  runtime.Version(),
		Compiler: DefaultCompiler,
	}

	// Prefer ldflags values injected by goreleaser at build time.
	if version != "" {
		binaryInfo.Version = version
	}
	if commitID != "" {
		binaryInfo.CommitID = commitID
	}
	if binaryInfo.Version != DefaultVersion && binaryInfo.CommitID != DefaultCommitID {
		return binaryInfo, nil
	}

	// Fall back to VCS build info for local `go build`.
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return binaryInfo, errors.New("sys: binary does not contain build info")
	}

	const (
		GitTimeKey     = "vcs.time"
		GitRevisionKey = "vcs.revision"
		CompilerKey    = "-compiler"
	)
	for _, setting := range info.Settings {
		switch setting.Key {
		case GitTimeKey:
			if binaryInfo.Version == DefaultVersion {
				binaryInfo.Version = strings.ReplaceAll(setting.Value, ":", "-")
			}
		case GitRevisionKey:
			if binaryInfo.CommitID == DefaultCommitID {
				binaryInfo.CommitID = setting.Value
			}
		case CompilerKey:
			binaryInfo.Compiler = setting.Value
		}
	}
	return binaryInfo, nil
})
