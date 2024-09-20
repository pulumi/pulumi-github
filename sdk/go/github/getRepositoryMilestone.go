// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a specific GitHub milestone in a repository.
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
//			_, err := github.LookupRepositoryMilestone(ctx, &github.LookupRepositoryMilestoneArgs{
//				Owner:      "example-owner",
//				Repository: "example-repository",
//				Number:     1,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRepositoryMilestone(ctx *pulumi.Context, args *LookupRepositoryMilestoneArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryMilestoneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRepositoryMilestoneResult
	err := ctx.Invoke("github:index/getRepositoryMilestone:getRepositoryMilestone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryMilestone.
type LookupRepositoryMilestoneArgs struct {
	// The number of the milestone.
	Number int `pulumi:"number"`
	// Owner of the repository.
	Owner string `pulumi:"owner"`
	// Name of the repository to retrieve the milestone from.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getRepositoryMilestone.
type LookupRepositoryMilestoneResult struct {
	// Description of the milestone.
	Description string `pulumi:"description"`
	// The milestone due date (in ISO-8601 `yyyy-mm-dd` format).
	DueDate string `pulumi:"dueDate"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Number     int    `pulumi:"number"`
	Owner      string `pulumi:"owner"`
	Repository string `pulumi:"repository"`
	// State of the milestone.
	State string `pulumi:"state"`
	// Title of the milestone.
	Title string `pulumi:"title"`
}

func LookupRepositoryMilestoneOutput(ctx *pulumi.Context, args LookupRepositoryMilestoneOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryMilestoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRepositoryMilestoneResultOutput, error) {
			args := v.(LookupRepositoryMilestoneArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupRepositoryMilestoneResult
			secret, err := ctx.InvokePackageRaw("github:index/getRepositoryMilestone:getRepositoryMilestone", args, &rv, "", opts...)
			if err != nil {
				return LookupRepositoryMilestoneResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupRepositoryMilestoneResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupRepositoryMilestoneResultOutput), nil
			}
			return output, nil
		}).(LookupRepositoryMilestoneResultOutput)
}

// A collection of arguments for invoking getRepositoryMilestone.
type LookupRepositoryMilestoneOutputArgs struct {
	// The number of the milestone.
	Number pulumi.IntInput `pulumi:"number"`
	// Owner of the repository.
	Owner pulumi.StringInput `pulumi:"owner"`
	// Name of the repository to retrieve the milestone from.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (LookupRepositoryMilestoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryMilestoneArgs)(nil)).Elem()
}

// A collection of values returned by getRepositoryMilestone.
type LookupRepositoryMilestoneResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryMilestoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryMilestoneResult)(nil)).Elem()
}

func (o LookupRepositoryMilestoneResultOutput) ToLookupRepositoryMilestoneResultOutput() LookupRepositoryMilestoneResultOutput {
	return o
}

func (o LookupRepositoryMilestoneResultOutput) ToLookupRepositoryMilestoneResultOutputWithContext(ctx context.Context) LookupRepositoryMilestoneResultOutput {
	return o
}

// Description of the milestone.
func (o LookupRepositoryMilestoneResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryMilestoneResult) string { return v.Description }).(pulumi.StringOutput)
}

// The milestone due date (in ISO-8601 `yyyy-mm-dd` format).
func (o LookupRepositoryMilestoneResultOutput) DueDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryMilestoneResult) string { return v.DueDate }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRepositoryMilestoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryMilestoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRepositoryMilestoneResultOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRepositoryMilestoneResult) int { return v.Number }).(pulumi.IntOutput)
}

func (o LookupRepositoryMilestoneResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryMilestoneResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupRepositoryMilestoneResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryMilestoneResult) string { return v.Repository }).(pulumi.StringOutput)
}

// State of the milestone.
func (o LookupRepositoryMilestoneResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryMilestoneResult) string { return v.State }).(pulumi.StringOutput)
}

// Title of the milestone.
func (o LookupRepositoryMilestoneResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryMilestoneResult) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryMilestoneResultOutput{})
}
