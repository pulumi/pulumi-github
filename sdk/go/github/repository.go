// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to create and manage repositories within your
// GitHub organization or personal account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v2/go/github"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.NewRepository(ctx, "example", &github.RepositoryArgs{
// 			Description: pulumi.String("My awesome codebase"),
// 			Template: &github.RepositoryTemplateArgs{
// 				Owner:      pulumi.String("github"),
// 				Repository: pulumi.String("terraform-module-template"),
// 			},
// 			Visibility: pulumi.String("public"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Repositories can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import github:index/repository:Repository terraform terraform
// ```
type Repository struct {
	pulumi.CustomResourceState

	// Set to `false` to disable merge commits on the repository.
	AllowMergeCommit pulumi.BoolPtrOutput `pulumi:"allowMergeCommit"`
	// Set to `false` to disable rebase merges on the repository.
	AllowRebaseMerge pulumi.BoolPtrOutput `pulumi:"allowRebaseMerge"`
	// Set to `false` to disable squash merges on the repository.
	AllowSquashMerge pulumi.BoolPtrOutput `pulumi:"allowSquashMerge"`
	// Set to `true` to archive the repository instead of deleting on destroy.
	ArchiveOnDestroy pulumi.BoolPtrOutput `pulumi:"archiveOnDestroy"`
	// Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
	Archived pulumi.BoolPtrOutput `pulumi:"archived"`
	// Set to `true` to produce an initial commit in the repository.
	AutoInit pulumi.BoolPtrOutput `pulumi:"autoInit"`
	// (Deprecated: Use `BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
	// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
	// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
	//
	// Deprecated: Use the github_branch_default resource instead
	DefaultBranch pulumi.StringOutput `pulumi:"defaultBranch"`
	// Automatically delete head branch after a pull request is merged. Defaults to `false`.
	DeleteBranchOnMerge pulumi.BoolPtrOutput `pulumi:"deleteBranchOnMerge"`
	// A description of the repository.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Etag        pulumi.StringOutput    `pulumi:"etag"`
	// A string of the form "orgname/reponame".
	FullName pulumi.StringOutput `pulumi:"fullName"`
	// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
	GitCloneUrl pulumi.StringOutput `pulumi:"gitCloneUrl"`
	// Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
	GitignoreTemplate pulumi.StringPtrOutput `pulumi:"gitignoreTemplate"`
	// Set to `true` to enable the (deprecated) downloads features on the repository.
	HasDownloads pulumi.BoolPtrOutput `pulumi:"hasDownloads"`
	// Set to `true` to enable the GitHub Issues features
	// on the repository.
	HasIssues pulumi.BoolPtrOutput `pulumi:"hasIssues"`
	// Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
	HasProjects pulumi.BoolPtrOutput `pulumi:"hasProjects"`
	// Set to `true` to enable the GitHub Wiki features on
	// the repository.
	HasWiki pulumi.BoolPtrOutput `pulumi:"hasWiki"`
	// URL of a page describing the project.
	HomepageUrl pulumi.StringPtrOutput `pulumi:"homepageUrl"`
	// URL to the repository on the web.
	HtmlUrl pulumi.StringOutput `pulumi:"htmlUrl"`
	// URL that can be provided to `git clone` to clone the repository via HTTPS.
	HttpCloneUrl pulumi.StringOutput `pulumi:"httpCloneUrl"`
	// Set to `true` to tell GitHub that this is a template repository.
	IsTemplate pulumi.BoolPtrOutput `pulumi:"isTemplate"`
	// Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
	LicenseTemplate pulumi.StringPtrOutput `pulumi:"licenseTemplate"`
	// The name of the repository.
	Name pulumi.StringOutput `pulumi:"name"`
	// GraphQL global node id for use with v4 API
	NodeId pulumi.StringOutput `pulumi:"nodeId"`
	// Set to `true` to create a private repository.
	// Repositories are created as public (e.g. open source) by default.
	//
	// Deprecated: use visibility instead
	Private pulumi.BoolOutput `pulumi:"private"`
	// Github ID for the repository
	RepoId pulumi.IntOutput `pulumi:"repoId"`
	// URL that can be provided to `git clone` to clone the repository via SSH.
	SshCloneUrl pulumi.StringOutput `pulumi:"sshCloneUrl"`
	// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
	SvnUrl pulumi.StringOutput `pulumi:"svnUrl"`
	// Use a template repository to create this resource. See Template Repositories below for details.
	Template RepositoryTemplatePtrOutput `pulumi:"template"`
	// The list of topics of the repository.
	Topics pulumi.StringArrayOutput `pulumi:"topics"`
	// Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
	// Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
	VulnerabilityAlerts pulumi.BoolPtrOutput `pulumi:"vulnerabilityAlerts"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		args = &RepositoryArgs{}
	}

	var resource Repository
	err := ctx.RegisterResource("github:index/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("github:index/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Set to `false` to disable merge commits on the repository.
	AllowMergeCommit *bool `pulumi:"allowMergeCommit"`
	// Set to `false` to disable rebase merges on the repository.
	AllowRebaseMerge *bool `pulumi:"allowRebaseMerge"`
	// Set to `false` to disable squash merges on the repository.
	AllowSquashMerge *bool `pulumi:"allowSquashMerge"`
	// Set to `true` to archive the repository instead of deleting on destroy.
	ArchiveOnDestroy *bool `pulumi:"archiveOnDestroy"`
	// Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
	Archived *bool `pulumi:"archived"`
	// Set to `true` to produce an initial commit in the repository.
	AutoInit *bool `pulumi:"autoInit"`
	// (Deprecated: Use `BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
	// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
	// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
	//
	// Deprecated: Use the github_branch_default resource instead
	DefaultBranch *string `pulumi:"defaultBranch"`
	// Automatically delete head branch after a pull request is merged. Defaults to `false`.
	DeleteBranchOnMerge *bool `pulumi:"deleteBranchOnMerge"`
	// A description of the repository.
	Description *string `pulumi:"description"`
	Etag        *string `pulumi:"etag"`
	// A string of the form "orgname/reponame".
	FullName *string `pulumi:"fullName"`
	// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
	GitCloneUrl *string `pulumi:"gitCloneUrl"`
	// Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
	GitignoreTemplate *string `pulumi:"gitignoreTemplate"`
	// Set to `true` to enable the (deprecated) downloads features on the repository.
	HasDownloads *bool `pulumi:"hasDownloads"`
	// Set to `true` to enable the GitHub Issues features
	// on the repository.
	HasIssues *bool `pulumi:"hasIssues"`
	// Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
	HasProjects *bool `pulumi:"hasProjects"`
	// Set to `true` to enable the GitHub Wiki features on
	// the repository.
	HasWiki *bool `pulumi:"hasWiki"`
	// URL of a page describing the project.
	HomepageUrl *string `pulumi:"homepageUrl"`
	// URL to the repository on the web.
	HtmlUrl *string `pulumi:"htmlUrl"`
	// URL that can be provided to `git clone` to clone the repository via HTTPS.
	HttpCloneUrl *string `pulumi:"httpCloneUrl"`
	// Set to `true` to tell GitHub that this is a template repository.
	IsTemplate *bool `pulumi:"isTemplate"`
	// Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
	LicenseTemplate *string `pulumi:"licenseTemplate"`
	// The name of the repository.
	Name *string `pulumi:"name"`
	// GraphQL global node id for use with v4 API
	NodeId *string `pulumi:"nodeId"`
	// Set to `true` to create a private repository.
	// Repositories are created as public (e.g. open source) by default.
	//
	// Deprecated: use visibility instead
	Private *bool `pulumi:"private"`
	// Github ID for the repository
	RepoId *int `pulumi:"repoId"`
	// URL that can be provided to `git clone` to clone the repository via SSH.
	SshCloneUrl *string `pulumi:"sshCloneUrl"`
	// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
	SvnUrl *string `pulumi:"svnUrl"`
	// Use a template repository to create this resource. See Template Repositories below for details.
	Template *RepositoryTemplate `pulumi:"template"`
	// The list of topics of the repository.
	Topics []string `pulumi:"topics"`
	// Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
	Visibility *string `pulumi:"visibility"`
	// Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
	VulnerabilityAlerts *bool `pulumi:"vulnerabilityAlerts"`
}

type RepositoryState struct {
	// Set to `false` to disable merge commits on the repository.
	AllowMergeCommit pulumi.BoolPtrInput
	// Set to `false` to disable rebase merges on the repository.
	AllowRebaseMerge pulumi.BoolPtrInput
	// Set to `false` to disable squash merges on the repository.
	AllowSquashMerge pulumi.BoolPtrInput
	// Set to `true` to archive the repository instead of deleting on destroy.
	ArchiveOnDestroy pulumi.BoolPtrInput
	// Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
	Archived pulumi.BoolPtrInput
	// Set to `true` to produce an initial commit in the repository.
	AutoInit pulumi.BoolPtrInput
	// (Deprecated: Use `BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
	// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
	// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
	//
	// Deprecated: Use the github_branch_default resource instead
	DefaultBranch pulumi.StringPtrInput
	// Automatically delete head branch after a pull request is merged. Defaults to `false`.
	DeleteBranchOnMerge pulumi.BoolPtrInput
	// A description of the repository.
	Description pulumi.StringPtrInput
	Etag        pulumi.StringPtrInput
	// A string of the form "orgname/reponame".
	FullName pulumi.StringPtrInput
	// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
	GitCloneUrl pulumi.StringPtrInput
	// Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
	GitignoreTemplate pulumi.StringPtrInput
	// Set to `true` to enable the (deprecated) downloads features on the repository.
	HasDownloads pulumi.BoolPtrInput
	// Set to `true` to enable the GitHub Issues features
	// on the repository.
	HasIssues pulumi.BoolPtrInput
	// Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
	HasProjects pulumi.BoolPtrInput
	// Set to `true` to enable the GitHub Wiki features on
	// the repository.
	HasWiki pulumi.BoolPtrInput
	// URL of a page describing the project.
	HomepageUrl pulumi.StringPtrInput
	// URL to the repository on the web.
	HtmlUrl pulumi.StringPtrInput
	// URL that can be provided to `git clone` to clone the repository via HTTPS.
	HttpCloneUrl pulumi.StringPtrInput
	// Set to `true` to tell GitHub that this is a template repository.
	IsTemplate pulumi.BoolPtrInput
	// Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
	LicenseTemplate pulumi.StringPtrInput
	// The name of the repository.
	Name pulumi.StringPtrInput
	// GraphQL global node id for use with v4 API
	NodeId pulumi.StringPtrInput
	// Set to `true` to create a private repository.
	// Repositories are created as public (e.g. open source) by default.
	//
	// Deprecated: use visibility instead
	Private pulumi.BoolPtrInput
	// Github ID for the repository
	RepoId pulumi.IntPtrInput
	// URL that can be provided to `git clone` to clone the repository via SSH.
	SshCloneUrl pulumi.StringPtrInput
	// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
	SvnUrl pulumi.StringPtrInput
	// Use a template repository to create this resource. See Template Repositories below for details.
	Template RepositoryTemplatePtrInput
	// The list of topics of the repository.
	Topics pulumi.StringArrayInput
	// Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
	Visibility pulumi.StringPtrInput
	// Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
	VulnerabilityAlerts pulumi.BoolPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Set to `false` to disable merge commits on the repository.
	AllowMergeCommit *bool `pulumi:"allowMergeCommit"`
	// Set to `false` to disable rebase merges on the repository.
	AllowRebaseMerge *bool `pulumi:"allowRebaseMerge"`
	// Set to `false` to disable squash merges on the repository.
	AllowSquashMerge *bool `pulumi:"allowSquashMerge"`
	// Set to `true` to archive the repository instead of deleting on destroy.
	ArchiveOnDestroy *bool `pulumi:"archiveOnDestroy"`
	// Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
	Archived *bool `pulumi:"archived"`
	// Set to `true` to produce an initial commit in the repository.
	AutoInit *bool `pulumi:"autoInit"`
	// (Deprecated: Use `BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
	// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
	// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
	//
	// Deprecated: Use the github_branch_default resource instead
	DefaultBranch *string `pulumi:"defaultBranch"`
	// Automatically delete head branch after a pull request is merged. Defaults to `false`.
	DeleteBranchOnMerge *bool `pulumi:"deleteBranchOnMerge"`
	// A description of the repository.
	Description *string `pulumi:"description"`
	// Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
	GitignoreTemplate *string `pulumi:"gitignoreTemplate"`
	// Set to `true` to enable the (deprecated) downloads features on the repository.
	HasDownloads *bool `pulumi:"hasDownloads"`
	// Set to `true` to enable the GitHub Issues features
	// on the repository.
	HasIssues *bool `pulumi:"hasIssues"`
	// Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
	HasProjects *bool `pulumi:"hasProjects"`
	// Set to `true` to enable the GitHub Wiki features on
	// the repository.
	HasWiki *bool `pulumi:"hasWiki"`
	// URL of a page describing the project.
	HomepageUrl *string `pulumi:"homepageUrl"`
	// Set to `true` to tell GitHub that this is a template repository.
	IsTemplate *bool `pulumi:"isTemplate"`
	// Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
	LicenseTemplate *string `pulumi:"licenseTemplate"`
	// The name of the repository.
	Name *string `pulumi:"name"`
	// Set to `true` to create a private repository.
	// Repositories are created as public (e.g. open source) by default.
	//
	// Deprecated: use visibility instead
	Private *bool `pulumi:"private"`
	// Use a template repository to create this resource. See Template Repositories below for details.
	Template *RepositoryTemplate `pulumi:"template"`
	// The list of topics of the repository.
	Topics []string `pulumi:"topics"`
	// Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
	Visibility *string `pulumi:"visibility"`
	// Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
	VulnerabilityAlerts *bool `pulumi:"vulnerabilityAlerts"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Set to `false` to disable merge commits on the repository.
	AllowMergeCommit pulumi.BoolPtrInput
	// Set to `false` to disable rebase merges on the repository.
	AllowRebaseMerge pulumi.BoolPtrInput
	// Set to `false` to disable squash merges on the repository.
	AllowSquashMerge pulumi.BoolPtrInput
	// Set to `true` to archive the repository instead of deleting on destroy.
	ArchiveOnDestroy pulumi.BoolPtrInput
	// Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
	Archived pulumi.BoolPtrInput
	// Set to `true` to produce an initial commit in the repository.
	AutoInit pulumi.BoolPtrInput
	// (Deprecated: Use `BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
	// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
	// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
	//
	// Deprecated: Use the github_branch_default resource instead
	DefaultBranch pulumi.StringPtrInput
	// Automatically delete head branch after a pull request is merged. Defaults to `false`.
	DeleteBranchOnMerge pulumi.BoolPtrInput
	// A description of the repository.
	Description pulumi.StringPtrInput
	// Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
	GitignoreTemplate pulumi.StringPtrInput
	// Set to `true` to enable the (deprecated) downloads features on the repository.
	HasDownloads pulumi.BoolPtrInput
	// Set to `true` to enable the GitHub Issues features
	// on the repository.
	HasIssues pulumi.BoolPtrInput
	// Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
	HasProjects pulumi.BoolPtrInput
	// Set to `true` to enable the GitHub Wiki features on
	// the repository.
	HasWiki pulumi.BoolPtrInput
	// URL of a page describing the project.
	HomepageUrl pulumi.StringPtrInput
	// Set to `true` to tell GitHub that this is a template repository.
	IsTemplate pulumi.BoolPtrInput
	// Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
	LicenseTemplate pulumi.StringPtrInput
	// The name of the repository.
	Name pulumi.StringPtrInput
	// Set to `true` to create a private repository.
	// Repositories are created as public (e.g. open source) by default.
	//
	// Deprecated: use visibility instead
	Private pulumi.BoolPtrInput
	// Use a template repository to create this resource. See Template Repositories below for details.
	Template RepositoryTemplatePtrInput
	// The list of topics of the repository.
	Topics pulumi.StringArrayInput
	// Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
	Visibility pulumi.StringPtrInput
	// Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
	VulnerabilityAlerts pulumi.BoolPtrInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (Repository) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil)).Elem()
}

func (i Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

type RepositoryOutput struct {
	*pulumi.OutputState
}

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryOutput)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RepositoryOutput{})
}
