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
    /// Provides a GitHub repository deploy key resource.
    /// 
    /// A deploy key is an SSH key that is stored on your server and grants
    /// access to a single GitHub repository. This key is attached directly to the repository instead of to a personal user
    /// account.
    /// 
    /// This resource allows you to add/remove repository deploy keys.
    /// 
    /// Further documentation on GitHub repository deploy keys:
    /// - [About deploy keys](https://developer.github.com/guides/managing-deploy-keys/#deploy-keys)
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// using Tls = Pulumi.Tls;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Generate an ssh key using provider "hashicorp/tls"
    ///     var exampleRepositoryDeployKeyPrivateKey = new Tls.PrivateKey("exampleRepositoryDeployKeyPrivateKey", new()
    ///     {
    ///         Algorithm = "ED25519",
    ///     });
    /// 
    ///     // Add the ssh key as a deploy key
    ///     var exampleRepositoryDeployKeyRepositoryDeployKey = new Github.RepositoryDeployKey("exampleRepositoryDeployKeyRepositoryDeployKey", new()
    ///     {
    ///         Title = "Repository test key",
    ///         Repository = "test-repo",
    ///         Key = exampleRepositoryDeployKeyPrivateKey.PublicKeyOpenssh,
    ///         ReadOnly = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Repository deploy keys can be imported using a colon-separated pair of repository name
    /// and GitHub's key id. The latter can be obtained by GitHub's SDKs and API.
    /// 
    /// ```sh
    /// $ pulumi import github:index/repositoryDeployKey:RepositoryDeployKey foo test-repo:23824728
    /// ```
    /// </summary>
    [GithubResourceType("github:index/repositoryDeployKey:RepositoryDeployKey")]
    public partial class RepositoryDeployKey : global::Pulumi.CustomResource
    {
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A SSH key.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// A boolean qualifying the key to be either read only or read/write.
        /// </summary>
        [Output("readOnly")]
        public Output<bool?> ReadOnly { get; private set; } = null!;

        /// <summary>
        /// Name of the GitHub repository.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// A title.
        /// 
        /// Changing any of the fields forces re-creating the resource.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryDeployKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryDeployKey(string name, RepositoryDeployKeyArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryDeployKey:RepositoryDeployKey", name, args ?? new RepositoryDeployKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryDeployKey(string name, Input<string> id, RepositoryDeployKeyState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryDeployKey:RepositoryDeployKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryDeployKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryDeployKey Get(string name, Input<string> id, RepositoryDeployKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryDeployKey(name, id, state, options);
        }
    }

    public sealed class RepositoryDeployKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A SSH key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// A boolean qualifying the key to be either read only or read/write.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Name of the GitHub repository.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// A title.
        /// 
        /// Changing any of the fields forces re-creating the resource.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public RepositoryDeployKeyArgs()
        {
        }
        public static new RepositoryDeployKeyArgs Empty => new RepositoryDeployKeyArgs();
    }

    public sealed class RepositoryDeployKeyState : global::Pulumi.ResourceArgs
    {
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// A SSH key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// A boolean qualifying the key to be either read only or read/write.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Name of the GitHub repository.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// A title.
        /// 
        /// Changing any of the fields forces re-creating the resource.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public RepositoryDeployKeyState()
        {
        }
        public static new RepositoryDeployKeyState Empty => new RepositoryDeployKeyState();
    }
}
