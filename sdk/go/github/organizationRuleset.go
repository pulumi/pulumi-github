// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a GitHub organization ruleset.
//
// This resource allows you to create and manage rulesets on the organization level. When applied, a new ruleset will be created. When destroyed, that ruleset will be removed.
//
// ## Import
//
// GitHub Organization Rulesets can be imported using the GitHub ruleset ID e.g.
//
// ```sh
//
//	$ pulumi import github:index/organizationRuleset:OrganizationRuleset example 12345`
//
// ```
type OrganizationRuleset struct {
	pulumi.CustomResourceState

	// (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
	BypassActors OrganizationRulesetBypassActorArrayOutput `pulumi:"bypassActors"`
	// (Block List, Max: 1) Parameters for an organization ruleset condition. `refName` is required alongside one of `repositoryName` or `repositoryId`. (see below for nested schema)
	Conditions OrganizationRulesetConditionsPtrOutput `pulumi:"conditions"`
	// (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
	Enforcement pulumi.StringOutput `pulumi:"enforcement"`
	// (String)
	Etag pulumi.StringOutput `pulumi:"etag"`
	// (String) The name of the ruleset.
	Name pulumi.StringOutput `pulumi:"name"`
	// (String) GraphQL global node id for use with v4 API.
	NodeId pulumi.StringOutput `pulumi:"nodeId"`
	// (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
	Rules OrganizationRulesetRulesOutput `pulumi:"rules"`
	// (Number) GitHub ID for the ruleset.
	RulesetId pulumi.IntOutput `pulumi:"rulesetId"`
	// (String) Possible values are `branch` and `tag`.
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewOrganizationRuleset registers a new resource with the given unique name, arguments, and options.
func NewOrganizationRuleset(ctx *pulumi.Context,
	name string, args *OrganizationRulesetArgs, opts ...pulumi.ResourceOption) (*OrganizationRuleset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enforcement == nil {
		return nil, errors.New("invalid value for required argument 'Enforcement'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationRuleset
	err := ctx.RegisterResource("github:index/organizationRuleset:OrganizationRuleset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationRuleset gets an existing OrganizationRuleset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationRuleset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationRulesetState, opts ...pulumi.ResourceOption) (*OrganizationRuleset, error) {
	var resource OrganizationRuleset
	err := ctx.ReadResource("github:index/organizationRuleset:OrganizationRuleset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationRuleset resources.
type organizationRulesetState struct {
	// (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
	BypassActors []OrganizationRulesetBypassActor `pulumi:"bypassActors"`
	// (Block List, Max: 1) Parameters for an organization ruleset condition. `refName` is required alongside one of `repositoryName` or `repositoryId`. (see below for nested schema)
	Conditions *OrganizationRulesetConditions `pulumi:"conditions"`
	// (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
	Enforcement *string `pulumi:"enforcement"`
	// (String)
	Etag *string `pulumi:"etag"`
	// (String) The name of the ruleset.
	Name *string `pulumi:"name"`
	// (String) GraphQL global node id for use with v4 API.
	NodeId *string `pulumi:"nodeId"`
	// (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
	Rules *OrganizationRulesetRules `pulumi:"rules"`
	// (Number) GitHub ID for the ruleset.
	RulesetId *int `pulumi:"rulesetId"`
	// (String) Possible values are `branch` and `tag`.
	Target *string `pulumi:"target"`
}

type OrganizationRulesetState struct {
	// (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
	BypassActors OrganizationRulesetBypassActorArrayInput
	// (Block List, Max: 1) Parameters for an organization ruleset condition. `refName` is required alongside one of `repositoryName` or `repositoryId`. (see below for nested schema)
	Conditions OrganizationRulesetConditionsPtrInput
	// (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
	Enforcement pulumi.StringPtrInput
	// (String)
	Etag pulumi.StringPtrInput
	// (String) The name of the ruleset.
	Name pulumi.StringPtrInput
	// (String) GraphQL global node id for use with v4 API.
	NodeId pulumi.StringPtrInput
	// (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
	Rules OrganizationRulesetRulesPtrInput
	// (Number) GitHub ID for the ruleset.
	RulesetId pulumi.IntPtrInput
	// (String) Possible values are `branch` and `tag`.
	Target pulumi.StringPtrInput
}

func (OrganizationRulesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationRulesetState)(nil)).Elem()
}

type organizationRulesetArgs struct {
	// (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
	BypassActors []OrganizationRulesetBypassActor `pulumi:"bypassActors"`
	// (Block List, Max: 1) Parameters for an organization ruleset condition. `refName` is required alongside one of `repositoryName` or `repositoryId`. (see below for nested schema)
	Conditions *OrganizationRulesetConditions `pulumi:"conditions"`
	// (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
	Enforcement string `pulumi:"enforcement"`
	// (String) The name of the ruleset.
	Name *string `pulumi:"name"`
	// (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
	Rules OrganizationRulesetRules `pulumi:"rules"`
	// (String) Possible values are `branch` and `tag`.
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a OrganizationRuleset resource.
type OrganizationRulesetArgs struct {
	// (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
	BypassActors OrganizationRulesetBypassActorArrayInput
	// (Block List, Max: 1) Parameters for an organization ruleset condition. `refName` is required alongside one of `repositoryName` or `repositoryId`. (see below for nested schema)
	Conditions OrganizationRulesetConditionsPtrInput
	// (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
	Enforcement pulumi.StringInput
	// (String) The name of the ruleset.
	Name pulumi.StringPtrInput
	// (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
	Rules OrganizationRulesetRulesInput
	// (String) Possible values are `branch` and `tag`.
	Target pulumi.StringInput
}

func (OrganizationRulesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationRulesetArgs)(nil)).Elem()
}

type OrganizationRulesetInput interface {
	pulumi.Input

	ToOrganizationRulesetOutput() OrganizationRulesetOutput
	ToOrganizationRulesetOutputWithContext(ctx context.Context) OrganizationRulesetOutput
}

func (*OrganizationRuleset) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationRuleset)(nil)).Elem()
}

func (i *OrganizationRuleset) ToOrganizationRulesetOutput() OrganizationRulesetOutput {
	return i.ToOrganizationRulesetOutputWithContext(context.Background())
}

func (i *OrganizationRuleset) ToOrganizationRulesetOutputWithContext(ctx context.Context) OrganizationRulesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationRulesetOutput)
}

// OrganizationRulesetArrayInput is an input type that accepts OrganizationRulesetArray and OrganizationRulesetArrayOutput values.
// You can construct a concrete instance of `OrganizationRulesetArrayInput` via:
//
//	OrganizationRulesetArray{ OrganizationRulesetArgs{...} }
type OrganizationRulesetArrayInput interface {
	pulumi.Input

	ToOrganizationRulesetArrayOutput() OrganizationRulesetArrayOutput
	ToOrganizationRulesetArrayOutputWithContext(context.Context) OrganizationRulesetArrayOutput
}

type OrganizationRulesetArray []OrganizationRulesetInput

func (OrganizationRulesetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationRuleset)(nil)).Elem()
}

func (i OrganizationRulesetArray) ToOrganizationRulesetArrayOutput() OrganizationRulesetArrayOutput {
	return i.ToOrganizationRulesetArrayOutputWithContext(context.Background())
}

func (i OrganizationRulesetArray) ToOrganizationRulesetArrayOutputWithContext(ctx context.Context) OrganizationRulesetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationRulesetArrayOutput)
}

// OrganizationRulesetMapInput is an input type that accepts OrganizationRulesetMap and OrganizationRulesetMapOutput values.
// You can construct a concrete instance of `OrganizationRulesetMapInput` via:
//
//	OrganizationRulesetMap{ "key": OrganizationRulesetArgs{...} }
type OrganizationRulesetMapInput interface {
	pulumi.Input

	ToOrganizationRulesetMapOutput() OrganizationRulesetMapOutput
	ToOrganizationRulesetMapOutputWithContext(context.Context) OrganizationRulesetMapOutput
}

type OrganizationRulesetMap map[string]OrganizationRulesetInput

func (OrganizationRulesetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationRuleset)(nil)).Elem()
}

func (i OrganizationRulesetMap) ToOrganizationRulesetMapOutput() OrganizationRulesetMapOutput {
	return i.ToOrganizationRulesetMapOutputWithContext(context.Background())
}

func (i OrganizationRulesetMap) ToOrganizationRulesetMapOutputWithContext(ctx context.Context) OrganizationRulesetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationRulesetMapOutput)
}

type OrganizationRulesetOutput struct{ *pulumi.OutputState }

func (OrganizationRulesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationRuleset)(nil)).Elem()
}

func (o OrganizationRulesetOutput) ToOrganizationRulesetOutput() OrganizationRulesetOutput {
	return o
}

func (o OrganizationRulesetOutput) ToOrganizationRulesetOutputWithContext(ctx context.Context) OrganizationRulesetOutput {
	return o
}

// (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
func (o OrganizationRulesetOutput) BypassActors() OrganizationRulesetBypassActorArrayOutput {
	return o.ApplyT(func(v *OrganizationRuleset) OrganizationRulesetBypassActorArrayOutput { return v.BypassActors }).(OrganizationRulesetBypassActorArrayOutput)
}

// (Block List, Max: 1) Parameters for an organization ruleset condition. `refName` is required alongside one of `repositoryName` or `repositoryId`. (see below for nested schema)
func (o OrganizationRulesetOutput) Conditions() OrganizationRulesetConditionsPtrOutput {
	return o.ApplyT(func(v *OrganizationRuleset) OrganizationRulesetConditionsPtrOutput { return v.Conditions }).(OrganizationRulesetConditionsPtrOutput)
}

// (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
func (o OrganizationRulesetOutput) Enforcement() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationRuleset) pulumi.StringOutput { return v.Enforcement }).(pulumi.StringOutput)
}

// (String)
func (o OrganizationRulesetOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationRuleset) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// (String) The name of the ruleset.
func (o OrganizationRulesetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationRuleset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (String) GraphQL global node id for use with v4 API.
func (o OrganizationRulesetOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationRuleset) pulumi.StringOutput { return v.NodeId }).(pulumi.StringOutput)
}

// (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
func (o OrganizationRulesetOutput) Rules() OrganizationRulesetRulesOutput {
	return o.ApplyT(func(v *OrganizationRuleset) OrganizationRulesetRulesOutput { return v.Rules }).(OrganizationRulesetRulesOutput)
}

// (Number) GitHub ID for the ruleset.
func (o OrganizationRulesetOutput) RulesetId() pulumi.IntOutput {
	return o.ApplyT(func(v *OrganizationRuleset) pulumi.IntOutput { return v.RulesetId }).(pulumi.IntOutput)
}

// (String) Possible values are `branch` and `tag`.
func (o OrganizationRulesetOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationRuleset) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

type OrganizationRulesetArrayOutput struct{ *pulumi.OutputState }

func (OrganizationRulesetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationRuleset)(nil)).Elem()
}

func (o OrganizationRulesetArrayOutput) ToOrganizationRulesetArrayOutput() OrganizationRulesetArrayOutput {
	return o
}

func (o OrganizationRulesetArrayOutput) ToOrganizationRulesetArrayOutputWithContext(ctx context.Context) OrganizationRulesetArrayOutput {
	return o
}

func (o OrganizationRulesetArrayOutput) Index(i pulumi.IntInput) OrganizationRulesetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationRuleset {
		return vs[0].([]*OrganizationRuleset)[vs[1].(int)]
	}).(OrganizationRulesetOutput)
}

type OrganizationRulesetMapOutput struct{ *pulumi.OutputState }

func (OrganizationRulesetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationRuleset)(nil)).Elem()
}

func (o OrganizationRulesetMapOutput) ToOrganizationRulesetMapOutput() OrganizationRulesetMapOutput {
	return o
}

func (o OrganizationRulesetMapOutput) ToOrganizationRulesetMapOutputWithContext(ctx context.Context) OrganizationRulesetMapOutput {
	return o
}

func (o OrganizationRulesetMapOutput) MapIndex(k pulumi.StringInput) OrganizationRulesetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationRuleset {
		return vs[0].(map[string]*OrganizationRuleset)[vs[1].(string)]
	}).(OrganizationRulesetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationRulesetInput)(nil)).Elem(), &OrganizationRuleset{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationRulesetArrayInput)(nil)).Elem(), OrganizationRulesetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationRulesetMapInput)(nil)).Elem(), OrganizationRulesetMap{})
	pulumi.RegisterOutputType(OrganizationRulesetOutput{})
	pulumi.RegisterOutputType(OrganizationRulesetArrayOutput{})
	pulumi.RegisterOutputType(OrganizationRulesetMapOutput{})
}
