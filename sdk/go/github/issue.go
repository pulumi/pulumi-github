// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a GitHub issue resource.
//
// This resource allows you to create and manage issue within your
// GitHub repository.
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
// 		testRepository, err := github.NewRepository(ctx, "testRepository", &github.RepositoryArgs{
// 			AutoInit:  pulumi.Bool(true),
// 			HasIssues: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = github.NewIssue(ctx, "testIssue", &github.IssueArgs{
// 			Repository: testRepository.Name,
// 			Title:      pulumi.String("My issue title"),
// 			Body:       pulumi.String("The body of my issue"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// GitHub Issues can be imported using an ID made up of `repository:number`, e.g.
//
// ```sh
//  $ pulumi import github:index/issue:Issue issue_15 myrepo:15
// ```
type Issue struct {
	pulumi.CustomResourceState

	// List of Logins to assign the to the issue.
	Assignees pulumi.StringArrayOutput `pulumi:"assignees"`
	// Title of the issue.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	Etag pulumi.StringOutput    `pulumi:"etag"`
	// (Computed) - The issue id
	IssueId pulumi.IntOutput `pulumi:"issueId"`
	// List of labels to attach to the issue.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Milestone number to assign to the issue
	MilestoneNumber pulumi.IntPtrOutput `pulumi:"milestoneNumber"`
	// (Computed) - The issue number
	Number pulumi.IntOutput `pulumi:"number"`
	// The GitHub repository name
	Repository pulumi.StringOutput `pulumi:"repository"`
	// Title of the issue.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewIssue registers a new resource with the given unique name, arguments, and options.
func NewIssue(ctx *pulumi.Context,
	name string, args *IssueArgs, opts ...pulumi.ResourceOption) (*Issue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource Issue
	err := ctx.RegisterResource("github:index/issue:Issue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIssue gets an existing Issue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIssue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IssueState, opts ...pulumi.ResourceOption) (*Issue, error) {
	var resource Issue
	err := ctx.ReadResource("github:index/issue:Issue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Issue resources.
type issueState struct {
	// List of Logins to assign the to the issue.
	Assignees []string `pulumi:"assignees"`
	// Title of the issue.
	Body *string `pulumi:"body"`
	Etag *string `pulumi:"etag"`
	// (Computed) - The issue id
	IssueId *int `pulumi:"issueId"`
	// List of labels to attach to the issue.
	Labels []string `pulumi:"labels"`
	// Milestone number to assign to the issue
	MilestoneNumber *int `pulumi:"milestoneNumber"`
	// (Computed) - The issue number
	Number *int `pulumi:"number"`
	// The GitHub repository name
	Repository *string `pulumi:"repository"`
	// Title of the issue.
	Title *string `pulumi:"title"`
}

type IssueState struct {
	// List of Logins to assign the to the issue.
	Assignees pulumi.StringArrayInput
	// Title of the issue.
	Body pulumi.StringPtrInput
	Etag pulumi.StringPtrInput
	// (Computed) - The issue id
	IssueId pulumi.IntPtrInput
	// List of labels to attach to the issue.
	Labels pulumi.StringArrayInput
	// Milestone number to assign to the issue
	MilestoneNumber pulumi.IntPtrInput
	// (Computed) - The issue number
	Number pulumi.IntPtrInput
	// The GitHub repository name
	Repository pulumi.StringPtrInput
	// Title of the issue.
	Title pulumi.StringPtrInput
}

func (IssueState) ElementType() reflect.Type {
	return reflect.TypeOf((*issueState)(nil)).Elem()
}

type issueArgs struct {
	// List of Logins to assign the to the issue.
	Assignees []string `pulumi:"assignees"`
	// Title of the issue.
	Body *string `pulumi:"body"`
	// List of labels to attach to the issue.
	Labels []string `pulumi:"labels"`
	// Milestone number to assign to the issue
	MilestoneNumber *int `pulumi:"milestoneNumber"`
	// The GitHub repository name
	Repository string `pulumi:"repository"`
	// Title of the issue.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a Issue resource.
type IssueArgs struct {
	// List of Logins to assign the to the issue.
	Assignees pulumi.StringArrayInput
	// Title of the issue.
	Body pulumi.StringPtrInput
	// List of labels to attach to the issue.
	Labels pulumi.StringArrayInput
	// Milestone number to assign to the issue
	MilestoneNumber pulumi.IntPtrInput
	// The GitHub repository name
	Repository pulumi.StringInput
	// Title of the issue.
	Title pulumi.StringInput
}

func (IssueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*issueArgs)(nil)).Elem()
}

type IssueInput interface {
	pulumi.Input

	ToIssueOutput() IssueOutput
	ToIssueOutputWithContext(ctx context.Context) IssueOutput
}

func (*Issue) ElementType() reflect.Type {
	return reflect.TypeOf((**Issue)(nil)).Elem()
}

func (i *Issue) ToIssueOutput() IssueOutput {
	return i.ToIssueOutputWithContext(context.Background())
}

func (i *Issue) ToIssueOutputWithContext(ctx context.Context) IssueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IssueOutput)
}

// IssueArrayInput is an input type that accepts IssueArray and IssueArrayOutput values.
// You can construct a concrete instance of `IssueArrayInput` via:
//
//          IssueArray{ IssueArgs{...} }
type IssueArrayInput interface {
	pulumi.Input

	ToIssueArrayOutput() IssueArrayOutput
	ToIssueArrayOutputWithContext(context.Context) IssueArrayOutput
}

type IssueArray []IssueInput

func (IssueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Issue)(nil)).Elem()
}

func (i IssueArray) ToIssueArrayOutput() IssueArrayOutput {
	return i.ToIssueArrayOutputWithContext(context.Background())
}

func (i IssueArray) ToIssueArrayOutputWithContext(ctx context.Context) IssueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IssueArrayOutput)
}

// IssueMapInput is an input type that accepts IssueMap and IssueMapOutput values.
// You can construct a concrete instance of `IssueMapInput` via:
//
//          IssueMap{ "key": IssueArgs{...} }
type IssueMapInput interface {
	pulumi.Input

	ToIssueMapOutput() IssueMapOutput
	ToIssueMapOutputWithContext(context.Context) IssueMapOutput
}

type IssueMap map[string]IssueInput

func (IssueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Issue)(nil)).Elem()
}

func (i IssueMap) ToIssueMapOutput() IssueMapOutput {
	return i.ToIssueMapOutputWithContext(context.Background())
}

func (i IssueMap) ToIssueMapOutputWithContext(ctx context.Context) IssueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IssueMapOutput)
}

type IssueOutput struct{ *pulumi.OutputState }

func (IssueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Issue)(nil)).Elem()
}

func (o IssueOutput) ToIssueOutput() IssueOutput {
	return o
}

func (o IssueOutput) ToIssueOutputWithContext(ctx context.Context) IssueOutput {
	return o
}

type IssueArrayOutput struct{ *pulumi.OutputState }

func (IssueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Issue)(nil)).Elem()
}

func (o IssueArrayOutput) ToIssueArrayOutput() IssueArrayOutput {
	return o
}

func (o IssueArrayOutput) ToIssueArrayOutputWithContext(ctx context.Context) IssueArrayOutput {
	return o
}

func (o IssueArrayOutput) Index(i pulumi.IntInput) IssueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Issue {
		return vs[0].([]*Issue)[vs[1].(int)]
	}).(IssueOutput)
}

type IssueMapOutput struct{ *pulumi.OutputState }

func (IssueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Issue)(nil)).Elem()
}

func (o IssueMapOutput) ToIssueMapOutput() IssueMapOutput {
	return o
}

func (o IssueMapOutput) ToIssueMapOutputWithContext(ctx context.Context) IssueMapOutput {
	return o
}

func (o IssueMapOutput) MapIndex(k pulumi.StringInput) IssueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Issue {
		return vs[0].(map[string]*Issue)[vs[1].(string)]
	}).(IssueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IssueInput)(nil)).Elem(), &Issue{})
	pulumi.RegisterInputType(reflect.TypeOf((*IssueArrayInput)(nil)).Elem(), IssueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IssueMapInput)(nil)).Elem(), IssueMap{})
	pulumi.RegisterOutputType(IssueOutput{})
	pulumi.RegisterOutputType(IssueArrayOutput{})
	pulumi.RegisterOutputType(IssueMapOutput{})
}
