// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of dependabot secrets for a GitHub repository.
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
//			_, err := github.GetDependabotSecrets(ctx, &github.GetDependabotSecretsArgs{
//				Name: pulumi.StringRef("example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDependabotSecrets(ctx *pulumi.Context, args *GetDependabotSecretsArgs, opts ...pulumi.InvokeOption) (*GetDependabotSecretsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDependabotSecretsResult
	err := ctx.Invoke("github:index/getDependabotSecrets:getDependabotSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDependabotSecrets.
type GetDependabotSecretsArgs struct {
	// Full name of the repository (in `org/name` format).
	FullName *string `pulumi:"fullName"`
	// The name of the repository.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDependabotSecrets.
type GetDependabotSecretsResult struct {
	FullName string `pulumi:"fullName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Secret name
	Name string `pulumi:"name"`
	// list of dependabot secrets for the repository
	Secrets []GetDependabotSecretsSecret `pulumi:"secrets"`
}

func GetDependabotSecretsOutput(ctx *pulumi.Context, args GetDependabotSecretsOutputArgs, opts ...pulumi.InvokeOption) GetDependabotSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDependabotSecretsResult, error) {
			args := v.(GetDependabotSecretsArgs)
			r, err := GetDependabotSecrets(ctx, &args, opts...)
			var s GetDependabotSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDependabotSecretsResultOutput)
}

// A collection of arguments for invoking getDependabotSecrets.
type GetDependabotSecretsOutputArgs struct {
	// Full name of the repository (in `org/name` format).
	FullName pulumi.StringPtrInput `pulumi:"fullName"`
	// The name of the repository.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetDependabotSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDependabotSecretsArgs)(nil)).Elem()
}

// A collection of values returned by getDependabotSecrets.
type GetDependabotSecretsResultOutput struct{ *pulumi.OutputState }

func (GetDependabotSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDependabotSecretsResult)(nil)).Elem()
}

func (o GetDependabotSecretsResultOutput) ToGetDependabotSecretsResultOutput() GetDependabotSecretsResultOutput {
	return o
}

func (o GetDependabotSecretsResultOutput) ToGetDependabotSecretsResultOutputWithContext(ctx context.Context) GetDependabotSecretsResultOutput {
	return o
}

func (o GetDependabotSecretsResultOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDependabotSecretsResult) string { return v.FullName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDependabotSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDependabotSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Secret name
func (o GetDependabotSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDependabotSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

// list of dependabot secrets for the repository
func (o GetDependabotSecretsResultOutput) Secrets() GetDependabotSecretsSecretArrayOutput {
	return o.ApplyT(func(v GetDependabotSecretsResult) []GetDependabotSecretsSecret { return v.Secrets }).(GetDependabotSecretsSecretArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDependabotSecretsResultOutput{})
}
