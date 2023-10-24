// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    /// <summary>
    /// This resource allows you to create and manage an OpenID Connect subject claim customization template within a GitHub
    /// organization.
    /// 
    /// More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
    /// available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the organization's name.
    /// 
    /// ```sh
    ///  $ pulumi import github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate test example_organization
    /// ```
    /// </summary>
    [GithubResourceType("github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate")]
    public partial class ActionsOrganizationOidcSubjectClaimCustomizationTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of OpenID Connect claims.
        /// </summary>
        [Output("includeClaimKeys")]
        public Output<ImmutableArray<string>> IncludeClaimKeys { get; private set; } = null!;


        /// <summary>
        /// Create a ActionsOrganizationOidcSubjectClaimCustomizationTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionsOrganizationOidcSubjectClaimCustomizationTemplate(string name, ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs args, CustomResourceOptions? options = null)
            : base("github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate", name, args ?? new ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionsOrganizationOidcSubjectClaimCustomizationTemplate(string name, Input<string> id, ActionsOrganizationOidcSubjectClaimCustomizationTemplateState? state = null, CustomResourceOptions? options = null)
            : base("github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ActionsOrganizationOidcSubjectClaimCustomizationTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionsOrganizationOidcSubjectClaimCustomizationTemplate Get(string name, Input<string> id, ActionsOrganizationOidcSubjectClaimCustomizationTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionsOrganizationOidcSubjectClaimCustomizationTemplate(name, id, state, options);
        }
    }

    public sealed class ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("includeClaimKeys", required: true)]
        private InputList<string>? _includeClaimKeys;

        /// <summary>
        /// A list of OpenID Connect claims.
        /// </summary>
        public InputList<string> IncludeClaimKeys
        {
            get => _includeClaimKeys ?? (_includeClaimKeys = new InputList<string>());
            set => _includeClaimKeys = value;
        }

        public ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs()
        {
        }
        public static new ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs Empty => new ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs();
    }

    public sealed class ActionsOrganizationOidcSubjectClaimCustomizationTemplateState : global::Pulumi.ResourceArgs
    {
        [Input("includeClaimKeys")]
        private InputList<string>? _includeClaimKeys;

        /// <summary>
        /// A list of OpenID Connect claims.
        /// </summary>
        public InputList<string> IncludeClaimKeys
        {
            get => _includeClaimKeys ?? (_includeClaimKeys = new InputList<string>());
            set => _includeClaimKeys = value;
        }

        public ActionsOrganizationOidcSubjectClaimCustomizationTemplateState()
        {
        }
        public static new ActionsOrganizationOidcSubjectClaimCustomizationTemplateState Empty => new ActionsOrganizationOidcSubjectClaimCustomizationTemplateState();
    }
}
