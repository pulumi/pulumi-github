// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to generate a [GitHub App JWT](https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-json-web-token-jwt-for-a-github-app).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := github.GetAppToken(ctx, &github.GetAppTokenArgs{
//				AppId:          "123456",
//				InstallationId: "78910",
//				PemFile: std.File(ctx, &std.FileArgs{
//					Input: "foo/bar.pem",
//				}, nil).Result,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAppToken(ctx *pulumi.Context, args *GetAppTokenArgs, opts ...pulumi.InvokeOption) (*GetAppTokenResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppTokenResult
	err := ctx.Invoke("github:index/getAppToken:getAppToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppToken.
type GetAppTokenArgs struct {
	// This is the ID of the GitHub App.
	AppId string `pulumi:"appId"`
	// This is the ID of the GitHub App installation.
	InstallationId string `pulumi:"installationId"`
	// This is the contents of the GitHub App private key PEM file.
	PemFile string `pulumi:"pemFile"`
}

// A collection of values returned by getAppToken.
type GetAppTokenResult struct {
	AppId string `pulumi:"appId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	InstallationId string `pulumi:"installationId"`
	PemFile        string `pulumi:"pemFile"`
	// The generated GitHub APP JWT.
	Token string `pulumi:"token"`
}

func GetAppTokenOutput(ctx *pulumi.Context, args GetAppTokenOutputArgs, opts ...pulumi.InvokeOption) GetAppTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppTokenResult, error) {
			args := v.(GetAppTokenArgs)
			r, err := GetAppToken(ctx, &args, opts...)
			var s GetAppTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppTokenResultOutput)
}

// A collection of arguments for invoking getAppToken.
type GetAppTokenOutputArgs struct {
	// This is the ID of the GitHub App.
	AppId pulumi.StringInput `pulumi:"appId"`
	// This is the ID of the GitHub App installation.
	InstallationId pulumi.StringInput `pulumi:"installationId"`
	// This is the contents of the GitHub App private key PEM file.
	PemFile pulumi.StringInput `pulumi:"pemFile"`
}

func (GetAppTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppTokenArgs)(nil)).Elem()
}

// A collection of values returned by getAppToken.
type GetAppTokenResultOutput struct{ *pulumi.OutputState }

func (GetAppTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppTokenResult)(nil)).Elem()
}

func (o GetAppTokenResultOutput) ToGetAppTokenResultOutput() GetAppTokenResultOutput {
	return o
}

func (o GetAppTokenResultOutput) ToGetAppTokenResultOutputWithContext(ctx context.Context) GetAppTokenResultOutput {
	return o
}

func (o GetAppTokenResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppTokenResult) string { return v.AppId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppTokenResultOutput) InstallationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppTokenResult) string { return v.InstallationId }).(pulumi.StringOutput)
}

func (o GetAppTokenResultOutput) PemFile() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppTokenResult) string { return v.PemFile }).(pulumi.StringOutput)
}

// The generated GitHub APP JWT.
func (o GetAppTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppTokenResultOutput{})
}
