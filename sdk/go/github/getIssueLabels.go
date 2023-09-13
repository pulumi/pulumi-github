// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve the labels for a given repository.
func GetIssueLabels(ctx *pulumi.Context, args *GetIssueLabelsArgs, opts ...pulumi.InvokeOption) (*GetIssueLabelsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIssueLabelsResult
	err := ctx.Invoke("github:index/getIssueLabels:getIssueLabels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIssueLabels.
type GetIssueLabelsArgs struct {
	// The name of the repository.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getIssueLabels.
type GetIssueLabelsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of this repository's labels. Each element of `labels` has the following attributes:
	Labels     []GetIssueLabelsLabel `pulumi:"labels"`
	Repository string                `pulumi:"repository"`
}

func GetIssueLabelsOutput(ctx *pulumi.Context, args GetIssueLabelsOutputArgs, opts ...pulumi.InvokeOption) GetIssueLabelsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIssueLabelsResult, error) {
			args := v.(GetIssueLabelsArgs)
			r, err := GetIssueLabels(ctx, &args, opts...)
			var s GetIssueLabelsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIssueLabelsResultOutput)
}

// A collection of arguments for invoking getIssueLabels.
type GetIssueLabelsOutputArgs struct {
	// The name of the repository.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetIssueLabelsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIssueLabelsArgs)(nil)).Elem()
}

// A collection of values returned by getIssueLabels.
type GetIssueLabelsResultOutput struct{ *pulumi.OutputState }

func (GetIssueLabelsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIssueLabelsResult)(nil)).Elem()
}

func (o GetIssueLabelsResultOutput) ToGetIssueLabelsResultOutput() GetIssueLabelsResultOutput {
	return o
}

func (o GetIssueLabelsResultOutput) ToGetIssueLabelsResultOutputWithContext(ctx context.Context) GetIssueLabelsResultOutput {
	return o
}

func (o GetIssueLabelsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetIssueLabelsResult] {
	return pulumix.Output[GetIssueLabelsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetIssueLabelsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIssueLabelsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of this repository's labels. Each element of `labels` has the following attributes:
func (o GetIssueLabelsResultOutput) Labels() GetIssueLabelsLabelArrayOutput {
	return o.ApplyT(func(v GetIssueLabelsResult) []GetIssueLabelsLabel { return v.Labels }).(GetIssueLabelsLabelArrayOutput)
}

func (o GetIssueLabelsResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetIssueLabelsResult) string { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIssueLabelsResultOutput{})
}
