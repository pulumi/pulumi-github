// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a GitHub Actions public key. This data source is required to be used with other GitHub secrets interactions.
// Note that the provider `token` must have admin rights to a repository to retrieve it's action public key.
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
//			_, err := github.GetActionsPublicKey(ctx, &github.GetActionsPublicKeyArgs{
//				Repository: "example_repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetActionsPublicKey(ctx *pulumi.Context, args *GetActionsPublicKeyArgs, opts ...pulumi.InvokeOption) (*GetActionsPublicKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetActionsPublicKeyResult
	err := ctx.Invoke("github:index/getActionsPublicKey:getActionsPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActionsPublicKey.
type GetActionsPublicKeyArgs struct {
	// Name of the repository to get public key from.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getActionsPublicKey.
type GetActionsPublicKeyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Actual key retrieved.
	Key string `pulumi:"key"`
	// ID of the key that has been retrieved.
	KeyId      string `pulumi:"keyId"`
	Repository string `pulumi:"repository"`
}

func GetActionsPublicKeyOutput(ctx *pulumi.Context, args GetActionsPublicKeyOutputArgs, opts ...pulumi.InvokeOption) GetActionsPublicKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetActionsPublicKeyResult, error) {
			args := v.(GetActionsPublicKeyArgs)
			r, err := GetActionsPublicKey(ctx, &args, opts...)
			var s GetActionsPublicKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetActionsPublicKeyResultOutput)
}

// A collection of arguments for invoking getActionsPublicKey.
type GetActionsPublicKeyOutputArgs struct {
	// Name of the repository to get public key from.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetActionsPublicKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsPublicKeyArgs)(nil)).Elem()
}

// A collection of values returned by getActionsPublicKey.
type GetActionsPublicKeyResultOutput struct{ *pulumi.OutputState }

func (GetActionsPublicKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsPublicKeyResult)(nil)).Elem()
}

func (o GetActionsPublicKeyResultOutput) ToGetActionsPublicKeyResultOutput() GetActionsPublicKeyResultOutput {
	return o
}

func (o GetActionsPublicKeyResultOutput) ToGetActionsPublicKeyResultOutputWithContext(ctx context.Context) GetActionsPublicKeyResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetActionsPublicKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsPublicKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Actual key retrieved.
func (o GetActionsPublicKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsPublicKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

// ID of the key that has been retrieved.
func (o GetActionsPublicKeyResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsPublicKeyResult) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o GetActionsPublicKeyResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsPublicKeyResult) string { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActionsPublicKeyResultOutput{})
}
