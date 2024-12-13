// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of codespaces secrets for a GitHub repository.
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
//			_, err := github.GetCodespacesSecrets(ctx, &github.GetCodespacesSecretsArgs{
//				Name: pulumi.StringRef("example_repository"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.GetCodespacesSecrets(ctx, &github.GetCodespacesSecretsArgs{
//				FullName: pulumi.StringRef("org/example_repository"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCodespacesSecrets(ctx *pulumi.Context, args *GetCodespacesSecretsArgs, opts ...pulumi.InvokeOption) (*GetCodespacesSecretsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCodespacesSecretsResult
	err := ctx.Invoke("github:index/getCodespacesSecrets:getCodespacesSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCodespacesSecrets.
type GetCodespacesSecretsArgs struct {
	// Full name of the repository (in `org/name` format).
	FullName *string `pulumi:"fullName"`
	// The name of the repository.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCodespacesSecrets.
type GetCodespacesSecretsResult struct {
	FullName string `pulumi:"fullName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Secret name
	Name string `pulumi:"name"`
	// list of codespaces secrets for the repository
	Secrets []GetCodespacesSecretsSecret `pulumi:"secrets"`
}

func GetCodespacesSecretsOutput(ctx *pulumi.Context, args GetCodespacesSecretsOutputArgs, opts ...pulumi.InvokeOption) GetCodespacesSecretsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCodespacesSecretsResultOutput, error) {
			args := v.(GetCodespacesSecretsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getCodespacesSecrets:getCodespacesSecrets", args, GetCodespacesSecretsResultOutput{}, options).(GetCodespacesSecretsResultOutput), nil
		}).(GetCodespacesSecretsResultOutput)
}

// A collection of arguments for invoking getCodespacesSecrets.
type GetCodespacesSecretsOutputArgs struct {
	// Full name of the repository (in `org/name` format).
	FullName pulumi.StringPtrInput `pulumi:"fullName"`
	// The name of the repository.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetCodespacesSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCodespacesSecretsArgs)(nil)).Elem()
}

// A collection of values returned by getCodespacesSecrets.
type GetCodespacesSecretsResultOutput struct{ *pulumi.OutputState }

func (GetCodespacesSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCodespacesSecretsResult)(nil)).Elem()
}

func (o GetCodespacesSecretsResultOutput) ToGetCodespacesSecretsResultOutput() GetCodespacesSecretsResultOutput {
	return o
}

func (o GetCodespacesSecretsResultOutput) ToGetCodespacesSecretsResultOutputWithContext(ctx context.Context) GetCodespacesSecretsResultOutput {
	return o
}

func (o GetCodespacesSecretsResultOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCodespacesSecretsResult) string { return v.FullName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCodespacesSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCodespacesSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Secret name
func (o GetCodespacesSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCodespacesSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

// list of codespaces secrets for the repository
func (o GetCodespacesSecretsResultOutput) Secrets() GetCodespacesSecretsSecretArrayOutput {
	return o.ApplyT(func(v GetCodespacesSecretsResult) []GetCodespacesSecretsSecret { return v.Secrets }).(GetCodespacesSecretsSecretArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCodespacesSecretsResultOutput{})
}
