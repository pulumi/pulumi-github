// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
// 		_, err := github.NewIssueLabel(ctx, "testRepo", &github.IssueLabelArgs{
// 			Color:      pulumi.String("FF0000"),
// 			Repository: pulumi.String("test-repo"),
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
// GitHub Issue Labels can be imported using an ID made up of `repository:name`, e.g.
//
// ```sh
//  $ pulumi import github:index/issueLabel:IssueLabel panic_label terraform:panic
// ```
type IssueLabel struct {
	pulumi.CustomResourceState

	// A 6 character hex code, **without the leading #**, identifying the color of the label.
	Color pulumi.StringOutput `pulumi:"color"`
	// A short description of the label.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Etag        pulumi.StringOutput    `pulumi:"etag"`
	// The name of the label.
	Name pulumi.StringOutput `pulumi:"name"`
	// The GitHub repository
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The URL to the issue label
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewIssueLabel registers a new resource with the given unique name, arguments, and options.
func NewIssueLabel(ctx *pulumi.Context,
	name string, args *IssueLabelArgs, opts ...pulumi.ResourceOption) (*IssueLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Color == nil {
		return nil, errors.New("invalid value for required argument 'Color'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource IssueLabel
	err := ctx.RegisterResource("github:index/issueLabel:IssueLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIssueLabel gets an existing IssueLabel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIssueLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IssueLabelState, opts ...pulumi.ResourceOption) (*IssueLabel, error) {
	var resource IssueLabel
	err := ctx.ReadResource("github:index/issueLabel:IssueLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IssueLabel resources.
type issueLabelState struct {
	// A 6 character hex code, **without the leading #**, identifying the color of the label.
	Color *string `pulumi:"color"`
	// A short description of the label.
	Description *string `pulumi:"description"`
	Etag        *string `pulumi:"etag"`
	// The name of the label.
	Name *string `pulumi:"name"`
	// The GitHub repository
	Repository *string `pulumi:"repository"`
	// The URL to the issue label
	Url *string `pulumi:"url"`
}

type IssueLabelState struct {
	// A 6 character hex code, **without the leading #**, identifying the color of the label.
	Color pulumi.StringPtrInput
	// A short description of the label.
	Description pulumi.StringPtrInput
	Etag        pulumi.StringPtrInput
	// The name of the label.
	Name pulumi.StringPtrInput
	// The GitHub repository
	Repository pulumi.StringPtrInput
	// The URL to the issue label
	Url pulumi.StringPtrInput
}

func (IssueLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*issueLabelState)(nil)).Elem()
}

type issueLabelArgs struct {
	// A 6 character hex code, **without the leading #**, identifying the color of the label.
	Color string `pulumi:"color"`
	// A short description of the label.
	Description *string `pulumi:"description"`
	// The name of the label.
	Name *string `pulumi:"name"`
	// The GitHub repository
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a IssueLabel resource.
type IssueLabelArgs struct {
	// A 6 character hex code, **without the leading #**, identifying the color of the label.
	Color pulumi.StringInput
	// A short description of the label.
	Description pulumi.StringPtrInput
	// The name of the label.
	Name pulumi.StringPtrInput
	// The GitHub repository
	Repository pulumi.StringInput
}

func (IssueLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*issueLabelArgs)(nil)).Elem()
}

type IssueLabelInput interface {
	pulumi.Input

	ToIssueLabelOutput() IssueLabelOutput
	ToIssueLabelOutputWithContext(ctx context.Context) IssueLabelOutput
}

func (*IssueLabel) ElementType() reflect.Type {
	return reflect.TypeOf((**IssueLabel)(nil)).Elem()
}

func (i *IssueLabel) ToIssueLabelOutput() IssueLabelOutput {
	return i.ToIssueLabelOutputWithContext(context.Background())
}

func (i *IssueLabel) ToIssueLabelOutputWithContext(ctx context.Context) IssueLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IssueLabelOutput)
}

// IssueLabelArrayInput is an input type that accepts IssueLabelArray and IssueLabelArrayOutput values.
// You can construct a concrete instance of `IssueLabelArrayInput` via:
//
//          IssueLabelArray{ IssueLabelArgs{...} }
type IssueLabelArrayInput interface {
	pulumi.Input

	ToIssueLabelArrayOutput() IssueLabelArrayOutput
	ToIssueLabelArrayOutputWithContext(context.Context) IssueLabelArrayOutput
}

type IssueLabelArray []IssueLabelInput

func (IssueLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IssueLabel)(nil)).Elem()
}

func (i IssueLabelArray) ToIssueLabelArrayOutput() IssueLabelArrayOutput {
	return i.ToIssueLabelArrayOutputWithContext(context.Background())
}

func (i IssueLabelArray) ToIssueLabelArrayOutputWithContext(ctx context.Context) IssueLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IssueLabelArrayOutput)
}

// IssueLabelMapInput is an input type that accepts IssueLabelMap and IssueLabelMapOutput values.
// You can construct a concrete instance of `IssueLabelMapInput` via:
//
//          IssueLabelMap{ "key": IssueLabelArgs{...} }
type IssueLabelMapInput interface {
	pulumi.Input

	ToIssueLabelMapOutput() IssueLabelMapOutput
	ToIssueLabelMapOutputWithContext(context.Context) IssueLabelMapOutput
}

type IssueLabelMap map[string]IssueLabelInput

func (IssueLabelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IssueLabel)(nil)).Elem()
}

func (i IssueLabelMap) ToIssueLabelMapOutput() IssueLabelMapOutput {
	return i.ToIssueLabelMapOutputWithContext(context.Background())
}

func (i IssueLabelMap) ToIssueLabelMapOutputWithContext(ctx context.Context) IssueLabelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IssueLabelMapOutput)
}

type IssueLabelOutput struct{ *pulumi.OutputState }

func (IssueLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IssueLabel)(nil)).Elem()
}

func (o IssueLabelOutput) ToIssueLabelOutput() IssueLabelOutput {
	return o
}

func (o IssueLabelOutput) ToIssueLabelOutputWithContext(ctx context.Context) IssueLabelOutput {
	return o
}

type IssueLabelArrayOutput struct{ *pulumi.OutputState }

func (IssueLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IssueLabel)(nil)).Elem()
}

func (o IssueLabelArrayOutput) ToIssueLabelArrayOutput() IssueLabelArrayOutput {
	return o
}

func (o IssueLabelArrayOutput) ToIssueLabelArrayOutputWithContext(ctx context.Context) IssueLabelArrayOutput {
	return o
}

func (o IssueLabelArrayOutput) Index(i pulumi.IntInput) IssueLabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IssueLabel {
		return vs[0].([]*IssueLabel)[vs[1].(int)]
	}).(IssueLabelOutput)
}

type IssueLabelMapOutput struct{ *pulumi.OutputState }

func (IssueLabelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IssueLabel)(nil)).Elem()
}

func (o IssueLabelMapOutput) ToIssueLabelMapOutput() IssueLabelMapOutput {
	return o
}

func (o IssueLabelMapOutput) ToIssueLabelMapOutputWithContext(ctx context.Context) IssueLabelMapOutput {
	return o
}

func (o IssueLabelMapOutput) MapIndex(k pulumi.StringInput) IssueLabelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IssueLabel {
		return vs[0].(map[string]*IssueLabel)[vs[1].(string)]
	}).(IssueLabelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IssueLabelInput)(nil)).Elem(), &IssueLabel{})
	pulumi.RegisterInputType(reflect.TypeOf((*IssueLabelArrayInput)(nil)).Elem(), IssueLabelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IssueLabelMapInput)(nil)).Elem(), IssueLabelMap{})
	pulumi.RegisterOutputType(IssueLabelOutput{})
	pulumi.RegisterOutputType(IssueLabelArrayOutput{})
	pulumi.RegisterOutputType(IssueLabelMapOutput{})
}
