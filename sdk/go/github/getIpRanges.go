// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about GitHub's IP addresses.
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
//			_, err := github.GetIpRanges(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIpRanges(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIpRangesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpRangesResult
	err := ctx.Invoke("github:index/getIpRanges:getIpRanges", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIpRanges.
type GetIpRangesResult struct {
	// An array of IP addresses in CIDR format specifying the addresses that incoming requests from GitHub actions will originate from.
	Actions []string `pulumi:"actions"`
	// A subset of the `actions` array that contains IP addresses in IPv4 CIDR format.
	ActionsIpv4s []string `pulumi:"actionsIpv4s"`
	// A subset of the `actions` array that contains IP addresses in IPv6 CIDR format.
	ActionsIpv6s []string `pulumi:"actionsIpv6s"`
	// A subset of the `api` array that contains IP addresses in IPv4 CIDR format.
	ApiIpv4s []string `pulumi:"apiIpv4s"`
	// A subset of the `api` array that contains IP addresses in IPv6 CIDR format.
	ApiIpv6s []string `pulumi:"apiIpv6s"`
	// An Array of IP addresses in CIDR format for the GitHub API.
	Apis []string `pulumi:"apis"`
	// A subset of the `dependabot` array that contains IP addresses in IPv4 CIDR format.
	DependabotIpv4s []string `pulumi:"dependabotIpv4s"`
	// A subset of the `dependabot` array that contains IP addresses in IPv6 CIDR format.
	DependabotIpv6s []string `pulumi:"dependabotIpv6s"`
	// An array of IP addresses in CIDR format specifying the A records for dependabot.
	Dependabots []string `pulumi:"dependabots"`
	// A subset of the `git` array that contains IP addresses in IPv4 CIDR format.
	GitIpv4s []string `pulumi:"gitIpv4s"`
	// A subset of the `git` array that contains IP addresses in IPv6 CIDR format.
	GitIpv6s []string `pulumi:"gitIpv6s"`
	// An Array of IP addresses in CIDR format specifying the Git servers.
	Gits []string `pulumi:"gits"`
	// An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
	Hooks []string `pulumi:"hooks"`
	// A subset of the `hooks` array that contains IP addresses in IPv4 CIDR format.
	HooksIpv4s []string `pulumi:"hooksIpv4s"`
	// A subset of the `hooks` array that contains IP addresses in IPv6 CIDR format.
	HooksIpv6s []string `pulumi:"hooksIpv6s"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A subset of the `importer` array that contains IP addresses in IPv4 CIDR format.
	ImporterIpv4s []string `pulumi:"importerIpv4s"`
	// A subset of the `importer` array that contains IP addresses in IPv6 CIDR format.
	ImporterIpv6s []string `pulumi:"importerIpv6s"`
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
	Importers []string `pulumi:"importers"`
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Packages.
	Packages []string `pulumi:"packages"`
	// A subset of the `packages` array that contains IP addresses in IPv4 CIDR format.
	PackagesIpv4s []string `pulumi:"packagesIpv4s"`
	// A subset of the `packages` array that contains IP addresses in IPv6 CIDR format.
	PackagesIpv6s []string `pulumi:"packagesIpv6s"`
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
	Pages []string `pulumi:"pages"`
	// A subset of the `pages` array that contains IP addresses in IPv4 CIDR format.
	PagesIpv4s []string `pulumi:"pagesIpv4s"`
	// A subset of the `pages` array that contains IP addresses in IPv6 CIDR format.
	PagesIpv6s []string `pulumi:"pagesIpv6s"`
	// A subset of the `web` array that contains IP addresses in IPv4 CIDR format.
	WebIpv4s []string `pulumi:"webIpv4s"`
	// A subset of the `web` array that contains IP addresses in IPv6 CIDR format.
	WebIpv6s []string `pulumi:"webIpv6s"`
	// An Array of IP addresses in CIDR format for GitHub Web.
	Webs []string `pulumi:"webs"`
}

func GetIpRangesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetIpRangesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetIpRangesResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("github:index/getIpRanges:getIpRanges", nil, GetIpRangesResultOutput{}, options).(GetIpRangesResultOutput), nil
	}).(GetIpRangesResultOutput)
}

// A collection of values returned by getIpRanges.
type GetIpRangesResultOutput struct{ *pulumi.OutputState }

func (GetIpRangesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpRangesResult)(nil)).Elem()
}

func (o GetIpRangesResultOutput) ToGetIpRangesResultOutput() GetIpRangesResultOutput {
	return o
}

func (o GetIpRangesResultOutput) ToGetIpRangesResultOutputWithContext(ctx context.Context) GetIpRangesResultOutput {
	return o
}

// An array of IP addresses in CIDR format specifying the addresses that incoming requests from GitHub actions will originate from.
func (o GetIpRangesResultOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

// A subset of the `actions` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) ActionsIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.ActionsIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `actions` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) ActionsIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.ActionsIpv6s }).(pulumi.StringArrayOutput)
}

// A subset of the `api` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) ApiIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.ApiIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `api` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) ApiIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.ApiIpv6s }).(pulumi.StringArrayOutput)
}

// An Array of IP addresses in CIDR format for the GitHub API.
func (o GetIpRangesResultOutput) Apis() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Apis }).(pulumi.StringArrayOutput)
}

// A subset of the `dependabot` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) DependabotIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.DependabotIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `dependabot` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) DependabotIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.DependabotIpv6s }).(pulumi.StringArrayOutput)
}

// An array of IP addresses in CIDR format specifying the A records for dependabot.
func (o GetIpRangesResultOutput) Dependabots() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Dependabots }).(pulumi.StringArrayOutput)
}

// A subset of the `git` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) GitIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.GitIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `git` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) GitIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.GitIpv6s }).(pulumi.StringArrayOutput)
}

// An Array of IP addresses in CIDR format specifying the Git servers.
func (o GetIpRangesResultOutput) Gits() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Gits }).(pulumi.StringArrayOutput)
}

// An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
func (o GetIpRangesResultOutput) Hooks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Hooks }).(pulumi.StringArrayOutput)
}

// A subset of the `hooks` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) HooksIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.HooksIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `hooks` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) HooksIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.HooksIpv6s }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpRangesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpRangesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A subset of the `importer` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) ImporterIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.ImporterIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `importer` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) ImporterIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.ImporterIpv6s }).(pulumi.StringArrayOutput)
}

// An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
func (o GetIpRangesResultOutput) Importers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Importers }).(pulumi.StringArrayOutput)
}

// An Array of IP addresses in CIDR format specifying the A records for GitHub Packages.
func (o GetIpRangesResultOutput) Packages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Packages }).(pulumi.StringArrayOutput)
}

// A subset of the `packages` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) PackagesIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.PackagesIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `packages` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) PackagesIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.PackagesIpv6s }).(pulumi.StringArrayOutput)
}

// An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
func (o GetIpRangesResultOutput) Pages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Pages }).(pulumi.StringArrayOutput)
}

// A subset of the `pages` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) PagesIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.PagesIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `pages` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) PagesIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.PagesIpv6s }).(pulumi.StringArrayOutput)
}

// A subset of the `web` array that contains IP addresses in IPv4 CIDR format.
func (o GetIpRangesResultOutput) WebIpv4s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.WebIpv4s }).(pulumi.StringArrayOutput)
}

// A subset of the `web` array that contains IP addresses in IPv6 CIDR format.
func (o GetIpRangesResultOutput) WebIpv6s() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.WebIpv6s }).(pulumi.StringArrayOutput)
}

// An Array of IP addresses in CIDR format for GitHub Web.
func (o GetIpRangesResultOutput) Webs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRangesResult) []string { return v.Webs }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpRangesResultOutput{})
}
