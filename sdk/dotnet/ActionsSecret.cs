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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var examplePublicKey = Github.GetActionsPublicKey.Invoke(new()
    ///     {
    ///         Repository = "example_repository",
    ///     });
    /// 
    ///     var exampleSecretActionsSecret = new Github.ActionsSecret("exampleSecretActionsSecret", new()
    ///     {
    ///         Repository = "example_repository",
    ///         SecretName = "example_secret_name",
    ///         PlaintextValue = @var.Some_secret_string,
    ///     });
    /// 
    ///     var exampleSecretIndex_actionsSecretActionsSecret = new Github.ActionsSecret("exampleSecretIndex/actionsSecretActionsSecret", new()
    ///     {
    ///         Repository = "example_repository",
    ///         SecretName = "example_secret_name",
    ///         EncryptedValue = @var.Some_encrypted_secret_string,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an ID made up of the `repository` and `secret_name`:
    /// 
    /// ```sh
    /// $ pulumi import github:index/actionsSecret:ActionsSecret example_secret repository/secret_name
    /// ```
    /// NOTE: the implementation is limited in that it won't fetch the value of the
    /// `plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
    /// </summary>
    [GithubResourceType("github:index/actionsSecret:ActionsSecret")]
    public partial class ActionsSecret : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date of actions_secret creation.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Encrypted value of the secret using the GitHub public key in Base64 format.
        /// </summary>
        [Output("encryptedValue")]
        public Output<string?> EncryptedValue { get; private set; } = null!;

        /// <summary>
        /// Plaintext value of the secret to be encrypted
        /// </summary>
        [Output("plaintextValue")]
        public Output<string?> PlaintextValue { get; private set; } = null!;

        /// <summary>
        /// Name of the repository
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// Name of the secret
        /// </summary>
        [Output("secretName")]
        public Output<string> SecretName { get; private set; } = null!;

        /// <summary>
        /// Date of actions_secret update.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ActionsSecret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionsSecret(string name, ActionsSecretArgs args, CustomResourceOptions? options = null)
            : base("github:index/actionsSecret:ActionsSecret", name, args ?? new ActionsSecretArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionsSecret(string name, Input<string> id, ActionsSecretState? state = null, CustomResourceOptions? options = null)
            : base("github:index/actionsSecret:ActionsSecret", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "encryptedValue",
                    "plaintextValue",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ActionsSecret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionsSecret Get(string name, Input<string> id, ActionsSecretState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionsSecret(name, id, state, options);
        }
    }

    public sealed class ActionsSecretArgs : global::Pulumi.ResourceArgs
    {
        [Input("encryptedValue")]
        private Input<string>? _encryptedValue;

        /// <summary>
        /// Encrypted value of the secret using the GitHub public key in Base64 format.
        /// </summary>
        public Input<string>? EncryptedValue
        {
            get => _encryptedValue;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _encryptedValue = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("plaintextValue")]
        private Input<string>? _plaintextValue;

        /// <summary>
        /// Plaintext value of the secret to be encrypted
        /// </summary>
        public Input<string>? PlaintextValue
        {
            get => _plaintextValue;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _plaintextValue = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of the repository
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// Name of the secret
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        public ActionsSecretArgs()
        {
        }
        public static new ActionsSecretArgs Empty => new ActionsSecretArgs();
    }

    public sealed class ActionsSecretState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date of actions_secret creation.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("encryptedValue")]
        private Input<string>? _encryptedValue;

        /// <summary>
        /// Encrypted value of the secret using the GitHub public key in Base64 format.
        /// </summary>
        public Input<string>? EncryptedValue
        {
            get => _encryptedValue;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _encryptedValue = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("plaintextValue")]
        private Input<string>? _plaintextValue;

        /// <summary>
        /// Plaintext value of the secret to be encrypted
        /// </summary>
        public Input<string>? PlaintextValue
        {
            get => _plaintextValue;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _plaintextValue = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of the repository
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// Name of the secret
        /// </summary>
        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        /// <summary>
        /// Date of actions_secret update.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ActionsSecretState()
        {
        }
        public static new ActionsSecretState Empty => new ActionsSecretState();
    }
}
