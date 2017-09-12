// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package webhooks

import (
    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/pulumi-fabric/pkg/resource"
    "github.com/pulumi/pulumi-fabric/pkg/resource/plugin"
    "github.com/pulumi/pulumi-fabric/pkg/tokens"
    "github.com/pulumi/pulumi-fabric/pkg/util/contract"
    "github.com/pulumi/pulumi-fabric/pkg/util/mapper"
    lumirpc "github.com/pulumi/pulumi-fabric/sdk/proto/go"
)

/* Marshalable Config structure(s) */

// Config is a marshalable representation of its corresponding IDL type.
type Config struct {
    URL string `lumi:"url"`
    ContentType *ContentType `lumi:"contentType,optional"`
    Secret *string `lumi:"secret,optional"`
    InsecureSSL *bool `lumi:"insecureSSL,optional"`
}

// Config's properties have constants to make dealing with diffs and property bags easier.
const (
    Config_URL = "url"
    Config_ContentType = "contentType"
    Config_Secret = "secret"
    Config_InsecureSSL = "insecureSSL"
)

/* RPC stubs for Subscription resource provider */

// SubscriptionToken is the type token corresponding to the Subscription package type.
const SubscriptionToken = tokens.Type("github:webhooks/subscription:Subscription")

// SubscriptionProviderOps is a pluggable interface for Subscription-related management functionality.
type SubscriptionProviderOps interface {
    Configure(ctx context.Context, vars map[tokens.ModuleMember]string) error
    Check(ctx context.Context, obj *Subscription, property string) error
    Diff(ctx context.Context, id resource.ID,
        old *Subscription, new *Subscription, diff *resource.ObjectDiff) ([]string, error)
    Create(ctx context.Context, obj *Subscription) (resource.ID, error)
    Update(ctx context.Context, id resource.ID,
        old *Subscription, new *Subscription, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID, obj Subscription) error
}

// SubscriptionProvider is a dynamic gRPC-based plugin for managing Subscription resources.
type SubscriptionProvider struct {
    ops SubscriptionProviderOps
}

// NewSubscriptionProvider allocates a resource provider that delegates to a ops instance.
func NewSubscriptionProvider(ops SubscriptionProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &SubscriptionProvider{ops: ops}
}

func (p *SubscriptionProvider) Configure(
    ctx context.Context, req *lumirpc.ConfigureRequest) (*pbempty.Empty, error) {
    vars := make(map[tokens.ModuleMember]string)
    for k, v := range req.GetVariables() {
        vars[tokens.ModuleMember(k)] = v
    }
    if err := p.ops.Configure(ctx, vars); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SubscriptionProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == SubscriptionToken)
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    if failure := p.ops.Check(ctx, obj, ""); failure != nil {
        failures = append(failures, failure)
    }
    if failure := p.ops.Check(ctx, obj, "service"); failure != nil {
        failures = append(failures,
            resource.NewPropertyError("Subscription", "service", failure))
    }
    if failure := p.ops.Check(ctx, obj, "config"); failure != nil {
        failures = append(failures,
            resource.NewPropertyError("Subscription", "config", failure))
    }
    if failure := p.ops.Check(ctx, obj, "events"); failure != nil {
        failures = append(failures,
            resource.NewPropertyError("Subscription", "events", failure))
    }
    if failure := p.ops.Check(ctx, obj, "active"); failure != nil {
        failures = append(failures,
            resource.NewPropertyError("Subscription", "active", failure))
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *SubscriptionProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == SubscriptionToken)
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{
        Id: string(id),
        Properties: plugin.MarshalProperties(
            resource.NewPropertyMap(obj), plugin.MarshalOptions{}),
    }, nil
}

func (p *SubscriptionProvider) Diff(
    ctx context.Context, req *lumirpc.DiffRequest) (*lumirpc.DiffResponse, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == SubscriptionToken)
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
    }
    more, err := p.ops.Diff(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.DiffResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *SubscriptionProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*lumirpc.UpdateResponse, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == SubscriptionToken)
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &lumirpc.UpdateResponse{
        Properties: plugin.MarshalProperties(
            resource.NewPropertyMap(new), plugin.MarshalOptions{}),
    }, nil
}

func (p *SubscriptionProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == SubscriptionToken)
    id := resource.ID(req.GetId())
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if err := p.ops.Delete(ctx, id, *obj); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SubscriptionProvider) Unmarshal(
    v *pbstruct.Struct) (*Subscription, resource.PropertyMap, error) {
    var obj Subscription
    props := plugin.UnmarshalProperties(v, plugin.MarshalOptions{})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable Subscription structure(s) */

// Subscription is a marshalable representation of its corresponding IDL type.
type Subscription struct {
    Service string `lumi:"service"`
    Config Config `lumi:"config"`
    Events *[]string `lumi:"events,optional"`
    Active *bool `lumi:"active,optional"`
}

// Subscription's properties have constants to make dealing with diffs and property bags easier.
const (
    Subscription_Service = "service"
    Subscription_Config = "config"
    Subscription_Events = "events"
    Subscription_Active = "active"
)

/* Typedefs */

type (
    ContentType string
)

/* Constants */

const (
    FormContentType ContentType = "form"
    JSONContentType ContentType = "json"
)


