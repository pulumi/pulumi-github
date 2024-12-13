// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the collaborators for a given repository.
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
//			_, err := github.GetCollaborators(ctx, &github.GetCollaboratorsArgs{
//				Owner:      "example_owner",
//				Repository: "example_repository",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCollaborators(ctx *pulumi.Context, args *GetCollaboratorsArgs, opts ...pulumi.InvokeOption) (*GetCollaboratorsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCollaboratorsResult
	err := ctx.Invoke("github:index/getCollaborators:getCollaborators", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCollaborators.
type GetCollaboratorsArgs struct {
	// Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
	Affiliation *string `pulumi:"affiliation"`
	// The organization that owns the repository.
	Owner string `pulumi:"owner"`
	// Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
	Permission *string `pulumi:"permission"`
	// The name of the repository.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getCollaborators.
type GetCollaboratorsResult struct {
	Affiliation *string `pulumi:"affiliation"`
	// An Array of GitHub collaborators.  Each `collaborator` block consists of the fields documented below.
	Collaborators []GetCollaboratorsCollaborator `pulumi:"collaborators"`
	// The provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	Owner string `pulumi:"owner"`
	// The permission of the collaborator.
	Permission *string `pulumi:"permission"`
	Repository string  `pulumi:"repository"`
}

func GetCollaboratorsOutput(ctx *pulumi.Context, args GetCollaboratorsOutputArgs, opts ...pulumi.InvokeOption) GetCollaboratorsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCollaboratorsResultOutput, error) {
			args := v.(GetCollaboratorsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getCollaborators:getCollaborators", args, GetCollaboratorsResultOutput{}, options).(GetCollaboratorsResultOutput), nil
		}).(GetCollaboratorsResultOutput)
}

// A collection of arguments for invoking getCollaborators.
type GetCollaboratorsOutputArgs struct {
	// Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
	Affiliation pulumi.StringPtrInput `pulumi:"affiliation"`
	// The organization that owns the repository.
	Owner pulumi.StringInput `pulumi:"owner"`
	// Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
	Permission pulumi.StringPtrInput `pulumi:"permission"`
	// The name of the repository.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetCollaboratorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCollaboratorsArgs)(nil)).Elem()
}

// A collection of values returned by getCollaborators.
type GetCollaboratorsResultOutput struct{ *pulumi.OutputState }

func (GetCollaboratorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCollaboratorsResult)(nil)).Elem()
}

func (o GetCollaboratorsResultOutput) ToGetCollaboratorsResultOutput() GetCollaboratorsResultOutput {
	return o
}

func (o GetCollaboratorsResultOutput) ToGetCollaboratorsResultOutputWithContext(ctx context.Context) GetCollaboratorsResultOutput {
	return o
}

func (o GetCollaboratorsResultOutput) Affiliation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCollaboratorsResult) *string { return v.Affiliation }).(pulumi.StringPtrOutput)
}

// An Array of GitHub collaborators.  Each `collaborator` block consists of the fields documented below.
func (o GetCollaboratorsResultOutput) Collaborators() GetCollaboratorsCollaboratorArrayOutput {
	return o.ApplyT(func(v GetCollaboratorsResult) []GetCollaboratorsCollaborator { return v.Collaborators }).(GetCollaboratorsCollaboratorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCollaboratorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCollaboratorsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCollaboratorsResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v GetCollaboratorsResult) string { return v.Owner }).(pulumi.StringOutput)
}

// The permission of the collaborator.
func (o GetCollaboratorsResultOutput) Permission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCollaboratorsResult) *string { return v.Permission }).(pulumi.StringPtrOutput)
}

func (o GetCollaboratorsResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetCollaboratorsResult) string { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCollaboratorsResultOutput{})
}
