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

// This resource allows you to manage repository allow list for existing GitHub Actions secrets within your GitHub organization.
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
//			_, err = github.NewActionsOrganizationSecretRepositories(ctx, "orgSecretRepos", &github.ActionsOrganizationSecretRepositoriesArgs{
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
// This resource can be imported using an ID made up of the secret name:
//
// ```sh
//
//	$ pulumi import github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories test_secret_repos test_secret_name
//
// ```
type ActionsOrganizationSecretRepositories struct {
	pulumi.CustomResourceState

	// Name of the existing secret
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayOutput `pulumi:"selectedRepositoryIds"`
}

// NewActionsOrganizationSecretRepositories registers a new resource with the given unique name, arguments, and options.
func NewActionsOrganizationSecretRepositories(ctx *pulumi.Context,
	name string, args *ActionsOrganizationSecretRepositoriesArgs, opts ...pulumi.ResourceOption) (*ActionsOrganizationSecretRepositories, error) {
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
	var resource ActionsOrganizationSecretRepositories
	err := ctx.RegisterResource("github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionsOrganizationSecretRepositories gets an existing ActionsOrganizationSecretRepositories resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionsOrganizationSecretRepositories(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionsOrganizationSecretRepositoriesState, opts ...pulumi.ResourceOption) (*ActionsOrganizationSecretRepositories, error) {
	var resource ActionsOrganizationSecretRepositories
	err := ctx.ReadResource("github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionsOrganizationSecretRepositories resources.
type actionsOrganizationSecretRepositoriesState struct {
	// Name of the existing secret
	SecretName *string `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
}

type ActionsOrganizationSecretRepositoriesState struct {
	// Name of the existing secret
	SecretName pulumi.StringPtrInput
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayInput
}

func (ActionsOrganizationSecretRepositoriesState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationSecretRepositoriesState)(nil)).Elem()
}

type actionsOrganizationSecretRepositoriesArgs struct {
	// Name of the existing secret
	SecretName string `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
}

// The set of arguments for constructing a ActionsOrganizationSecretRepositories resource.
type ActionsOrganizationSecretRepositoriesArgs struct {
	// Name of the existing secret
	SecretName pulumi.StringInput
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayInput
}

func (ActionsOrganizationSecretRepositoriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationSecretRepositoriesArgs)(nil)).Elem()
}

type ActionsOrganizationSecretRepositoriesInput interface {
	pulumi.Input

	ToActionsOrganizationSecretRepositoriesOutput() ActionsOrganizationSecretRepositoriesOutput
	ToActionsOrganizationSecretRepositoriesOutputWithContext(ctx context.Context) ActionsOrganizationSecretRepositoriesOutput
}

func (*ActionsOrganizationSecretRepositories) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationSecretRepositories)(nil)).Elem()
}

func (i *ActionsOrganizationSecretRepositories) ToActionsOrganizationSecretRepositoriesOutput() ActionsOrganizationSecretRepositoriesOutput {
	return i.ToActionsOrganizationSecretRepositoriesOutputWithContext(context.Background())
}

func (i *ActionsOrganizationSecretRepositories) ToActionsOrganizationSecretRepositoriesOutputWithContext(ctx context.Context) ActionsOrganizationSecretRepositoriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationSecretRepositoriesOutput)
}

// ActionsOrganizationSecretRepositoriesArrayInput is an input type that accepts ActionsOrganizationSecretRepositoriesArray and ActionsOrganizationSecretRepositoriesArrayOutput values.
// You can construct a concrete instance of `ActionsOrganizationSecretRepositoriesArrayInput` via:
//
//	ActionsOrganizationSecretRepositoriesArray{ ActionsOrganizationSecretRepositoriesArgs{...} }
type ActionsOrganizationSecretRepositoriesArrayInput interface {
	pulumi.Input

	ToActionsOrganizationSecretRepositoriesArrayOutput() ActionsOrganizationSecretRepositoriesArrayOutput
	ToActionsOrganizationSecretRepositoriesArrayOutputWithContext(context.Context) ActionsOrganizationSecretRepositoriesArrayOutput
}

type ActionsOrganizationSecretRepositoriesArray []ActionsOrganizationSecretRepositoriesInput

func (ActionsOrganizationSecretRepositoriesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationSecretRepositories)(nil)).Elem()
}

func (i ActionsOrganizationSecretRepositoriesArray) ToActionsOrganizationSecretRepositoriesArrayOutput() ActionsOrganizationSecretRepositoriesArrayOutput {
	return i.ToActionsOrganizationSecretRepositoriesArrayOutputWithContext(context.Background())
}

func (i ActionsOrganizationSecretRepositoriesArray) ToActionsOrganizationSecretRepositoriesArrayOutputWithContext(ctx context.Context) ActionsOrganizationSecretRepositoriesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationSecretRepositoriesArrayOutput)
}

// ActionsOrganizationSecretRepositoriesMapInput is an input type that accepts ActionsOrganizationSecretRepositoriesMap and ActionsOrganizationSecretRepositoriesMapOutput values.
// You can construct a concrete instance of `ActionsOrganizationSecretRepositoriesMapInput` via:
//
//	ActionsOrganizationSecretRepositoriesMap{ "key": ActionsOrganizationSecretRepositoriesArgs{...} }
type ActionsOrganizationSecretRepositoriesMapInput interface {
	pulumi.Input

	ToActionsOrganizationSecretRepositoriesMapOutput() ActionsOrganizationSecretRepositoriesMapOutput
	ToActionsOrganizationSecretRepositoriesMapOutputWithContext(context.Context) ActionsOrganizationSecretRepositoriesMapOutput
}

type ActionsOrganizationSecretRepositoriesMap map[string]ActionsOrganizationSecretRepositoriesInput

func (ActionsOrganizationSecretRepositoriesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationSecretRepositories)(nil)).Elem()
}

func (i ActionsOrganizationSecretRepositoriesMap) ToActionsOrganizationSecretRepositoriesMapOutput() ActionsOrganizationSecretRepositoriesMapOutput {
	return i.ToActionsOrganizationSecretRepositoriesMapOutputWithContext(context.Background())
}

func (i ActionsOrganizationSecretRepositoriesMap) ToActionsOrganizationSecretRepositoriesMapOutputWithContext(ctx context.Context) ActionsOrganizationSecretRepositoriesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationSecretRepositoriesMapOutput)
}

type ActionsOrganizationSecretRepositoriesOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationSecretRepositoriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationSecretRepositories)(nil)).Elem()
}

func (o ActionsOrganizationSecretRepositoriesOutput) ToActionsOrganizationSecretRepositoriesOutput() ActionsOrganizationSecretRepositoriesOutput {
	return o
}

func (o ActionsOrganizationSecretRepositoriesOutput) ToActionsOrganizationSecretRepositoriesOutputWithContext(ctx context.Context) ActionsOrganizationSecretRepositoriesOutput {
	return o
}

// Name of the existing secret
func (o ActionsOrganizationSecretRepositoriesOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecretRepositories) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// An array of repository ids that can access the organization secret.
func (o ActionsOrganizationSecretRepositoriesOutput) SelectedRepositoryIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecretRepositories) pulumi.IntArrayOutput { return v.SelectedRepositoryIds }).(pulumi.IntArrayOutput)
}

type ActionsOrganizationSecretRepositoriesArrayOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationSecretRepositoriesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationSecretRepositories)(nil)).Elem()
}

func (o ActionsOrganizationSecretRepositoriesArrayOutput) ToActionsOrganizationSecretRepositoriesArrayOutput() ActionsOrganizationSecretRepositoriesArrayOutput {
	return o
}

func (o ActionsOrganizationSecretRepositoriesArrayOutput) ToActionsOrganizationSecretRepositoriesArrayOutputWithContext(ctx context.Context) ActionsOrganizationSecretRepositoriesArrayOutput {
	return o
}

func (o ActionsOrganizationSecretRepositoriesArrayOutput) Index(i pulumi.IntInput) ActionsOrganizationSecretRepositoriesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionsOrganizationSecretRepositories {
		return vs[0].([]*ActionsOrganizationSecretRepositories)[vs[1].(int)]
	}).(ActionsOrganizationSecretRepositoriesOutput)
}

type ActionsOrganizationSecretRepositoriesMapOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationSecretRepositoriesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationSecretRepositories)(nil)).Elem()
}

func (o ActionsOrganizationSecretRepositoriesMapOutput) ToActionsOrganizationSecretRepositoriesMapOutput() ActionsOrganizationSecretRepositoriesMapOutput {
	return o
}

func (o ActionsOrganizationSecretRepositoriesMapOutput) ToActionsOrganizationSecretRepositoriesMapOutputWithContext(ctx context.Context) ActionsOrganizationSecretRepositoriesMapOutput {
	return o
}

func (o ActionsOrganizationSecretRepositoriesMapOutput) MapIndex(k pulumi.StringInput) ActionsOrganizationSecretRepositoriesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionsOrganizationSecretRepositories {
		return vs[0].(map[string]*ActionsOrganizationSecretRepositories)[vs[1].(string)]
	}).(ActionsOrganizationSecretRepositoriesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationSecretRepositoriesInput)(nil)).Elem(), &ActionsOrganizationSecretRepositories{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationSecretRepositoriesArrayInput)(nil)).Elem(), ActionsOrganizationSecretRepositoriesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationSecretRepositoriesMapInput)(nil)).Elem(), ActionsOrganizationSecretRepositoriesMap{})
	pulumi.RegisterOutputType(ActionsOrganizationSecretRepositoriesOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationSecretRepositoriesArrayOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationSecretRepositoriesMapOutput{})
}
