// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to manage repository allow list for existing GitHub Codespaces secrets within your GitHub organization.
//
// You must have write access to an organization secret to use this resource.
//
// This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.
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
//			repo, err := github.LookupRepository(ctx, &github.LookupRepositoryArgs{
//				FullName: pulumi.StringRef("my-org/repo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.NewCodespacesOrganizationSecretRepositories(ctx, "orgSecretRepos", &github.CodespacesOrganizationSecretRepositoriesArgs{
//				SecretName: pulumi.String("existing_secret_name"),
//				SelectedRepositoryIds: pulumi.IntArray{
//					*pulumi.Int(repo.RepoId),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # This resource can be imported using an ID made up of the secret name
//
// ```sh
//
//	$ pulumi import github:index/codespacesOrganizationSecretRepositories:CodespacesOrganizationSecretRepositories org_secret_repos existing_secret_name
//
// ```
type CodespacesOrganizationSecretRepositories struct {
	pulumi.CustomResourceState

	// Name of the existing secret
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayOutput `pulumi:"selectedRepositoryIds"`
}

// NewCodespacesOrganizationSecretRepositories registers a new resource with the given unique name, arguments, and options.
func NewCodespacesOrganizationSecretRepositories(ctx *pulumi.Context,
	name string, args *CodespacesOrganizationSecretRepositoriesArgs, opts ...pulumi.ResourceOption) (*CodespacesOrganizationSecretRepositories, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecretName == nil {
		return nil, errors.New("invalid value for required argument 'SecretName'")
	}
	if args.SelectedRepositoryIds == nil {
		return nil, errors.New("invalid value for required argument 'SelectedRepositoryIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CodespacesOrganizationSecretRepositories
	err := ctx.RegisterResource("github:index/codespacesOrganizationSecretRepositories:CodespacesOrganizationSecretRepositories", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCodespacesOrganizationSecretRepositories gets an existing CodespacesOrganizationSecretRepositories resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCodespacesOrganizationSecretRepositories(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodespacesOrganizationSecretRepositoriesState, opts ...pulumi.ResourceOption) (*CodespacesOrganizationSecretRepositories, error) {
	var resource CodespacesOrganizationSecretRepositories
	err := ctx.ReadResource("github:index/codespacesOrganizationSecretRepositories:CodespacesOrganizationSecretRepositories", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CodespacesOrganizationSecretRepositories resources.
type codespacesOrganizationSecretRepositoriesState struct {
	// Name of the existing secret
	SecretName *string `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
}

type CodespacesOrganizationSecretRepositoriesState struct {
	// Name of the existing secret
	SecretName pulumi.StringPtrInput
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayInput
}

func (CodespacesOrganizationSecretRepositoriesState) ElementType() reflect.Type {
	return reflect.TypeOf((*codespacesOrganizationSecretRepositoriesState)(nil)).Elem()
}

type codespacesOrganizationSecretRepositoriesArgs struct {
	// Name of the existing secret
	SecretName string `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
}

// The set of arguments for constructing a CodespacesOrganizationSecretRepositories resource.
type CodespacesOrganizationSecretRepositoriesArgs struct {
	// Name of the existing secret
	SecretName pulumi.StringInput
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayInput
}

func (CodespacesOrganizationSecretRepositoriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codespacesOrganizationSecretRepositoriesArgs)(nil)).Elem()
}

type CodespacesOrganizationSecretRepositoriesInput interface {
	pulumi.Input

	ToCodespacesOrganizationSecretRepositoriesOutput() CodespacesOrganizationSecretRepositoriesOutput
	ToCodespacesOrganizationSecretRepositoriesOutputWithContext(ctx context.Context) CodespacesOrganizationSecretRepositoriesOutput
}

func (*CodespacesOrganizationSecretRepositories) ElementType() reflect.Type {
	return reflect.TypeOf((**CodespacesOrganizationSecretRepositories)(nil)).Elem()
}

func (i *CodespacesOrganizationSecretRepositories) ToCodespacesOrganizationSecretRepositoriesOutput() CodespacesOrganizationSecretRepositoriesOutput {
	return i.ToCodespacesOrganizationSecretRepositoriesOutputWithContext(context.Background())
}

func (i *CodespacesOrganizationSecretRepositories) ToCodespacesOrganizationSecretRepositoriesOutputWithContext(ctx context.Context) CodespacesOrganizationSecretRepositoriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodespacesOrganizationSecretRepositoriesOutput)
}

// CodespacesOrganizationSecretRepositoriesArrayInput is an input type that accepts CodespacesOrganizationSecretRepositoriesArray and CodespacesOrganizationSecretRepositoriesArrayOutput values.
// You can construct a concrete instance of `CodespacesOrganizationSecretRepositoriesArrayInput` via:
//
//	CodespacesOrganizationSecretRepositoriesArray{ CodespacesOrganizationSecretRepositoriesArgs{...} }
type CodespacesOrganizationSecretRepositoriesArrayInput interface {
	pulumi.Input

	ToCodespacesOrganizationSecretRepositoriesArrayOutput() CodespacesOrganizationSecretRepositoriesArrayOutput
	ToCodespacesOrganizationSecretRepositoriesArrayOutputWithContext(context.Context) CodespacesOrganizationSecretRepositoriesArrayOutput
}

type CodespacesOrganizationSecretRepositoriesArray []CodespacesOrganizationSecretRepositoriesInput

func (CodespacesOrganizationSecretRepositoriesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CodespacesOrganizationSecretRepositories)(nil)).Elem()
}

func (i CodespacesOrganizationSecretRepositoriesArray) ToCodespacesOrganizationSecretRepositoriesArrayOutput() CodespacesOrganizationSecretRepositoriesArrayOutput {
	return i.ToCodespacesOrganizationSecretRepositoriesArrayOutputWithContext(context.Background())
}

func (i CodespacesOrganizationSecretRepositoriesArray) ToCodespacesOrganizationSecretRepositoriesArrayOutputWithContext(ctx context.Context) CodespacesOrganizationSecretRepositoriesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodespacesOrganizationSecretRepositoriesArrayOutput)
}

// CodespacesOrganizationSecretRepositoriesMapInput is an input type that accepts CodespacesOrganizationSecretRepositoriesMap and CodespacesOrganizationSecretRepositoriesMapOutput values.
// You can construct a concrete instance of `CodespacesOrganizationSecretRepositoriesMapInput` via:
//
//	CodespacesOrganizationSecretRepositoriesMap{ "key": CodespacesOrganizationSecretRepositoriesArgs{...} }
type CodespacesOrganizationSecretRepositoriesMapInput interface {
	pulumi.Input

	ToCodespacesOrganizationSecretRepositoriesMapOutput() CodespacesOrganizationSecretRepositoriesMapOutput
	ToCodespacesOrganizationSecretRepositoriesMapOutputWithContext(context.Context) CodespacesOrganizationSecretRepositoriesMapOutput
}

type CodespacesOrganizationSecretRepositoriesMap map[string]CodespacesOrganizationSecretRepositoriesInput

func (CodespacesOrganizationSecretRepositoriesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CodespacesOrganizationSecretRepositories)(nil)).Elem()
}

func (i CodespacesOrganizationSecretRepositoriesMap) ToCodespacesOrganizationSecretRepositoriesMapOutput() CodespacesOrganizationSecretRepositoriesMapOutput {
	return i.ToCodespacesOrganizationSecretRepositoriesMapOutputWithContext(context.Background())
}

func (i CodespacesOrganizationSecretRepositoriesMap) ToCodespacesOrganizationSecretRepositoriesMapOutputWithContext(ctx context.Context) CodespacesOrganizationSecretRepositoriesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodespacesOrganizationSecretRepositoriesMapOutput)
}

type CodespacesOrganizationSecretRepositoriesOutput struct{ *pulumi.OutputState }

func (CodespacesOrganizationSecretRepositoriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodespacesOrganizationSecretRepositories)(nil)).Elem()
}

func (o CodespacesOrganizationSecretRepositoriesOutput) ToCodespacesOrganizationSecretRepositoriesOutput() CodespacesOrganizationSecretRepositoriesOutput {
	return o
}

func (o CodespacesOrganizationSecretRepositoriesOutput) ToCodespacesOrganizationSecretRepositoriesOutputWithContext(ctx context.Context) CodespacesOrganizationSecretRepositoriesOutput {
	return o
}

// Name of the existing secret
func (o CodespacesOrganizationSecretRepositoriesOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *CodespacesOrganizationSecretRepositories) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// An array of repository ids that can access the organization secret.
func (o CodespacesOrganizationSecretRepositoriesOutput) SelectedRepositoryIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *CodespacesOrganizationSecretRepositories) pulumi.IntArrayOutput {
		return v.SelectedRepositoryIds
	}).(pulumi.IntArrayOutput)
}

type CodespacesOrganizationSecretRepositoriesArrayOutput struct{ *pulumi.OutputState }

func (CodespacesOrganizationSecretRepositoriesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CodespacesOrganizationSecretRepositories)(nil)).Elem()
}

func (o CodespacesOrganizationSecretRepositoriesArrayOutput) ToCodespacesOrganizationSecretRepositoriesArrayOutput() CodespacesOrganizationSecretRepositoriesArrayOutput {
	return o
}

func (o CodespacesOrganizationSecretRepositoriesArrayOutput) ToCodespacesOrganizationSecretRepositoriesArrayOutputWithContext(ctx context.Context) CodespacesOrganizationSecretRepositoriesArrayOutput {
	return o
}

func (o CodespacesOrganizationSecretRepositoriesArrayOutput) Index(i pulumi.IntInput) CodespacesOrganizationSecretRepositoriesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CodespacesOrganizationSecretRepositories {
		return vs[0].([]*CodespacesOrganizationSecretRepositories)[vs[1].(int)]
	}).(CodespacesOrganizationSecretRepositoriesOutput)
}

type CodespacesOrganizationSecretRepositoriesMapOutput struct{ *pulumi.OutputState }

func (CodespacesOrganizationSecretRepositoriesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CodespacesOrganizationSecretRepositories)(nil)).Elem()
}

func (o CodespacesOrganizationSecretRepositoriesMapOutput) ToCodespacesOrganizationSecretRepositoriesMapOutput() CodespacesOrganizationSecretRepositoriesMapOutput {
	return o
}

func (o CodespacesOrganizationSecretRepositoriesMapOutput) ToCodespacesOrganizationSecretRepositoriesMapOutputWithContext(ctx context.Context) CodespacesOrganizationSecretRepositoriesMapOutput {
	return o
}

func (o CodespacesOrganizationSecretRepositoriesMapOutput) MapIndex(k pulumi.StringInput) CodespacesOrganizationSecretRepositoriesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CodespacesOrganizationSecretRepositories {
		return vs[0].(map[string]*CodespacesOrganizationSecretRepositories)[vs[1].(string)]
	}).(CodespacesOrganizationSecretRepositoriesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CodespacesOrganizationSecretRepositoriesInput)(nil)).Elem(), &CodespacesOrganizationSecretRepositories{})
	pulumi.RegisterInputType(reflect.TypeOf((*CodespacesOrganizationSecretRepositoriesArrayInput)(nil)).Elem(), CodespacesOrganizationSecretRepositoriesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CodespacesOrganizationSecretRepositoriesMapInput)(nil)).Elem(), CodespacesOrganizationSecretRepositoriesMap{})
	pulumi.RegisterOutputType(CodespacesOrganizationSecretRepositoriesOutput{})
	pulumi.RegisterOutputType(CodespacesOrganizationSecretRepositoriesArrayOutput{})
	pulumi.RegisterOutputType(CodespacesOrganizationSecretRepositoriesMapOutput{})
}
