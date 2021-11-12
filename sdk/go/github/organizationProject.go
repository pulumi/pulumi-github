// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage projects for GitHub organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.NewOrganizationProject(ctx, "project", &github.OrganizationProjectArgs{
// 			Body: pulumi.String("This is a organization project."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OrganizationProject struct {
	pulumi.CustomResourceState

	// The body of the project.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	Etag pulumi.StringOutput    `pulumi:"etag"`
	// The name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the project
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewOrganizationProject registers a new resource with the given unique name, arguments, and options.
func NewOrganizationProject(ctx *pulumi.Context,
	name string, args *OrganizationProjectArgs, opts ...pulumi.ResourceOption) (*OrganizationProject, error) {
	if args == nil {
		args = &OrganizationProjectArgs{}
	}

	var resource OrganizationProject
	err := ctx.RegisterResource("github:index/organizationProject:OrganizationProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationProject gets an existing OrganizationProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationProjectState, opts ...pulumi.ResourceOption) (*OrganizationProject, error) {
	var resource OrganizationProject
	err := ctx.ReadResource("github:index/organizationProject:OrganizationProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationProject resources.
type organizationProjectState struct {
	// The body of the project.
	Body *string `pulumi:"body"`
	Etag *string `pulumi:"etag"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// URL of the project
	Url *string `pulumi:"url"`
}

type OrganizationProjectState struct {
	// The body of the project.
	Body pulumi.StringPtrInput
	Etag pulumi.StringPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// URL of the project
	Url pulumi.StringPtrInput
}

func (OrganizationProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationProjectState)(nil)).Elem()
}

type organizationProjectArgs struct {
	// The body of the project.
	Body *string `pulumi:"body"`
	// The name of the project.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a OrganizationProject resource.
type OrganizationProjectArgs struct {
	// The body of the project.
	Body pulumi.StringPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
}

func (OrganizationProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationProjectArgs)(nil)).Elem()
}

type OrganizationProjectInput interface {
	pulumi.Input

	ToOrganizationProjectOutput() OrganizationProjectOutput
	ToOrganizationProjectOutputWithContext(ctx context.Context) OrganizationProjectOutput
}

func (*OrganizationProject) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationProject)(nil)).Elem()
}

func (i *OrganizationProject) ToOrganizationProjectOutput() OrganizationProjectOutput {
	return i.ToOrganizationProjectOutputWithContext(context.Background())
}

func (i *OrganizationProject) ToOrganizationProjectOutputWithContext(ctx context.Context) OrganizationProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationProjectOutput)
}

// OrganizationProjectArrayInput is an input type that accepts OrganizationProjectArray and OrganizationProjectArrayOutput values.
// You can construct a concrete instance of `OrganizationProjectArrayInput` via:
//
//          OrganizationProjectArray{ OrganizationProjectArgs{...} }
type OrganizationProjectArrayInput interface {
	pulumi.Input

	ToOrganizationProjectArrayOutput() OrganizationProjectArrayOutput
	ToOrganizationProjectArrayOutputWithContext(context.Context) OrganizationProjectArrayOutput
}

type OrganizationProjectArray []OrganizationProjectInput

func (OrganizationProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationProject)(nil)).Elem()
}

func (i OrganizationProjectArray) ToOrganizationProjectArrayOutput() OrganizationProjectArrayOutput {
	return i.ToOrganizationProjectArrayOutputWithContext(context.Background())
}

func (i OrganizationProjectArray) ToOrganizationProjectArrayOutputWithContext(ctx context.Context) OrganizationProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationProjectArrayOutput)
}

// OrganizationProjectMapInput is an input type that accepts OrganizationProjectMap and OrganizationProjectMapOutput values.
// You can construct a concrete instance of `OrganizationProjectMapInput` via:
//
//          OrganizationProjectMap{ "key": OrganizationProjectArgs{...} }
type OrganizationProjectMapInput interface {
	pulumi.Input

	ToOrganizationProjectMapOutput() OrganizationProjectMapOutput
	ToOrganizationProjectMapOutputWithContext(context.Context) OrganizationProjectMapOutput
}

type OrganizationProjectMap map[string]OrganizationProjectInput

func (OrganizationProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationProject)(nil)).Elem()
}

func (i OrganizationProjectMap) ToOrganizationProjectMapOutput() OrganizationProjectMapOutput {
	return i.ToOrganizationProjectMapOutputWithContext(context.Background())
}

func (i OrganizationProjectMap) ToOrganizationProjectMapOutputWithContext(ctx context.Context) OrganizationProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationProjectMapOutput)
}

type OrganizationProjectOutput struct{ *pulumi.OutputState }

func (OrganizationProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationProject)(nil)).Elem()
}

func (o OrganizationProjectOutput) ToOrganizationProjectOutput() OrganizationProjectOutput {
	return o
}

func (o OrganizationProjectOutput) ToOrganizationProjectOutputWithContext(ctx context.Context) OrganizationProjectOutput {
	return o
}

type OrganizationProjectArrayOutput struct{ *pulumi.OutputState }

func (OrganizationProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationProject)(nil)).Elem()
}

func (o OrganizationProjectArrayOutput) ToOrganizationProjectArrayOutput() OrganizationProjectArrayOutput {
	return o
}

func (o OrganizationProjectArrayOutput) ToOrganizationProjectArrayOutputWithContext(ctx context.Context) OrganizationProjectArrayOutput {
	return o
}

func (o OrganizationProjectArrayOutput) Index(i pulumi.IntInput) OrganizationProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationProject {
		return vs[0].([]*OrganizationProject)[vs[1].(int)]
	}).(OrganizationProjectOutput)
}

type OrganizationProjectMapOutput struct{ *pulumi.OutputState }

func (OrganizationProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationProject)(nil)).Elem()
}

func (o OrganizationProjectMapOutput) ToOrganizationProjectMapOutput() OrganizationProjectMapOutput {
	return o
}

func (o OrganizationProjectMapOutput) ToOrganizationProjectMapOutputWithContext(ctx context.Context) OrganizationProjectMapOutput {
	return o
}

func (o OrganizationProjectMapOutput) MapIndex(k pulumi.StringInput) OrganizationProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationProject {
		return vs[0].(map[string]*OrganizationProject)[vs[1].(string)]
	}).(OrganizationProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationProjectInput)(nil)).Elem(), &OrganizationProject{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationProjectArrayInput)(nil)).Elem(), OrganizationProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationProjectMapInput)(nil)).Elem(), OrganizationProjectMap{})
	pulumi.RegisterOutputType(OrganizationProjectOutput{})
	pulumi.RegisterOutputType(OrganizationProjectArrayOutput{})
	pulumi.RegisterOutputType(OrganizationProjectMapOutput{})
}
