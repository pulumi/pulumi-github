// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a repository ref.
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
//			_, err := github.GetRef(ctx, &github.GetRefArgs{
//				Owner:      pulumi.StringRef("example"),
//				Repository: "example",
//				Ref:        "heads/development",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRef(ctx *pulumi.Context, args *GetRefArgs, opts ...pulumi.InvokeOption) (*GetRefResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRefResult
	err := ctx.Invoke("github:index/getRef:getRef", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRef.
type GetRefArgs struct {
	// Owner of the repository.
	Owner *string `pulumi:"owner"`
	// The repository ref to look up. Must be formatted `heads/<ref>` for branches, and `tags/<ref>` for tags.
	Ref string `pulumi:"ref"`
	// The GitHub repository name.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getRef.
type GetRefResult struct {
	// An etag representing the ref.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	Owner      *string `pulumi:"owner"`
	Ref        string  `pulumi:"ref"`
	Repository string  `pulumi:"repository"`
	// A string storing the reference's `HEAD` commit's SHA1.
	Sha string `pulumi:"sha"`
}

func GetRefOutput(ctx *pulumi.Context, args GetRefOutputArgs, opts ...pulumi.InvokeOption) GetRefResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRefResultOutput, error) {
			args := v.(GetRefArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetRefResult
			secret, err := ctx.InvokePackageRaw("github:index/getRef:getRef", args, &rv, "", opts...)
			if err != nil {
				return GetRefResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetRefResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetRefResultOutput), nil
			}
			return output, nil
		}).(GetRefResultOutput)
}

// A collection of arguments for invoking getRef.
type GetRefOutputArgs struct {
	// Owner of the repository.
	Owner pulumi.StringPtrInput `pulumi:"owner"`
	// The repository ref to look up. Must be formatted `heads/<ref>` for branches, and `tags/<ref>` for tags.
	Ref pulumi.StringInput `pulumi:"ref"`
	// The GitHub repository name.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetRefOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRefArgs)(nil)).Elem()
}

// A collection of values returned by getRef.
type GetRefResultOutput struct{ *pulumi.OutputState }

func (GetRefResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRefResult)(nil)).Elem()
}

func (o GetRefResultOutput) ToGetRefResultOutput() GetRefResultOutput {
	return o
}

func (o GetRefResultOutput) ToGetRefResultOutputWithContext(ctx context.Context) GetRefResultOutput {
	return o
}

// An etag representing the ref.
func (o GetRefResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRefResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRefResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRefResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o GetRefResultOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Ref }).(pulumi.StringOutput)
}

func (o GetRefResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Repository }).(pulumi.StringOutput)
}

// A string storing the reference's `HEAD` commit's SHA1.
func (o GetRefResultOutput) Sha() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Sha }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRefResultOutput{})
}
