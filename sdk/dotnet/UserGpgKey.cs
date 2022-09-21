// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/userGpgKey:UserGpgKey")]
    public partial class UserGpgKey : global::Pulumi.CustomResource
    {
        [Output("armoredPublicKey")]
        public Output<string> ArmoredPublicKey { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;


        /// <summary>
        /// Create a UserGpgKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserGpgKey(string name, UserGpgKeyArgs args, CustomResourceOptions? options = null)
            : base("github:index/userGpgKey:UserGpgKey", name, args ?? new UserGpgKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserGpgKey(string name, Input<string> id, UserGpgKeyState? state = null, CustomResourceOptions? options = null)
            : base("github:index/userGpgKey:UserGpgKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserGpgKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserGpgKey Get(string name, Input<string> id, UserGpgKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new UserGpgKey(name, id, state, options);
        }
    }

    public sealed class UserGpgKeyArgs : global::Pulumi.ResourceArgs
    {
        [Input("armoredPublicKey", required: true)]
        public Input<string> ArmoredPublicKey { get; set; } = null!;

        public UserGpgKeyArgs()
        {
        }
        public static new UserGpgKeyArgs Empty => new UserGpgKeyArgs();
    }

    public sealed class UserGpgKeyState : global::Pulumi.ResourceArgs
    {
        [Input("armoredPublicKey")]
        public Input<string>? ArmoredPublicKey { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        public UserGpgKeyState()
        {
        }
        public static new UserGpgKeyState Empty => new UserGpgKeyState();
    }
}
