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
    /// This resource allows you to create and manage an OpenID Connect subject claim customization template for a GitHub
    /// repository.
    /// 
    /// More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
    /// available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).
    /// 
    /// The following table lists the behaviour of `use_default`:
    /// 
    /// | `use_default` | `include_claim_keys` | Template used                                             |
    /// |---------------|----------------------|-----------------------------------------------------------|
    /// | `true`        | Unset                | GitHub's default                                          |
    /// | `false`       | Set                  | `include_claim_keys`                                      |
    /// | `false`       | Unset                | Organization's default if set, otherwise GitHub's default |
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Github.Repository("example");
    /// 
    ///     var exampleTemplate = new Github.ActionsRepositoryOidcSubjectClaimCustomizationTemplate("exampleTemplate", new()
    ///     {
    ///         Repository = example.Name,
    ///         UseDefault = false,
    ///         IncludeClaimKeys = new[]
    ///         {
    ///             "actor",
    ///             "context",
    ///             "repository_owner",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the repository's name.
    /// 
    /// ```sh
    /// $ pulumi import github:index/actionsRepositoryOidcSubjectClaimCustomizationTemplate:ActionsRepositoryOidcSubjectClaimCustomizationTemplate test example_repository
    /// ```
    /// </summary>
    [GithubResourceType("github:index/actionsRepositoryOidcSubjectClaimCustomizationTemplate:ActionsRepositoryOidcSubjectClaimCustomizationTemplate")]
    public partial class ActionsRepositoryOidcSubjectClaimCustomizationTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of OpenID Connect claims.
        /// </summary>
        [Output("includeClaimKeys")]
        public Output<ImmutableArray<string>> IncludeClaimKeys { get; private set; } = null!;

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// Whether to use the default template or not. If `true`, `include_claim_keys` must not
        /// be set.
        /// </summary>
        [Output("useDefault")]
        public Output<bool> UseDefault { get; private set; } = null!;


        /// <summary>
        /// Create a ActionsRepositoryOidcSubjectClaimCustomizationTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionsRepositoryOidcSubjectClaimCustomizationTemplate(string name, ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs args, CustomResourceOptions? options = null)
            : base("github:index/actionsRepositoryOidcSubjectClaimCustomizationTemplate:ActionsRepositoryOidcSubjectClaimCustomizationTemplate", name, args ?? new ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionsRepositoryOidcSubjectClaimCustomizationTemplate(string name, Input<string> id, ActionsRepositoryOidcSubjectClaimCustomizationTemplateState? state = null, CustomResourceOptions? options = null)
            : base("github:index/actionsRepositoryOidcSubjectClaimCustomizationTemplate:ActionsRepositoryOidcSubjectClaimCustomizationTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActionsRepositoryOidcSubjectClaimCustomizationTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionsRepositoryOidcSubjectClaimCustomizationTemplate Get(string name, Input<string> id, ActionsRepositoryOidcSubjectClaimCustomizationTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionsRepositoryOidcSubjectClaimCustomizationTemplate(name, id, state, options);
        }
    }

    public sealed class ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// Whether to use the default template or not. If `true`, `include_claim_keys` must not
        /// be set.
        /// </summary>
        [Input("useDefault", required: true)]
        public Input<bool> UseDefault { get; set; } = null!;

        public ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs()
        {
        }
        public static new ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs Empty => new ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs();
    }

    public sealed class ActionsRepositoryOidcSubjectClaimCustomizationTemplateState : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// Whether to use the default template or not. If `true`, `include_claim_keys` must not
        /// be set.
        /// </summary>
        [Input("useDefault")]
        public Input<bool>? UseDefault { get; set; }

        public ActionsRepositoryOidcSubjectClaimCustomizationTemplateState()
        {
        }
        public static new ActionsRepositoryOidcSubjectClaimCustomizationTemplateState Empty => new ActionsRepositoryOidcSubjectClaimCustomizationTemplateState();
    }
}
