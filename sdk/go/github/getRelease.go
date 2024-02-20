// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a GitHub release in a specific repository.
func LookupRelease(ctx *pulumi.Context, args *LookupReleaseArgs, opts ...pulumi.InvokeOption) (*LookupReleaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReleaseResult
	err := ctx.Invoke("github:index/getRelease:getRelease", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRelease.
type LookupReleaseArgs struct {
	// Owner of the repository.
	Owner string `pulumi:"owner"`
	// ID of the release to retrieve. Must be specified when `retrieveBy` = `id`.
	ReleaseId *int `pulumi:"releaseId"`
	// Tag of the release to retrieve. Must be specified when `retrieveBy` = `tag`.
	ReleaseTag *string `pulumi:"releaseTag"`
	// Name of the repository to retrieve the release from.
	Repository string `pulumi:"repository"`
	// Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.
	RetrieveBy string `pulumi:"retrieveBy"`
}

// A collection of values returned by getRelease.
type LookupReleaseResult struct {
	// **Deprecated**: Use `assetsUrl` resource instead
	//
	// Deprecated: use assets_url instead
	AssertsUrl string `pulumi:"assertsUrl"`
	// Collection of assets for the release. Each asset conforms to the following schema:
	Assets []GetReleaseAsset `pulumi:"assets"`
	// URL of any associated assets with the release
	AssetsUrl string `pulumi:"assetsUrl"`
	// Contents of the description (body) of a release
	Body string `pulumi:"body"`
	// Date the asset was created
	CreatedAt string `pulumi:"createdAt"`
	// (`Boolean`) indicates whether the release is a draft
	Draft bool `pulumi:"draft"`
	// URL directing to detailed information on the release
	HtmlUrl string `pulumi:"htmlUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The file name of the asset
	Name  string `pulumi:"name"`
	Owner string `pulumi:"owner"`
	// (`Boolean`) indicates whether the release is a prerelease
	Prerelease bool `pulumi:"prerelease"`
	// Date of release publishing
	PublishedAt string `pulumi:"publishedAt"`
	// ID of release
	ReleaseId *int `pulumi:"releaseId"`
	// Tag of release
	ReleaseTag *string `pulumi:"releaseTag"`
	Repository string  `pulumi:"repository"`
	RetrieveBy string  `pulumi:"retrieveBy"`
	// Download URL of a specific release in `tar.gz` format
	TarballUrl string `pulumi:"tarballUrl"`
	// Commitish value that determines where the Git release is created from
	TargetCommitish string `pulumi:"targetCommitish"`
	// URL that can be used to upload Assets to the release
	UploadUrl string `pulumi:"uploadUrl"`
	// URL of the asset
	Url string `pulumi:"url"`
	// Download URL of a specific release in `zip` format
	ZipballUrl string `pulumi:"zipballUrl"`
}

func LookupReleaseOutput(ctx *pulumi.Context, args LookupReleaseOutputArgs, opts ...pulumi.InvokeOption) LookupReleaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReleaseResult, error) {
			args := v.(LookupReleaseArgs)
			r, err := LookupRelease(ctx, &args, opts...)
			var s LookupReleaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReleaseResultOutput)
}

// A collection of arguments for invoking getRelease.
type LookupReleaseOutputArgs struct {
	// Owner of the repository.
	Owner pulumi.StringInput `pulumi:"owner"`
	// ID of the release to retrieve. Must be specified when `retrieveBy` = `id`.
	ReleaseId pulumi.IntPtrInput `pulumi:"releaseId"`
	// Tag of the release to retrieve. Must be specified when `retrieveBy` = `tag`.
	ReleaseTag pulumi.StringPtrInput `pulumi:"releaseTag"`
	// Name of the repository to retrieve the release from.
	Repository pulumi.StringInput `pulumi:"repository"`
	// Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.
	RetrieveBy pulumi.StringInput `pulumi:"retrieveBy"`
}

func (LookupReleaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReleaseArgs)(nil)).Elem()
}

// A collection of values returned by getRelease.
type LookupReleaseResultOutput struct{ *pulumi.OutputState }

func (LookupReleaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReleaseResult)(nil)).Elem()
}

func (o LookupReleaseResultOutput) ToLookupReleaseResultOutput() LookupReleaseResultOutput {
	return o
}

func (o LookupReleaseResultOutput) ToLookupReleaseResultOutputWithContext(ctx context.Context) LookupReleaseResultOutput {
	return o
}

// **Deprecated**: Use `assetsUrl` resource instead
//
// Deprecated: use assets_url instead
func (o LookupReleaseResultOutput) AssertsUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.AssertsUrl }).(pulumi.StringOutput)
}

// Collection of assets for the release. Each asset conforms to the following schema:
func (o LookupReleaseResultOutput) Assets() GetReleaseAssetArrayOutput {
	return o.ApplyT(func(v LookupReleaseResult) []GetReleaseAsset { return v.Assets }).(GetReleaseAssetArrayOutput)
}

// URL of any associated assets with the release
func (o LookupReleaseResultOutput) AssetsUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.AssetsUrl }).(pulumi.StringOutput)
}

// Contents of the description (body) of a release
func (o LookupReleaseResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Body }).(pulumi.StringOutput)
}

// Date the asset was created
func (o LookupReleaseResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// (`Boolean`) indicates whether the release is a draft
func (o LookupReleaseResultOutput) Draft() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReleaseResult) bool { return v.Draft }).(pulumi.BoolOutput)
}

// URL directing to detailed information on the release
func (o LookupReleaseResultOutput) HtmlUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.HtmlUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupReleaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The file name of the asset
func (o LookupReleaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Owner }).(pulumi.StringOutput)
}

// (`Boolean`) indicates whether the release is a prerelease
func (o LookupReleaseResultOutput) Prerelease() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReleaseResult) bool { return v.Prerelease }).(pulumi.BoolOutput)
}

// Date of release publishing
func (o LookupReleaseResultOutput) PublishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.PublishedAt }).(pulumi.StringOutput)
}

// ID of release
func (o LookupReleaseResultOutput) ReleaseId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupReleaseResult) *int { return v.ReleaseId }).(pulumi.IntPtrOutput)
}

// Tag of release
func (o LookupReleaseResultOutput) ReleaseTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReleaseResult) *string { return v.ReleaseTag }).(pulumi.StringPtrOutput)
}

func (o LookupReleaseResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Repository }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) RetrieveBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.RetrieveBy }).(pulumi.StringOutput)
}

// Download URL of a specific release in `tar.gz` format
func (o LookupReleaseResultOutput) TarballUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.TarballUrl }).(pulumi.StringOutput)
}

// Commitish value that determines where the Git release is created from
func (o LookupReleaseResultOutput) TargetCommitish() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.TargetCommitish }).(pulumi.StringOutput)
}

// URL that can be used to upload Assets to the release
func (o LookupReleaseResultOutput) UploadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.UploadUrl }).(pulumi.StringOutput)
}

// URL of the asset
func (o LookupReleaseResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Url }).(pulumi.StringOutput)
}

// Download URL of a specific release in `zip` format
func (o LookupReleaseResultOutput) ZipballUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.ZipballUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReleaseResultOutput{})
}
