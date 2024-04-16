// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage files within a
// GitHub repository.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			foo, err := github.NewRepository(ctx, "foo", &github.RepositoryArgs{
//				Name:     pulumi.String("tf-acc-test-%s"),
//				AutoInit: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryFile(ctx, "foo", &github.RepositoryFileArgs{
//				Repository:        foo.Name,
//				Branch:            pulumi.String("main"),
//				File:              pulumi.String(".gitignore"),
//				Content:           pulumi.String("**/*.tfstate"),
//				CommitMessage:     pulumi.String("Managed by Terraform"),
//				CommitAuthor:      pulumi.String("Terraform User"),
//				CommitEmail:       pulumi.String("terraform@example.com"),
//				OverwriteOnCreate: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Repository files can be imported using a combination of the `repo` and `file`, e.g.
//
// ```sh
// $ pulumi import github:index/repositoryFile:RepositoryFile gitignore example/.gitignore
// ```
// To import a file from a branch other than the default branch, append `:` and the branch name, e.g.
//
// ```sh
// $ pulumi import github:index/repositoryFile:RepositoryFile gitignore example/.gitignore:dev
// ```
type RepositoryFile struct {
	pulumi.CustomResourceState

	// Git branch (defaults to the repository's default branch).
	// The branch must already exist, it will not be created if it does not already exist.
	Branch pulumi.StringPtrOutput `pulumi:"branch"`
	// Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
	CommitAuthor pulumi.StringPtrOutput `pulumi:"commitAuthor"`
	// Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
	CommitEmail pulumi.StringPtrOutput `pulumi:"commitEmail"`
	// The commit message when creating, updating or deleting the managed file.
	CommitMessage pulumi.StringOutput `pulumi:"commitMessage"`
	// The SHA of the commit that modified the file.
	CommitSha pulumi.StringOutput `pulumi:"commitSha"`
	// The file content.
	Content pulumi.StringOutput `pulumi:"content"`
	// The path of the file to manage.
	File pulumi.StringOutput `pulumi:"file"`
	// Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
	OverwriteOnCreate pulumi.BoolPtrOutput `pulumi:"overwriteOnCreate"`
	// The name of the commit/branch/tag.
	Ref pulumi.StringOutput `pulumi:"ref"`
	// The repository to create the file in.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The SHA blob of the file.
	Sha pulumi.StringOutput `pulumi:"sha"`
}

// NewRepositoryFile registers a new resource with the given unique name, arguments, and options.
func NewRepositoryFile(ctx *pulumi.Context,
	name string, args *RepositoryFileArgs, opts ...pulumi.ResourceOption) (*RepositoryFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.File == nil {
		return nil, errors.New("invalid value for required argument 'File'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryFile
	err := ctx.RegisterResource("github:index/repositoryFile:RepositoryFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryFile gets an existing RepositoryFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryFileState, opts ...pulumi.ResourceOption) (*RepositoryFile, error) {
	var resource RepositoryFile
	err := ctx.ReadResource("github:index/repositoryFile:RepositoryFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryFile resources.
type repositoryFileState struct {
	// Git branch (defaults to the repository's default branch).
	// The branch must already exist, it will not be created if it does not already exist.
	Branch *string `pulumi:"branch"`
	// Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
	CommitAuthor *string `pulumi:"commitAuthor"`
	// Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
	CommitEmail *string `pulumi:"commitEmail"`
	// The commit message when creating, updating or deleting the managed file.
	CommitMessage *string `pulumi:"commitMessage"`
	// The SHA of the commit that modified the file.
	CommitSha *string `pulumi:"commitSha"`
	// The file content.
	Content *string `pulumi:"content"`
	// The path of the file to manage.
	File *string `pulumi:"file"`
	// Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
	OverwriteOnCreate *bool `pulumi:"overwriteOnCreate"`
	// The name of the commit/branch/tag.
	Ref *string `pulumi:"ref"`
	// The repository to create the file in.
	Repository *string `pulumi:"repository"`
	// The SHA blob of the file.
	Sha *string `pulumi:"sha"`
}

type RepositoryFileState struct {
	// Git branch (defaults to the repository's default branch).
	// The branch must already exist, it will not be created if it does not already exist.
	Branch pulumi.StringPtrInput
	// Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
	CommitAuthor pulumi.StringPtrInput
	// Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
	CommitEmail pulumi.StringPtrInput
	// The commit message when creating, updating or deleting the managed file.
	CommitMessage pulumi.StringPtrInput
	// The SHA of the commit that modified the file.
	CommitSha pulumi.StringPtrInput
	// The file content.
	Content pulumi.StringPtrInput
	// The path of the file to manage.
	File pulumi.StringPtrInput
	// Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
	OverwriteOnCreate pulumi.BoolPtrInput
	// The name of the commit/branch/tag.
	Ref pulumi.StringPtrInput
	// The repository to create the file in.
	Repository pulumi.StringPtrInput
	// The SHA blob of the file.
	Sha pulumi.StringPtrInput
}

func (RepositoryFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryFileState)(nil)).Elem()
}

type repositoryFileArgs struct {
	// Git branch (defaults to the repository's default branch).
	// The branch must already exist, it will not be created if it does not already exist.
	Branch *string `pulumi:"branch"`
	// Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
	CommitAuthor *string `pulumi:"commitAuthor"`
	// Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
	CommitEmail *string `pulumi:"commitEmail"`
	// The commit message when creating, updating or deleting the managed file.
	CommitMessage *string `pulumi:"commitMessage"`
	// The file content.
	Content string `pulumi:"content"`
	// The path of the file to manage.
	File string `pulumi:"file"`
	// Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
	OverwriteOnCreate *bool `pulumi:"overwriteOnCreate"`
	// The repository to create the file in.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryFile resource.
type RepositoryFileArgs struct {
	// Git branch (defaults to the repository's default branch).
	// The branch must already exist, it will not be created if it does not already exist.
	Branch pulumi.StringPtrInput
	// Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
	CommitAuthor pulumi.StringPtrInput
	// Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
	CommitEmail pulumi.StringPtrInput
	// The commit message when creating, updating or deleting the managed file.
	CommitMessage pulumi.StringPtrInput
	// The file content.
	Content pulumi.StringInput
	// The path of the file to manage.
	File pulumi.StringInput
	// Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
	OverwriteOnCreate pulumi.BoolPtrInput
	// The repository to create the file in.
	Repository pulumi.StringInput
}

func (RepositoryFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryFileArgs)(nil)).Elem()
}

type RepositoryFileInput interface {
	pulumi.Input

	ToRepositoryFileOutput() RepositoryFileOutput
	ToRepositoryFileOutputWithContext(ctx context.Context) RepositoryFileOutput
}

func (*RepositoryFile) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryFile)(nil)).Elem()
}

func (i *RepositoryFile) ToRepositoryFileOutput() RepositoryFileOutput {
	return i.ToRepositoryFileOutputWithContext(context.Background())
}

func (i *RepositoryFile) ToRepositoryFileOutputWithContext(ctx context.Context) RepositoryFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryFileOutput)
}

// RepositoryFileArrayInput is an input type that accepts RepositoryFileArray and RepositoryFileArrayOutput values.
// You can construct a concrete instance of `RepositoryFileArrayInput` via:
//
//	RepositoryFileArray{ RepositoryFileArgs{...} }
type RepositoryFileArrayInput interface {
	pulumi.Input

	ToRepositoryFileArrayOutput() RepositoryFileArrayOutput
	ToRepositoryFileArrayOutputWithContext(context.Context) RepositoryFileArrayOutput
}

type RepositoryFileArray []RepositoryFileInput

func (RepositoryFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryFile)(nil)).Elem()
}

func (i RepositoryFileArray) ToRepositoryFileArrayOutput() RepositoryFileArrayOutput {
	return i.ToRepositoryFileArrayOutputWithContext(context.Background())
}

func (i RepositoryFileArray) ToRepositoryFileArrayOutputWithContext(ctx context.Context) RepositoryFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryFileArrayOutput)
}

// RepositoryFileMapInput is an input type that accepts RepositoryFileMap and RepositoryFileMapOutput values.
// You can construct a concrete instance of `RepositoryFileMapInput` via:
//
//	RepositoryFileMap{ "key": RepositoryFileArgs{...} }
type RepositoryFileMapInput interface {
	pulumi.Input

	ToRepositoryFileMapOutput() RepositoryFileMapOutput
	ToRepositoryFileMapOutputWithContext(context.Context) RepositoryFileMapOutput
}

type RepositoryFileMap map[string]RepositoryFileInput

func (RepositoryFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryFile)(nil)).Elem()
}

func (i RepositoryFileMap) ToRepositoryFileMapOutput() RepositoryFileMapOutput {
	return i.ToRepositoryFileMapOutputWithContext(context.Background())
}

func (i RepositoryFileMap) ToRepositoryFileMapOutputWithContext(ctx context.Context) RepositoryFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryFileMapOutput)
}

type RepositoryFileOutput struct{ *pulumi.OutputState }

func (RepositoryFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryFile)(nil)).Elem()
}

func (o RepositoryFileOutput) ToRepositoryFileOutput() RepositoryFileOutput {
	return o
}

func (o RepositoryFileOutput) ToRepositoryFileOutputWithContext(ctx context.Context) RepositoryFileOutput {
	return o
}

// Git branch (defaults to the repository's default branch).
// The branch must already exist, it will not be created if it does not already exist.
func (o RepositoryFileOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

// Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
func (o RepositoryFileOutput) CommitAuthor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringPtrOutput { return v.CommitAuthor }).(pulumi.StringPtrOutput)
}

// Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
func (o RepositoryFileOutput) CommitEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringPtrOutput { return v.CommitEmail }).(pulumi.StringPtrOutput)
}

// The commit message when creating, updating or deleting the managed file.
func (o RepositoryFileOutput) CommitMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.CommitMessage }).(pulumi.StringOutput)
}

// The SHA of the commit that modified the file.
func (o RepositoryFileOutput) CommitSha() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.CommitSha }).(pulumi.StringOutput)
}

// The file content.
func (o RepositoryFileOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// The path of the file to manage.
func (o RepositoryFileOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.File }).(pulumi.StringOutput)
}

// Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
func (o RepositoryFileOutput) OverwriteOnCreate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.BoolPtrOutput { return v.OverwriteOnCreate }).(pulumi.BoolPtrOutput)
}

// The name of the commit/branch/tag.
func (o RepositoryFileOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Ref }).(pulumi.StringOutput)
}

// The repository to create the file in.
func (o RepositoryFileOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The SHA blob of the file.
func (o RepositoryFileOutput) Sha() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Sha }).(pulumi.StringOutput)
}

type RepositoryFileArrayOutput struct{ *pulumi.OutputState }

func (RepositoryFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryFile)(nil)).Elem()
}

func (o RepositoryFileArrayOutput) ToRepositoryFileArrayOutput() RepositoryFileArrayOutput {
	return o
}

func (o RepositoryFileArrayOutput) ToRepositoryFileArrayOutputWithContext(ctx context.Context) RepositoryFileArrayOutput {
	return o
}

func (o RepositoryFileArrayOutput) Index(i pulumi.IntInput) RepositoryFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryFile {
		return vs[0].([]*RepositoryFile)[vs[1].(int)]
	}).(RepositoryFileOutput)
}

type RepositoryFileMapOutput struct{ *pulumi.OutputState }

func (RepositoryFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryFile)(nil)).Elem()
}

func (o RepositoryFileMapOutput) ToRepositoryFileMapOutput() RepositoryFileMapOutput {
	return o
}

func (o RepositoryFileMapOutput) ToRepositoryFileMapOutputWithContext(ctx context.Context) RepositoryFileMapOutput {
	return o
}

func (o RepositoryFileMapOutput) MapIndex(k pulumi.StringInput) RepositoryFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryFile {
		return vs[0].(map[string]*RepositoryFile)[vs[1].(string)]
	}).(RepositoryFileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryFileInput)(nil)).Elem(), &RepositoryFile{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryFileArrayInput)(nil)).Elem(), RepositoryFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryFileMapInput)(nil)).Elem(), RepositoryFileMap{})
	pulumi.RegisterOutputType(RepositoryFileOutput{})
	pulumi.RegisterOutputType(RepositoryFileArrayOutput{})
	pulumi.RegisterOutputType(RepositoryFileMapOutput{})
}
