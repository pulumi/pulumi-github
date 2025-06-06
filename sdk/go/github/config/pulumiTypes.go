// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type AppAuth struct {
	// The GitHub App ID.
	Id string `pulumi:"id"`
	// The GitHub App installation instance ID.
	InstallationId string `pulumi:"installationId"`
	// The GitHub App PEM file contents.
	PemFile string `pulumi:"pemFile"`
}

// AppAuthInput is an input type that accepts AppAuthArgs and AppAuthOutput values.
// You can construct a concrete instance of `AppAuthInput` via:
//
//	AppAuthArgs{...}
type AppAuthInput interface {
	pulumi.Input

	ToAppAuthOutput() AppAuthOutput
	ToAppAuthOutputWithContext(context.Context) AppAuthOutput
}

type AppAuthArgs struct {
	// The GitHub App ID.
	Id pulumi.StringInput `pulumi:"id"`
	// The GitHub App installation instance ID.
	InstallationId pulumi.StringInput `pulumi:"installationId"`
	// The GitHub App PEM file contents.
	PemFile pulumi.StringInput `pulumi:"pemFile"`
}

func (AppAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAuth)(nil)).Elem()
}

func (i AppAuthArgs) ToAppAuthOutput() AppAuthOutput {
	return i.ToAppAuthOutputWithContext(context.Background())
}

func (i AppAuthArgs) ToAppAuthOutputWithContext(ctx context.Context) AppAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAuthOutput)
}

type AppAuthOutput struct{ *pulumi.OutputState }

func (AppAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAuth)(nil)).Elem()
}

func (o AppAuthOutput) ToAppAuthOutput() AppAuthOutput {
	return o
}

func (o AppAuthOutput) ToAppAuthOutputWithContext(ctx context.Context) AppAuthOutput {
	return o
}

// The GitHub App ID.
func (o AppAuthOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AppAuth) string { return v.Id }).(pulumi.StringOutput)
}

// The GitHub App installation instance ID.
func (o AppAuthOutput) InstallationId() pulumi.StringOutput {
	return o.ApplyT(func(v AppAuth) string { return v.InstallationId }).(pulumi.StringOutput)
}

// The GitHub App PEM file contents.
func (o AppAuthOutput) PemFile() pulumi.StringOutput {
	return o.ApplyT(func(v AppAuth) string { return v.PemFile }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppAuthInput)(nil)).Elem(), AppAuthArgs{})
	pulumi.RegisterOutputType(AppAuthOutput{})
}
