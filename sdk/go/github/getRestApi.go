// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a GitHub resource through REST API.
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
//			_, err := github.GetRestApi(ctx, &github.GetRestApiArgs{
//				Endpoint: "repos/example_repo/git/refs/heads/main",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRestApi(ctx *pulumi.Context, args *GetRestApiArgs, opts ...pulumi.InvokeOption) (*GetRestApiResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRestApiResult
	err := ctx.Invoke("github:index/getRestApi:getRestApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRestApi.
type GetRestApiArgs struct {
	// REST API endpoint to send the GET request to.
	Endpoint string `pulumi:"endpoint"`
}

// A collection of values returned by getRestApi.
type GetRestApiResult struct {
	// A JSON string containing response body.
	Body string `pulumi:"body"`
	// A response status code.
	Code     int    `pulumi:"code"`
	Endpoint string `pulumi:"endpoint"`
	// A JSON string containing response headers.
	Headers string `pulumi:"headers"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A response status string.
	Status string `pulumi:"status"`
}

func GetRestApiOutput(ctx *pulumi.Context, args GetRestApiOutputArgs, opts ...pulumi.InvokeOption) GetRestApiResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetRestApiResultOutput, error) {
			args := v.(GetRestApiArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getRestApi:getRestApi", args, GetRestApiResultOutput{}, options).(GetRestApiResultOutput), nil
		}).(GetRestApiResultOutput)
}

// A collection of arguments for invoking getRestApi.
type GetRestApiOutputArgs struct {
	// REST API endpoint to send the GET request to.
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
}

func (GetRestApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRestApiArgs)(nil)).Elem()
}

// A collection of values returned by getRestApi.
type GetRestApiResultOutput struct{ *pulumi.OutputState }

func (GetRestApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRestApiResult)(nil)).Elem()
}

func (o GetRestApiResultOutput) ToGetRestApiResultOutput() GetRestApiResultOutput {
	return o
}

func (o GetRestApiResultOutput) ToGetRestApiResultOutputWithContext(ctx context.Context) GetRestApiResultOutput {
	return o
}

// A JSON string containing response body.
func (o GetRestApiResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v GetRestApiResult) string { return v.Body }).(pulumi.StringOutput)
}

// A response status code.
func (o GetRestApiResultOutput) Code() pulumi.IntOutput {
	return o.ApplyT(func(v GetRestApiResult) int { return v.Code }).(pulumi.IntOutput)
}

func (o GetRestApiResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetRestApiResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// A JSON string containing response headers.
func (o GetRestApiResultOutput) Headers() pulumi.StringOutput {
	return o.ApplyT(func(v GetRestApiResult) string { return v.Headers }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRestApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRestApiResult) string { return v.Id }).(pulumi.StringOutput)
}

// A response status string.
func (o GetRestApiResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetRestApiResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRestApiResultOutput{})
}
