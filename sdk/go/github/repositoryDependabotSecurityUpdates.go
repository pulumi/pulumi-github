// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to manage dependabot automated security fixes for a single repository. See the
// [documentation](https://docs.github.com/en/code-security/dependabot/dependabot-security-updates/about-dependabot-security-updates)
// for details of usage and how this will impact your repository
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := github.NewRepository(ctx, "repo", &github.RepositoryArgs{
//				Name:                pulumi.String("my-repo"),
//				Description:         pulumi.String("GitHub repo managed by Terraform"),
//				Private:             pulumi.Bool(false),
//				VulnerabilityAlerts: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryDependabotSecurityUpdates(ctx, "example", &github.RepositoryDependabotSecurityUpdatesArgs{
//				Repository: pulumi.Any(test.Id),
//				Enabled:    pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ### Import by name
//
// ```sh
// $ pulumi import github:index/repositoryDependabotSecurityUpdates:RepositoryDependabotSecurityUpdates example my-repo
// ```
type RepositoryDependabotSecurityUpdates struct {
	pulumi.CustomResourceState

	// The state of the automated security fixes.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The repository to manage.
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewRepositoryDependabotSecurityUpdates registers a new resource with the given unique name, arguments, and options.
func NewRepositoryDependabotSecurityUpdates(ctx *pulumi.Context,
	name string, args *RepositoryDependabotSecurityUpdatesArgs, opts ...pulumi.ResourceOption) (*RepositoryDependabotSecurityUpdates, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryDependabotSecurityUpdates
	err := ctx.RegisterResource("github:index/repositoryDependabotSecurityUpdates:RepositoryDependabotSecurityUpdates", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryDependabotSecurityUpdates gets an existing RepositoryDependabotSecurityUpdates resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryDependabotSecurityUpdates(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryDependabotSecurityUpdatesState, opts ...pulumi.ResourceOption) (*RepositoryDependabotSecurityUpdates, error) {
	var resource RepositoryDependabotSecurityUpdates
	err := ctx.ReadResource("github:index/repositoryDependabotSecurityUpdates:RepositoryDependabotSecurityUpdates", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryDependabotSecurityUpdates resources.
type repositoryDependabotSecurityUpdatesState struct {
	// The state of the automated security fixes.
	Enabled *bool `pulumi:"enabled"`
	// The repository to manage.
	Repository *string `pulumi:"repository"`
}

type RepositoryDependabotSecurityUpdatesState struct {
	// The state of the automated security fixes.
	Enabled pulumi.BoolPtrInput
	// The repository to manage.
	Repository pulumi.StringPtrInput
}

func (RepositoryDependabotSecurityUpdatesState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryDependabotSecurityUpdatesState)(nil)).Elem()
}

type repositoryDependabotSecurityUpdatesArgs struct {
	// The state of the automated security fixes.
	Enabled bool `pulumi:"enabled"`
	// The repository to manage.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryDependabotSecurityUpdates resource.
type RepositoryDependabotSecurityUpdatesArgs struct {
	// The state of the automated security fixes.
	Enabled pulumi.BoolInput
	// The repository to manage.
	Repository pulumi.StringInput
}

func (RepositoryDependabotSecurityUpdatesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryDependabotSecurityUpdatesArgs)(nil)).Elem()
}

type RepositoryDependabotSecurityUpdatesInput interface {
	pulumi.Input

	ToRepositoryDependabotSecurityUpdatesOutput() RepositoryDependabotSecurityUpdatesOutput
	ToRepositoryDependabotSecurityUpdatesOutputWithContext(ctx context.Context) RepositoryDependabotSecurityUpdatesOutput
}

func (*RepositoryDependabotSecurityUpdates) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryDependabotSecurityUpdates)(nil)).Elem()
}

func (i *RepositoryDependabotSecurityUpdates) ToRepositoryDependabotSecurityUpdatesOutput() RepositoryDependabotSecurityUpdatesOutput {
	return i.ToRepositoryDependabotSecurityUpdatesOutputWithContext(context.Background())
}

func (i *RepositoryDependabotSecurityUpdates) ToRepositoryDependabotSecurityUpdatesOutputWithContext(ctx context.Context) RepositoryDependabotSecurityUpdatesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryDependabotSecurityUpdatesOutput)
}

// RepositoryDependabotSecurityUpdatesArrayInput is an input type that accepts RepositoryDependabotSecurityUpdatesArray and RepositoryDependabotSecurityUpdatesArrayOutput values.
// You can construct a concrete instance of `RepositoryDependabotSecurityUpdatesArrayInput` via:
//
//	RepositoryDependabotSecurityUpdatesArray{ RepositoryDependabotSecurityUpdatesArgs{...} }
type RepositoryDependabotSecurityUpdatesArrayInput interface {
	pulumi.Input

	ToRepositoryDependabotSecurityUpdatesArrayOutput() RepositoryDependabotSecurityUpdatesArrayOutput
	ToRepositoryDependabotSecurityUpdatesArrayOutputWithContext(context.Context) RepositoryDependabotSecurityUpdatesArrayOutput
}

type RepositoryDependabotSecurityUpdatesArray []RepositoryDependabotSecurityUpdatesInput

func (RepositoryDependabotSecurityUpdatesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryDependabotSecurityUpdates)(nil)).Elem()
}

func (i RepositoryDependabotSecurityUpdatesArray) ToRepositoryDependabotSecurityUpdatesArrayOutput() RepositoryDependabotSecurityUpdatesArrayOutput {
	return i.ToRepositoryDependabotSecurityUpdatesArrayOutputWithContext(context.Background())
}

func (i RepositoryDependabotSecurityUpdatesArray) ToRepositoryDependabotSecurityUpdatesArrayOutputWithContext(ctx context.Context) RepositoryDependabotSecurityUpdatesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryDependabotSecurityUpdatesArrayOutput)
}

// RepositoryDependabotSecurityUpdatesMapInput is an input type that accepts RepositoryDependabotSecurityUpdatesMap and RepositoryDependabotSecurityUpdatesMapOutput values.
// You can construct a concrete instance of `RepositoryDependabotSecurityUpdatesMapInput` via:
//
//	RepositoryDependabotSecurityUpdatesMap{ "key": RepositoryDependabotSecurityUpdatesArgs{...} }
type RepositoryDependabotSecurityUpdatesMapInput interface {
	pulumi.Input

	ToRepositoryDependabotSecurityUpdatesMapOutput() RepositoryDependabotSecurityUpdatesMapOutput
	ToRepositoryDependabotSecurityUpdatesMapOutputWithContext(context.Context) RepositoryDependabotSecurityUpdatesMapOutput
}

type RepositoryDependabotSecurityUpdatesMap map[string]RepositoryDependabotSecurityUpdatesInput

func (RepositoryDependabotSecurityUpdatesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryDependabotSecurityUpdates)(nil)).Elem()
}

func (i RepositoryDependabotSecurityUpdatesMap) ToRepositoryDependabotSecurityUpdatesMapOutput() RepositoryDependabotSecurityUpdatesMapOutput {
	return i.ToRepositoryDependabotSecurityUpdatesMapOutputWithContext(context.Background())
}

func (i RepositoryDependabotSecurityUpdatesMap) ToRepositoryDependabotSecurityUpdatesMapOutputWithContext(ctx context.Context) RepositoryDependabotSecurityUpdatesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryDependabotSecurityUpdatesMapOutput)
}

type RepositoryDependabotSecurityUpdatesOutput struct{ *pulumi.OutputState }

func (RepositoryDependabotSecurityUpdatesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryDependabotSecurityUpdates)(nil)).Elem()
}

func (o RepositoryDependabotSecurityUpdatesOutput) ToRepositoryDependabotSecurityUpdatesOutput() RepositoryDependabotSecurityUpdatesOutput {
	return o
}

func (o RepositoryDependabotSecurityUpdatesOutput) ToRepositoryDependabotSecurityUpdatesOutputWithContext(ctx context.Context) RepositoryDependabotSecurityUpdatesOutput {
	return o
}

// The state of the automated security fixes.
func (o RepositoryDependabotSecurityUpdatesOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *RepositoryDependabotSecurityUpdates) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The repository to manage.
func (o RepositoryDependabotSecurityUpdatesOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryDependabotSecurityUpdates) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

type RepositoryDependabotSecurityUpdatesArrayOutput struct{ *pulumi.OutputState }

func (RepositoryDependabotSecurityUpdatesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryDependabotSecurityUpdates)(nil)).Elem()
}

func (o RepositoryDependabotSecurityUpdatesArrayOutput) ToRepositoryDependabotSecurityUpdatesArrayOutput() RepositoryDependabotSecurityUpdatesArrayOutput {
	return o
}

func (o RepositoryDependabotSecurityUpdatesArrayOutput) ToRepositoryDependabotSecurityUpdatesArrayOutputWithContext(ctx context.Context) RepositoryDependabotSecurityUpdatesArrayOutput {
	return o
}

func (o RepositoryDependabotSecurityUpdatesArrayOutput) Index(i pulumi.IntInput) RepositoryDependabotSecurityUpdatesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryDependabotSecurityUpdates {
		return vs[0].([]*RepositoryDependabotSecurityUpdates)[vs[1].(int)]
	}).(RepositoryDependabotSecurityUpdatesOutput)
}

type RepositoryDependabotSecurityUpdatesMapOutput struct{ *pulumi.OutputState }

func (RepositoryDependabotSecurityUpdatesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryDependabotSecurityUpdates)(nil)).Elem()
}

func (o RepositoryDependabotSecurityUpdatesMapOutput) ToRepositoryDependabotSecurityUpdatesMapOutput() RepositoryDependabotSecurityUpdatesMapOutput {
	return o
}

func (o RepositoryDependabotSecurityUpdatesMapOutput) ToRepositoryDependabotSecurityUpdatesMapOutputWithContext(ctx context.Context) RepositoryDependabotSecurityUpdatesMapOutput {
	return o
}

func (o RepositoryDependabotSecurityUpdatesMapOutput) MapIndex(k pulumi.StringInput) RepositoryDependabotSecurityUpdatesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryDependabotSecurityUpdates {
		return vs[0].(map[string]*RepositoryDependabotSecurityUpdates)[vs[1].(string)]
	}).(RepositoryDependabotSecurityUpdatesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryDependabotSecurityUpdatesInput)(nil)).Elem(), &RepositoryDependabotSecurityUpdates{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryDependabotSecurityUpdatesArrayInput)(nil)).Elem(), RepositoryDependabotSecurityUpdatesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryDependabotSecurityUpdatesMapInput)(nil)).Elem(), RepositoryDependabotSecurityUpdatesMap{})
	pulumi.RegisterOutputType(RepositoryDependabotSecurityUpdatesOutput{})
	pulumi.RegisterOutputType(RepositoryDependabotSecurityUpdatesArrayOutput{})
	pulumi.RegisterOutputType(RepositoryDependabotSecurityUpdatesMapOutput{})
}
