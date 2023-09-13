// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource allows you to create and manage projects for GitHub repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := github.NewRepository(ctx, "example", &github.RepositoryArgs{
//				Description: pulumi.String("My awesome codebase"),
//				HasProjects: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryProject(ctx, "project", &github.RepositoryProjectArgs{
//				Body:       pulumi.String("This is a repository project."),
//				Repository: example.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type RepositoryProject struct {
	pulumi.CustomResourceState

	// The body of the project.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	Etag pulumi.StringOutput    `pulumi:"etag"`
	// The name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The repository of the project.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// URL of the project
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewRepositoryProject registers a new resource with the given unique name, arguments, and options.
func NewRepositoryProject(ctx *pulumi.Context,
	name string, args *RepositoryProjectArgs, opts ...pulumi.ResourceOption) (*RepositoryProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryProject
	err := ctx.RegisterResource("github:index/repositoryProject:RepositoryProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryProject gets an existing RepositoryProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryProjectState, opts ...pulumi.ResourceOption) (*RepositoryProject, error) {
	var resource RepositoryProject
	err := ctx.ReadResource("github:index/repositoryProject:RepositoryProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryProject resources.
type repositoryProjectState struct {
	// The body of the project.
	Body *string `pulumi:"body"`
	Etag *string `pulumi:"etag"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The repository of the project.
	Repository *string `pulumi:"repository"`
	// URL of the project
	Url *string `pulumi:"url"`
}

type RepositoryProjectState struct {
	// The body of the project.
	Body pulumi.StringPtrInput
	Etag pulumi.StringPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The repository of the project.
	Repository pulumi.StringPtrInput
	// URL of the project
	Url pulumi.StringPtrInput
}

func (RepositoryProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryProjectState)(nil)).Elem()
}

type repositoryProjectArgs struct {
	// The body of the project.
	Body *string `pulumi:"body"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The repository of the project.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryProject resource.
type RepositoryProjectArgs struct {
	// The body of the project.
	Body pulumi.StringPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The repository of the project.
	Repository pulumi.StringInput
}

func (RepositoryProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryProjectArgs)(nil)).Elem()
}

type RepositoryProjectInput interface {
	pulumi.Input

	ToRepositoryProjectOutput() RepositoryProjectOutput
	ToRepositoryProjectOutputWithContext(ctx context.Context) RepositoryProjectOutput
}

func (*RepositoryProject) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryProject)(nil)).Elem()
}

func (i *RepositoryProject) ToRepositoryProjectOutput() RepositoryProjectOutput {
	return i.ToRepositoryProjectOutputWithContext(context.Background())
}

func (i *RepositoryProject) ToRepositoryProjectOutputWithContext(ctx context.Context) RepositoryProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryProjectOutput)
}

func (i *RepositoryProject) ToOutput(ctx context.Context) pulumix.Output[*RepositoryProject] {
	return pulumix.Output[*RepositoryProject]{
		OutputState: i.ToRepositoryProjectOutputWithContext(ctx).OutputState,
	}
}

// RepositoryProjectArrayInput is an input type that accepts RepositoryProjectArray and RepositoryProjectArrayOutput values.
// You can construct a concrete instance of `RepositoryProjectArrayInput` via:
//
//	RepositoryProjectArray{ RepositoryProjectArgs{...} }
type RepositoryProjectArrayInput interface {
	pulumi.Input

	ToRepositoryProjectArrayOutput() RepositoryProjectArrayOutput
	ToRepositoryProjectArrayOutputWithContext(context.Context) RepositoryProjectArrayOutput
}

type RepositoryProjectArray []RepositoryProjectInput

func (RepositoryProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryProject)(nil)).Elem()
}

func (i RepositoryProjectArray) ToRepositoryProjectArrayOutput() RepositoryProjectArrayOutput {
	return i.ToRepositoryProjectArrayOutputWithContext(context.Background())
}

func (i RepositoryProjectArray) ToRepositoryProjectArrayOutputWithContext(ctx context.Context) RepositoryProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryProjectArrayOutput)
}

func (i RepositoryProjectArray) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryProject] {
	return pulumix.Output[[]*RepositoryProject]{
		OutputState: i.ToRepositoryProjectArrayOutputWithContext(ctx).OutputState,
	}
}

// RepositoryProjectMapInput is an input type that accepts RepositoryProjectMap and RepositoryProjectMapOutput values.
// You can construct a concrete instance of `RepositoryProjectMapInput` via:
//
//	RepositoryProjectMap{ "key": RepositoryProjectArgs{...} }
type RepositoryProjectMapInput interface {
	pulumi.Input

	ToRepositoryProjectMapOutput() RepositoryProjectMapOutput
	ToRepositoryProjectMapOutputWithContext(context.Context) RepositoryProjectMapOutput
}

type RepositoryProjectMap map[string]RepositoryProjectInput

func (RepositoryProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryProject)(nil)).Elem()
}

func (i RepositoryProjectMap) ToRepositoryProjectMapOutput() RepositoryProjectMapOutput {
	return i.ToRepositoryProjectMapOutputWithContext(context.Background())
}

func (i RepositoryProjectMap) ToRepositoryProjectMapOutputWithContext(ctx context.Context) RepositoryProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryProjectMapOutput)
}

func (i RepositoryProjectMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryProject] {
	return pulumix.Output[map[string]*RepositoryProject]{
		OutputState: i.ToRepositoryProjectMapOutputWithContext(ctx).OutputState,
	}
}

type RepositoryProjectOutput struct{ *pulumi.OutputState }

func (RepositoryProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryProject)(nil)).Elem()
}

func (o RepositoryProjectOutput) ToRepositoryProjectOutput() RepositoryProjectOutput {
	return o
}

func (o RepositoryProjectOutput) ToRepositoryProjectOutputWithContext(ctx context.Context) RepositoryProjectOutput {
	return o
}

func (o RepositoryProjectOutput) ToOutput(ctx context.Context) pulumix.Output[*RepositoryProject] {
	return pulumix.Output[*RepositoryProject]{
		OutputState: o.OutputState,
	}
}

// The body of the project.
func (o RepositoryProjectOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryProject) pulumi.StringPtrOutput { return v.Body }).(pulumi.StringPtrOutput)
}

func (o RepositoryProjectOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryProject) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the project.
func (o RepositoryProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryProject) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The repository of the project.
func (o RepositoryProjectOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryProject) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// URL of the project
func (o RepositoryProjectOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryProject) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type RepositoryProjectArrayOutput struct{ *pulumi.OutputState }

func (RepositoryProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryProject)(nil)).Elem()
}

func (o RepositoryProjectArrayOutput) ToRepositoryProjectArrayOutput() RepositoryProjectArrayOutput {
	return o
}

func (o RepositoryProjectArrayOutput) ToRepositoryProjectArrayOutputWithContext(ctx context.Context) RepositoryProjectArrayOutput {
	return o
}

func (o RepositoryProjectArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryProject] {
	return pulumix.Output[[]*RepositoryProject]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryProjectArrayOutput) Index(i pulumi.IntInput) RepositoryProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryProject {
		return vs[0].([]*RepositoryProject)[vs[1].(int)]
	}).(RepositoryProjectOutput)
}

type RepositoryProjectMapOutput struct{ *pulumi.OutputState }

func (RepositoryProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryProject)(nil)).Elem()
}

func (o RepositoryProjectMapOutput) ToRepositoryProjectMapOutput() RepositoryProjectMapOutput {
	return o
}

func (o RepositoryProjectMapOutput) ToRepositoryProjectMapOutputWithContext(ctx context.Context) RepositoryProjectMapOutput {
	return o
}

func (o RepositoryProjectMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryProject] {
	return pulumix.Output[map[string]*RepositoryProject]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryProjectMapOutput) MapIndex(k pulumi.StringInput) RepositoryProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryProject {
		return vs[0].(map[string]*RepositoryProject)[vs[1].(string)]
	}).(RepositoryProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryProjectInput)(nil)).Elem(), &RepositoryProject{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryProjectArrayInput)(nil)).Elem(), RepositoryProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryProjectMapInput)(nil)).Elem(), RepositoryProjectMap{})
	pulumi.RegisterOutputType(RepositoryProjectOutput{})
	pulumi.RegisterOutputType(RepositoryProjectArrayOutput{})
	pulumi.RegisterOutputType(RepositoryProjectMapOutput{})
}
