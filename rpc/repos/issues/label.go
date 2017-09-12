// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package issues

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

/* RPC stubs for Label resource provider */

// LabelToken is the type token corresponding to the Label package type.
const LabelToken = tokens.Type("github:repos/issues/label:Label")

// LabelProviderOps is a pluggable interface for Label-related management functionality.
type LabelProviderOps interface {
    Configure(ctx context.Context, vars map[tokens.ModuleMember]string) error
    Check(ctx context.Context, obj *Label, property string) error
    Diff(ctx context.Context, id resource.ID,
        old *Label, new *Label, diff *resource.ObjectDiff) ([]string, error)
    Create(ctx context.Context, obj *Label) (resource.ID, error)
    Update(ctx context.Context, id resource.ID,
        old *Label, new *Label, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID, obj Label) error
}

// LabelProvider is a dynamic gRPC-based plugin for managing Label resources.
type LabelProvider struct {
    ops LabelProviderOps
}

// NewLabelProvider allocates a resource provider that delegates to a ops instance.
func NewLabelProvider(ops LabelProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &LabelProvider{ops: ops}
}

func (p *LabelProvider) Configure(
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

func (p *LabelProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == LabelToken)
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    if failure := p.ops.Check(ctx, obj, ""); failure != nil {
        failures = append(failures, failure)
    }
    if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
        failures = append(failures,
            resource.NewPropertyError("Label", "name", failure))
    }
    if failure := p.ops.Check(ctx, obj, "color"); failure != nil {
        failures = append(failures,
            resource.NewPropertyError("Label", "color", failure))
    }
    if failure := p.ops.Check(ctx, obj, "repo"); failure != nil {
        failures = append(failures,
            resource.NewPropertyError("Label", "repo", failure))
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *LabelProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == LabelToken)
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

func (p *LabelProvider) Diff(
    ctx context.Context, req *lumirpc.DiffRequest) (*lumirpc.DiffResponse, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == LabelToken)
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
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("repo") {
            replaces = append(replaces, "repo")
        }
    }
    more, err := p.ops.Diff(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.DiffResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *LabelProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*lumirpc.UpdateResponse, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == LabelToken)
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

func (p *LabelProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(resource.URN(req.GetUrn()).Type() == LabelToken)
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

func (p *LabelProvider) Unmarshal(
    v *pbstruct.Struct) (*Label, resource.PropertyMap, error) {
    var obj Label
    props := plugin.UnmarshalProperties(v, plugin.MarshalOptions{})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable Label structure(s) */

// Label is a marshalable representation of its corresponding IDL type.
type Label struct {
    Name string `lumi:"name"`
    Color string `lumi:"color"`
    Repo *string `lumi:"repo,optional"`
}

// Label's properties have constants to make dealing with diffs and property bags easier.
const (
    Label_Name = "name"
    Label_Color = "color"
    Label_Repo = "repo"
)


