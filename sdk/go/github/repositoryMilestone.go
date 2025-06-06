// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a GitHub repository milestone resource.
//
// This resource allows you to create and manage milestones for a GitHub Repository within an organization or user account.
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
//			// Create a milestone for a repository
//			_, err := github.NewRepositoryMilestone(ctx, "example", &github.RepositoryMilestoneArgs{
//				Owner:      pulumi.String("example-owner"),
//				Repository: pulumi.String("example-repository"),
//				Title:      pulumi.String("v1.1.0"),
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
// A GitHub Repository Milestone can be imported using an ID made up of `owner/repository/number`, e.g.
//
// ```sh
// $ pulumi import github:index/repositoryMilestone:RepositoryMilestone example example-owner/example-repository/1
// ```
type RepositoryMilestone struct {
	pulumi.CustomResourceState

	// A description of the milestone.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate pulumi.StringPtrOutput `pulumi:"dueDate"`
	// The number of the milestone.
	Number pulumi.IntOutput `pulumi:"number"`
	// The owner of the GitHub Repository.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The name of the GitHub Repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The title of the milestone.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewRepositoryMilestone registers a new resource with the given unique name, arguments, and options.
func NewRepositoryMilestone(ctx *pulumi.Context,
	name string, args *RepositoryMilestoneArgs, opts ...pulumi.ResourceOption) (*RepositoryMilestone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Owner == nil {
		return nil, errors.New("invalid value for required argument 'Owner'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryMilestone
	err := ctx.RegisterResource("github:index/repositoryMilestone:RepositoryMilestone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryMilestone gets an existing RepositoryMilestone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryMilestone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryMilestoneState, opts ...pulumi.ResourceOption) (*RepositoryMilestone, error) {
	var resource RepositoryMilestone
	err := ctx.ReadResource("github:index/repositoryMilestone:RepositoryMilestone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryMilestone resources.
type repositoryMilestoneState struct {
	// A description of the milestone.
	Description *string `pulumi:"description"`
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate *string `pulumi:"dueDate"`
	// The number of the milestone.
	Number *int `pulumi:"number"`
	// The owner of the GitHub Repository.
	Owner *string `pulumi:"owner"`
	// The name of the GitHub Repository.
	Repository *string `pulumi:"repository"`
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State *string `pulumi:"state"`
	// The title of the milestone.
	Title *string `pulumi:"title"`
}

type RepositoryMilestoneState struct {
	// A description of the milestone.
	Description pulumi.StringPtrInput
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate pulumi.StringPtrInput
	// The number of the milestone.
	Number pulumi.IntPtrInput
	// The owner of the GitHub Repository.
	Owner pulumi.StringPtrInput
	// The name of the GitHub Repository.
	Repository pulumi.StringPtrInput
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State pulumi.StringPtrInput
	// The title of the milestone.
	Title pulumi.StringPtrInput
}

func (RepositoryMilestoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryMilestoneState)(nil)).Elem()
}

type repositoryMilestoneArgs struct {
	// A description of the milestone.
	Description *string `pulumi:"description"`
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate *string `pulumi:"dueDate"`
	// The owner of the GitHub Repository.
	Owner string `pulumi:"owner"`
	// The name of the GitHub Repository.
	Repository string `pulumi:"repository"`
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State *string `pulumi:"state"`
	// The title of the milestone.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a RepositoryMilestone resource.
type RepositoryMilestoneArgs struct {
	// A description of the milestone.
	Description pulumi.StringPtrInput
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate pulumi.StringPtrInput
	// The owner of the GitHub Repository.
	Owner pulumi.StringInput
	// The name of the GitHub Repository.
	Repository pulumi.StringInput
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State pulumi.StringPtrInput
	// The title of the milestone.
	Title pulumi.StringInput
}

func (RepositoryMilestoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryMilestoneArgs)(nil)).Elem()
}

type RepositoryMilestoneInput interface {
	pulumi.Input

	ToRepositoryMilestoneOutput() RepositoryMilestoneOutput
	ToRepositoryMilestoneOutputWithContext(ctx context.Context) RepositoryMilestoneOutput
}

func (*RepositoryMilestone) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryMilestone)(nil)).Elem()
}

func (i *RepositoryMilestone) ToRepositoryMilestoneOutput() RepositoryMilestoneOutput {
	return i.ToRepositoryMilestoneOutputWithContext(context.Background())
}

func (i *RepositoryMilestone) ToRepositoryMilestoneOutputWithContext(ctx context.Context) RepositoryMilestoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMilestoneOutput)
}

// RepositoryMilestoneArrayInput is an input type that accepts RepositoryMilestoneArray and RepositoryMilestoneArrayOutput values.
// You can construct a concrete instance of `RepositoryMilestoneArrayInput` via:
//
//	RepositoryMilestoneArray{ RepositoryMilestoneArgs{...} }
type RepositoryMilestoneArrayInput interface {
	pulumi.Input

	ToRepositoryMilestoneArrayOutput() RepositoryMilestoneArrayOutput
	ToRepositoryMilestoneArrayOutputWithContext(context.Context) RepositoryMilestoneArrayOutput
}

type RepositoryMilestoneArray []RepositoryMilestoneInput

func (RepositoryMilestoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryMilestone)(nil)).Elem()
}

func (i RepositoryMilestoneArray) ToRepositoryMilestoneArrayOutput() RepositoryMilestoneArrayOutput {
	return i.ToRepositoryMilestoneArrayOutputWithContext(context.Background())
}

func (i RepositoryMilestoneArray) ToRepositoryMilestoneArrayOutputWithContext(ctx context.Context) RepositoryMilestoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMilestoneArrayOutput)
}

// RepositoryMilestoneMapInput is an input type that accepts RepositoryMilestoneMap and RepositoryMilestoneMapOutput values.
// You can construct a concrete instance of `RepositoryMilestoneMapInput` via:
//
//	RepositoryMilestoneMap{ "key": RepositoryMilestoneArgs{...} }
type RepositoryMilestoneMapInput interface {
	pulumi.Input

	ToRepositoryMilestoneMapOutput() RepositoryMilestoneMapOutput
	ToRepositoryMilestoneMapOutputWithContext(context.Context) RepositoryMilestoneMapOutput
}

type RepositoryMilestoneMap map[string]RepositoryMilestoneInput

func (RepositoryMilestoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryMilestone)(nil)).Elem()
}

func (i RepositoryMilestoneMap) ToRepositoryMilestoneMapOutput() RepositoryMilestoneMapOutput {
	return i.ToRepositoryMilestoneMapOutputWithContext(context.Background())
}

func (i RepositoryMilestoneMap) ToRepositoryMilestoneMapOutputWithContext(ctx context.Context) RepositoryMilestoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMilestoneMapOutput)
}

type RepositoryMilestoneOutput struct{ *pulumi.OutputState }

func (RepositoryMilestoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryMilestone)(nil)).Elem()
}

func (o RepositoryMilestoneOutput) ToRepositoryMilestoneOutput() RepositoryMilestoneOutput {
	return o
}

func (o RepositoryMilestoneOutput) ToRepositoryMilestoneOutputWithContext(ctx context.Context) RepositoryMilestoneOutput {
	return o
}

// A description of the milestone.
func (o RepositoryMilestoneOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryMilestone) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The milestone due date. In `yyyy-mm-dd` format.
func (o RepositoryMilestoneOutput) DueDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryMilestone) pulumi.StringPtrOutput { return v.DueDate }).(pulumi.StringPtrOutput)
}

// The number of the milestone.
func (o RepositoryMilestoneOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v *RepositoryMilestone) pulumi.IntOutput { return v.Number }).(pulumi.IntOutput)
}

// The owner of the GitHub Repository.
func (o RepositoryMilestoneOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryMilestone) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The name of the GitHub Repository.
func (o RepositoryMilestoneOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryMilestone) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The state of the milestone. Either `open` or `closed`. Default: `open`
func (o RepositoryMilestoneOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryMilestone) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The title of the milestone.
func (o RepositoryMilestoneOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryMilestone) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

type RepositoryMilestoneArrayOutput struct{ *pulumi.OutputState }

func (RepositoryMilestoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryMilestone)(nil)).Elem()
}

func (o RepositoryMilestoneArrayOutput) ToRepositoryMilestoneArrayOutput() RepositoryMilestoneArrayOutput {
	return o
}

func (o RepositoryMilestoneArrayOutput) ToRepositoryMilestoneArrayOutputWithContext(ctx context.Context) RepositoryMilestoneArrayOutput {
	return o
}

func (o RepositoryMilestoneArrayOutput) Index(i pulumi.IntInput) RepositoryMilestoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryMilestone {
		return vs[0].([]*RepositoryMilestone)[vs[1].(int)]
	}).(RepositoryMilestoneOutput)
}

type RepositoryMilestoneMapOutput struct{ *pulumi.OutputState }

func (RepositoryMilestoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryMilestone)(nil)).Elem()
}

func (o RepositoryMilestoneMapOutput) ToRepositoryMilestoneMapOutput() RepositoryMilestoneMapOutput {
	return o
}

func (o RepositoryMilestoneMapOutput) ToRepositoryMilestoneMapOutputWithContext(ctx context.Context) RepositoryMilestoneMapOutput {
	return o
}

func (o RepositoryMilestoneMapOutput) MapIndex(k pulumi.StringInput) RepositoryMilestoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryMilestone {
		return vs[0].(map[string]*RepositoryMilestone)[vs[1].(string)]
	}).(RepositoryMilestoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMilestoneInput)(nil)).Elem(), &RepositoryMilestone{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMilestoneArrayInput)(nil)).Elem(), RepositoryMilestoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMilestoneMapInput)(nil)).Elem(), RepositoryMilestoneMap{})
	pulumi.RegisterOutputType(RepositoryMilestoneOutput{})
	pulumi.RegisterOutputType(RepositoryMilestoneArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMilestoneMapOutput{})
}
