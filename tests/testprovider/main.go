// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// A provider with resources for use in tests.
package main

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	pbempty "github.com/golang/protobuf/ptypes/empty"
)

const (
	providerName = "testprovider"
	version      = "0.0.1"
)

// Minimal set of methods to implement a basic provider.
type resourceProvider interface {
	Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error)
	Diff(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error)
	Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error)
	Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error)
	Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error)
	Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error)
}

var resourceProviders = map[string]resourceProvider{
	"testprovider:index:Random": &randomResourceProvider{},
	"testprovider:index:Echo":   &echoResourceProvider{},
}

func providerForURN(urn string) (resourceProvider, string, bool) {
	ty := string(resource.URN(urn).Type())
	provider, ok := resourceProviders[ty]
	return provider, ty, ok
}

//nolint: unused,deadcode
func main() {
	if err := provider.Main(providerName, func(host *provider.HostClient) (rpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version)
	}); err != nil {
		cmdutil.ExitError(err.Error())
	}
}

type testproviderProvider struct {
	host    *provider.HostClient
	name    string
	version string
}

func makeProvider(host *provider.HostClient, name, version string) (rpc.ResourceProviderServer, error) {
	// Return the new provider
	return &testproviderProvider{
		host:    host,
		name:    name,
		version: version,
	}, nil
}

// CheckConfig validates the configuration for this provider.
func (k *testproviderProvider) CheckConfig(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (k *testproviderProvider) DiffConfig(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (k *testproviderProvider) Configure(_ context.Context, req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	return &rpc.ConfigureResponse{}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (k *testproviderProvider) Invoke(_ context.Context, req *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	tok := req.GetTok()
	return nil, fmt.Errorf("Unknown Invoke token '%s'", tok)
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (k *testproviderProvider) StreamInvoke(req *rpc.InvokeRequest,
	server rpc.ResourceProvider_StreamInvokeServer) error {
	tok := req.GetTok()
	return fmt.Errorf("Unknown StreamInvoke token '%s'", tok)
}

func (k *testproviderProvider) Call(_ context.Context, req *rpc.CallRequest) (*rpc.CallResponse, error) {
	tok := req.GetTok()
	return nil, fmt.Errorf("Unknown Call token '%s'", tok)
}

func (k *testproviderProvider) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	provider, ty, ok := providerForURN(req.GetUrn())
	if !ok {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}
	return provider.Check(ctx, req)
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (k *testproviderProvider) Diff(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	provider, ty, ok := providerForURN(req.GetUrn())
	if !ok {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}
	return provider.Diff(ctx, req)
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (k *testproviderProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	provider, ty, ok := providerForURN(req.GetUrn())
	if !ok {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}
	return provider.Create(ctx, req)
}

// Read the current live state associated with a resource.
func (k *testproviderProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	provider, ty, ok := providerForURN(req.GetUrn())
	if !ok {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}
	return provider.Read(ctx, req)
}

// Update updates an existing resource with new values.
func (k *testproviderProvider) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	provider, ty, ok := providerForURN(req.GetUrn())
	if !ok {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}
	return provider.Update(ctx, req)
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (k *testproviderProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	provider, ty, ok := providerForURN(req.GetUrn())
	if !ok {
		return nil, fmt.Errorf("Unknown resource type '%s'", ty)
	}
	return provider.Delete(ctx, req)
}

// Construct creates a new component resource.
func (k *testproviderProvider) Construct(_ context.Context, _ *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	panic("Construct not implemented")
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (k *testproviderProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: k.version,
	}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (k *testproviderProvider) GetSchema(ctx context.Context,
	req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	return &rpc.GetSchemaResponse{}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (k *testproviderProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	return &pbempty.Empty{}, nil
}
