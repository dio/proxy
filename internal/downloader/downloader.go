// Copyright 2022 Dhi Aurrahman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package downloader

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/bazelbuild/bazelisk/httputil"
	"github.com/codeclysm/extract"
	"github.com/mitchellh/go-homedir"
	"github.com/ulikunitz/xz"

	"github.com/dio/proxy/internal/archives"
)

const (
	funcEHomeEnvKey = "FUNC_E_HOME"
)

// DownloadVersionedBinary returns the downloaded binary file path.
func DownloadVersionedBinary(ctx context.Context, archive archives.Archive, destDir string) (string, error) {
	destinationDir := filepath.Join(destDir, archive.BinaryDir())
	err := os.MkdirAll(destinationDir, 0o750)
	if err != nil {
		return "", fmt.Errorf("could not create directory %s: %v", destinationDir, err)
	}

	destinationPath := filepath.Join(destinationDir, archive.BinaryName())
	if _, err := os.Stat(destinationPath); err != nil {
		downloadURL := GetArchiveURL(archive)
		// TODO(dio): Streaming the bytes from remote file. We decided to use this for skipping copying
		// the retry logic that has already implemented in github.com/bazelbuild/bazelisk/httputil.
		data, _, err := httputil.ReadRemoteFile(downloadURL, "")
		if err != nil {
			return "", fmt.Errorf("failed to read remote file: %s: %w", downloadURL, err)
		}
		br := bufio.NewReader(bytes.NewBuffer(data))
		maybeXzHeader, err := br.Peek(xz.HeaderLen)
		if err != nil {
			return "", err
		}
		if xz.ValidHeader(maybeXzHeader) {
			var r *xz.Reader
			r, err = xz.NewReader(br)
			if err != nil {
				return "", err
			}
			err = extract.Tar(ctx, r, destDir, archive.Renamer())
		} else {
			err = extract.Gz(ctx, br, destDir, archive.Renamer())
		}
		if err != nil {
			return "", err
		}
		if _, err = os.Stat(destinationPath); err != nil {
			return "", fmt.Errorf("failed to extract the remote file from: %s: %w", downloadURL, err)
		}
		if err = os.Chmod(destinationPath, 0o755); err != nil { //nolint:gosec
			return "", fmt.Errorf("could not chmod file %s: %v", destinationPath, err)
		}
	}
	return destinationPath, nil
}

// GetArchiveURL renders the archive URL pattern to return the actual archive URL.
func GetArchiveURL(archive archives.Archive) string {
	// TODO(dio): Use template instead of simple fmt.Sprintf.
	return fmt.Sprintf(archive.URLPattern(), archive.Version(), archive.Version(), runtime.GOOS) // We always do amd64, ignore the GOARCH for now.
}

func Download(ctx context.Context) (string, error) {
	// The func-e home is defined by $FUNC_E_HOME or if not defined: ~/.func-e.
	// The binary should be downloaded to $FUNC_E_HOME/versions/1.20.1/bin/envoy
	funcEHome := os.Getenv(funcEHomeEnvKey)
	if funcEHome == "" {
		home, _ := homedir.Dir()
		funcEHome = filepath.Join(home, ".func-e")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	archive := &archives.Proxy{VersionUsed: proxyVersion(funcEHome)} // TODO(dio): fallback to default when older than the minimum version.
	return DownloadVersionedBinary(ctx, archive, funcEHome)
}

func proxyVersion(dir string) string {
	b, err := os.ReadFile(filepath.Join(dir, "version"))
	if err != nil {
		return ""
	}
	captured := string(b)
	parts := strings.Split(captured, ".")
	if len(parts) == 3 {
		// Check if the binary is there. TODO(dio): Make sure we accommodate windows as well.
		_, err = os.Lstat(filepath.Join(dir, "versions", captured, "bin", "envoy"))
		if err != nil {
			return ""
		}
		return captured
	}
	dirs, _ := ioutil.ReadDir(filepath.Join(dir, "versions"))
	for i := len(dirs); i >= 0; i-- {
		version := fmt.Sprintf(captured+".%d", i)
		// Check if the binary is there. TODO(dio): Make sure we accommodate windows as well.
		_, err = os.Lstat(filepath.Join(dir, "versions", version, "bin", "envoy"))
		if err == nil {
			return version
		}
	}
	return ""
}
