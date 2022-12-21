// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TeamMembership struct {
	pulumi.CustomResourceState

	Etag     pulumi.StringOutput    `pulumi:"etag"`
	Role     pulumi.StringPtrOutput `pulumi:"role"`
	TeamId   pulumi.StringOutput    `pulumi:"teamId"`
	Username pulumi.StringOutput    `pulumi:"username"`
}

// NewTeamMembership registers a new resource with the given unique name, arguments, and options.
func NewTeamMembership(ctx *pulumi.Context,
	name string, args *TeamMembershipArgs, opts ...pulumi.ResourceOption) (*TeamMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource TeamMembership
	err := ctx.RegisterResource("github:index/teamMembership:TeamMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamMembership gets an existing TeamMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamMembershipState, opts ...pulumi.ResourceOption) (*TeamMembership, error) {
	var resource TeamMembership
	err := ctx.ReadResource("github:index/teamMembership:TeamMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamMembership resources.
type teamMembershipState struct {
	Etag     *string `pulumi:"etag"`
	Role     *string `pulumi:"role"`
	TeamId   *string `pulumi:"teamId"`
	Username *string `pulumi:"username"`
}

type TeamMembershipState struct {
	Etag     pulumi.StringPtrInput
	Role     pulumi.StringPtrInput
	TeamId   pulumi.StringPtrInput
	Username pulumi.StringPtrInput
}

func (TeamMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamMembershipState)(nil)).Elem()
}

type teamMembershipArgs struct {
	Role     *string `pulumi:"role"`
	TeamId   string  `pulumi:"teamId"`
	Username string  `pulumi:"username"`
}

// The set of arguments for constructing a TeamMembership resource.
type TeamMembershipArgs struct {
	Role     pulumi.StringPtrInput
	TeamId   pulumi.StringInput
	Username pulumi.StringInput
}

func (TeamMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamMembershipArgs)(nil)).Elem()
}

type TeamMembershipInput interface {
	pulumi.Input

	ToTeamMembershipOutput() TeamMembershipOutput
	ToTeamMembershipOutputWithContext(ctx context.Context) TeamMembershipOutput
}

func (*TeamMembership) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamMembership)(nil)).Elem()
}

func (i *TeamMembership) ToTeamMembershipOutput() TeamMembershipOutput {
	return i.ToTeamMembershipOutputWithContext(context.Background())
}

func (i *TeamMembership) ToTeamMembershipOutputWithContext(ctx context.Context) TeamMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembershipOutput)
}

// TeamMembershipArrayInput is an input type that accepts TeamMembershipArray and TeamMembershipArrayOutput values.
// You can construct a concrete instance of `TeamMembershipArrayInput` via:
//
//          TeamMembershipArray{ TeamMembershipArgs{...} }
type TeamMembershipArrayInput interface {
	pulumi.Input

	ToTeamMembershipArrayOutput() TeamMembershipArrayOutput
	ToTeamMembershipArrayOutputWithContext(context.Context) TeamMembershipArrayOutput
}

type TeamMembershipArray []TeamMembershipInput

func (TeamMembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamMembership)(nil)).Elem()
}

func (i TeamMembershipArray) ToTeamMembershipArrayOutput() TeamMembershipArrayOutput {
	return i.ToTeamMembershipArrayOutputWithContext(context.Background())
}

func (i TeamMembershipArray) ToTeamMembershipArrayOutputWithContext(ctx context.Context) TeamMembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembershipArrayOutput)
}

// TeamMembershipMapInput is an input type that accepts TeamMembershipMap and TeamMembershipMapOutput values.
// You can construct a concrete instance of `TeamMembershipMapInput` via:
//
//          TeamMembershipMap{ "key": TeamMembershipArgs{...} }
type TeamMembershipMapInput interface {
	pulumi.Input

	ToTeamMembershipMapOutput() TeamMembershipMapOutput
	ToTeamMembershipMapOutputWithContext(context.Context) TeamMembershipMapOutput
}

type TeamMembershipMap map[string]TeamMembershipInput

func (TeamMembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamMembership)(nil)).Elem()
}

func (i TeamMembershipMap) ToTeamMembershipMapOutput() TeamMembershipMapOutput {
	return i.ToTeamMembershipMapOutputWithContext(context.Background())
}

func (i TeamMembershipMap) ToTeamMembershipMapOutputWithContext(ctx context.Context) TeamMembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembershipMapOutput)
}

type TeamMembershipOutput struct{ *pulumi.OutputState }

func (TeamMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamMembership)(nil)).Elem()
}

func (o TeamMembershipOutput) ToTeamMembershipOutput() TeamMembershipOutput {
	return o
}

func (o TeamMembershipOutput) ToTeamMembershipOutputWithContext(ctx context.Context) TeamMembershipOutput {
	return o
}

func (o TeamMembershipOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamMembership) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o TeamMembershipOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TeamMembership) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

func (o TeamMembershipOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamMembership) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

func (o TeamMembershipOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamMembership) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type TeamMembershipArrayOutput struct{ *pulumi.OutputState }

func (TeamMembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamMembership)(nil)).Elem()
}

func (o TeamMembershipArrayOutput) ToTeamMembershipArrayOutput() TeamMembershipArrayOutput {
	return o
}

func (o TeamMembershipArrayOutput) ToTeamMembershipArrayOutputWithContext(ctx context.Context) TeamMembershipArrayOutput {
	return o
}

func (o TeamMembershipArrayOutput) Index(i pulumi.IntInput) TeamMembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TeamMembership {
		return vs[0].([]*TeamMembership)[vs[1].(int)]
	}).(TeamMembershipOutput)
}

type TeamMembershipMapOutput struct{ *pulumi.OutputState }

func (TeamMembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamMembership)(nil)).Elem()
}

func (o TeamMembershipMapOutput) ToTeamMembershipMapOutput() TeamMembershipMapOutput {
	return o
}

func (o TeamMembershipMapOutput) ToTeamMembershipMapOutputWithContext(ctx context.Context) TeamMembershipMapOutput {
	return o
}

func (o TeamMembershipMapOutput) MapIndex(k pulumi.StringInput) TeamMembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TeamMembership {
		return vs[0].(map[string]*TeamMembership)[vs[1].(string)]
	}).(TeamMembershipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembershipInput)(nil)).Elem(), &TeamMembership{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembershipArrayInput)(nil)).Elem(), TeamMembershipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembershipMapInput)(nil)).Elem(), TeamMembershipMap{})
	pulumi.RegisterOutputType(TeamMembershipOutput{})
	pulumi.RegisterOutputType(TeamMembershipArrayOutput{})
	pulumi.RegisterOutputType(TeamMembershipMapOutput{})
}
