// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package main

import (
	"fmt"

	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi-fabric/pkg/resource"
	"github.com/pulumi/pulumi-fabric/pkg/resource/provider"
	"github.com/pulumi/pulumi-fabric/pkg/tokens"
	lumirpc "github.com/pulumi/pulumi-fabric/sdk/proto/go"
	"golang.org/x/net/context"

	"github.com/pulumi/pulumi-github/provider/ghctx"
	"github.com/pulumi/pulumi-github/provider/repos/issues"
	"github.com/pulumi/pulumi-github/provider/webhooks"
)

// Provider implements the AWS resource provider's operations for all known AWS types.
type Provider struct {
	impls map[tokens.Type]lumirpc.ResourceProviderServer
}

// NewProvider creates a new provider instance with server objects registered for every resource type.
func NewProvider(host *provider.HostClient) (*Provider, error) {
	return &Provider{
		impls: map[tokens.Type]lumirpc.ResourceProviderServer{
			issues.LabelToken:          issues.NewLabelProvider(host),
			issues.MilestoneToken:      issues.NewMilestoneProvider(host),
			webhooks.SubscriptionToken: webhooks.NewSubscriptionProvider(host),
		},
	}, nil
}

var _ lumirpc.ResourceProviderServer = (*Provider)(nil)

// Configure configures the resource provider with "globals" that control its behavior.
func (p *Provider) Configure(ctx context.Context, req *lumirpc.ConfigureRequest) (*pbempty.Empty, error) {
	// Configure the context.
	vars := make(map[tokens.ModuleMember]string)
	for k, v := range req.GetVariables() {
		vars[tokens.ModuleMember(k)] = v
	}
	if err := ghctx.Configure(vars); err != nil {
		return nil, err
	}

	// Give all other type providers a chance to observe the variables if desired.
	for _, impl := range p.impls {
		if _, err := impl.Configure(ctx, req); err != nil {
			return nil, err
		}
	}
	return &pbempty.Empty{}, nil
}

// Check validates that the given property bag is valid for a resource of the given type.
func (p *Provider) Check(ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
	t := resource.URN(req.GetUrn()).Type()
	if prov, has := p.impls[t]; has {
		return prov.Check(ctx, req)
	}
	return nil, fmt.Errorf("Unrecognized resource type (Check): %v", t)
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
func (p *Provider) Create(ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
	t := resource.URN(req.GetUrn()).Type()
	if prov, has := p.impls[t]; has {
		return prov.Create(ctx, req)
	}
	return nil, fmt.Errorf("Unrecognized resource type (Create): %v", t)
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (p *Provider) Diff(
	ctx context.Context, req *lumirpc.DiffRequest) (*lumirpc.DiffResponse, error) {
	t := resource.URN(req.GetUrn()).Type()
	if prov, has := p.impls[t]; has {
		return prov.Diff(ctx, req)
	}
	return nil, fmt.Errorf("Unrecognized resource type (InspectChange): %v", t)
}

// Update updates an existing resource with new values.  Only those values in the provided property bag are updated
// to new values.  The resource ID is returned and may be different if the resource had to be recreated.
func (p *Provider) Update(ctx context.Context, req *lumirpc.UpdateRequest) (*lumirpc.UpdateResponse, error) {
	t := resource.URN(req.GetUrn()).Type()
	if prov, has := p.impls[t]; has {
		return prov.Update(ctx, req)
	}
	return nil, fmt.Errorf("Unrecognized resource type (Update): %v", t)
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
func (p *Provider) Delete(ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
	t := resource.URN(req.GetUrn()).Type()
	if prov, has := p.impls[t]; has {
		return prov.Delete(ctx, req)
	}
	return nil, fmt.Errorf("Unrecognized resource type (Delete): %v", t)
}
