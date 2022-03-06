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

package handler

import (
	"bufio"
	"bytes"
	"context"
	_ "embed" // to allow embedding files.
	"errors"
	"fmt"
	"os"
	"text/template"

	bootstrapv3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"sigs.k8s.io/yaml"

	"github.com/dio/proxy/config"
	_ "github.com/dio/proxy/internal/extensions" // to allow resolving type URLs.
	"github.com/dio/proxy/internal/options"
)

// TODO(dio): To add xds-v3-google.yaml for using google_grpc.
//go:embed templates/xds-v3-envoy.yaml
var xdsV3Envoy string

//go:embed templates/xds-v3-google.yaml
var xdsV3Google string

// Args holds arguments values and a cleanup function to remove all produced files (or objects).
type Args struct {
	Values  []string
	Cleanup func() error
}

// New returns a new handler.
func New(c *config.Bootstrap) *Handler {
	return &Handler{
		c: c,
	}
}

// Handler prepares proxy config bootstrapping process.
type Handler struct {
	c *config.Bootstrap
}

func (h *Handler) Run(ctx context.Context) error {
	// Currently, handler does nothing.
	<-ctx.Done()
	return nil
}

// Args produces arguments for spawning the proxy process. When building the arguments, it may
// produces some files. This returns an Args instance which has a Cleanup() to remove all of the
// produced files.
func (h *Handler) Args() (*Args, error) {
	var err error
	var args []string
	var configPath string
	var adminAddressPath string

	// Path to configuration file.
	if !contains(options.ForwardedArgs, "-c") || !contains(options.ForwardedArgs, "--config-path") {
		configPath, err = buildConfigPath(h.c)
		if err != nil {
			return nil, fmt.Errorf("failed to build config: %w", err)
		}
		args = append(args, "-c", configPath)
	}

	if !contains(options.ForwardedArgs, "--use-dynamic-base-id") || !contains(options.ForwardedArgs, "--base-id") {
		// The server chooses a base ID dynamically. Supersedes a static base ID. May not be used when
		// the restart epoch is non-zero.
		args = append(args, "--use-dynamic-base-id") // So we can run multiple proxies.
	}

	if !contains(options.ForwardedArgs, "--admin-address-path") {
		adminAddressPath, err = createAdminAddressPath()
		if err != nil {
			return nil, err
		}
		args = append(args, "--admin-address-path", adminAddressPath)
	}

	if !contains(options.ForwardedArgs, "--disable-hot-restart") {
		args = append(args, "--disable-hot-restart") // Disable hot restart functionality.
	}

	args = append(args, options.ForwardedArgs...)
	return &Args{
		Values: args,
		Cleanup: func() error {
			if configPath != "" {
				return os.Remove(configPath)
			}
			if adminAddressPath != "" {
				return os.Remove(adminAddressPath)
			}
			return nil
		},
	}, nil
}

func buildConfigPath(c *config.Bootstrap) (string, error) {
	out, err := os.CreateTemp("", "*_config.yaml")
	if err != nil {
		return "", err
	}
	defer func() {
		_ = out.Close()
	}()

	cfg := xdsV3Envoy
	if c.UseGoogleGRPC {
		cfg = xdsV3Google
	}
	tmpl, err := template.New("bootstrap").Parse(cfg)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	tmpl.Execute(writer, c)
	if err != nil {
		return "", err
	}
	_ = writer.Flush()

	err = validateBootstrap(buf.Bytes())
	if err != nil {
		return "", err
	}

	_, err = out.Write(buf.Bytes())
	if err != nil {
		return "", err
	}

	return out.Name(), nil
}

func createAdminAddressPath() (string, error) {
	out, err := os.CreateTemp("", "*_admin.txt")
	if err != nil {
		return "", err
	}
	defer func() {
		_ = out.Close()
	}()

	return out.Name(), nil
}

func contains(entries []string, target string) bool {
	for _, entry := range entries {
		if entry == target {
			return true
		}
	}
	return false
}

func validateBootstrap(bytes []byte) error {
	j, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		return err
	}
	var bootstrap bootstrapv3.Bootstrap
	err = protojson.Unmarshal(j, &bootstrap)
	if err != nil {
		return err
	}

	if bootstrap.GetNode().GetId() == "" {
		return errors.New("missing node ID")
	}

	// Currently, we require dynamic resources to be always available, but when the config is
	// static-only, we should relax this and do another way of validating the required bootstrap
	// components.
	if (bootstrap.DynamicResources) == nil {
		return errors.New("missing dynamic resources")
	}

	// TODO(dio): Validate using the fine-grained JSON schema so we can refer to the offending lines.
	err = bootstrap.ValidateAll()
	if err != nil {
		return err
	}
	return nil
}
