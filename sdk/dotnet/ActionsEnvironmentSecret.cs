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
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleSecretActionsEnvironmentSecret = new Github.ActionsEnvironmentSecret("exampleSecretActionsEnvironmentSecret", new Github.ActionsEnvironmentSecretArgs
    ///         {
    ///             Environment = "example_environment",
    ///             SecretName = "example_secret_name",
    ///             PlaintextValue = @var.Some_secret_string,
    ///         });
    ///         var exampleSecretIndex_actionsEnvironmentSecretActionsEnvironmentSecret = new Github.ActionsEnvironmentSecret("exampleSecretIndex/actionsEnvironmentSecretActionsEnvironmentSecret", new Github.ActionsEnvironmentSecretArgs
    ///         {
    ///             Environment = "example_environment",
    ///             SecretName = "example_secret_name",
    ///             EncryptedValue = @var.Some_encrypted_secret_string,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var repo = Output.Create(Github.GetRepository.InvokeAsync(new Github.GetRepositoryArgs
    ///         {
    ///             FullName = "my-org/repo",
    ///         }));
    ///         var repoEnvironment = new Github.RepositoryEnvironment("repoEnvironment", new Github.RepositoryEnvironmentArgs
    ///         {
    ///             Repository = repo.Apply(repo =&gt; repo.Name),
    ///             Environment = "example_environment",
    ///         });
    ///         var testSecret = new Github.ActionsEnvironmentSecret("testSecret", new Github.ActionsEnvironmentSecretArgs
    ///         {
    ///             Repository = repo.Apply(repo =&gt; repo.Name),
    ///             Environment = repoEnvironment.Environment,
    ///             SecretName = "test_secret_name",
    ///             PlaintextValue = "%s",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an ID made up of the secret name
    /// 
    /// ```sh
    ///  $ pulumi import github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret test_secret test_secret_name
    /// ```
    /// 
    ///  NOTEthe implementation is limited in that it won't fetch the value of the `plaintext_value` field when importing. You may need to ignore changes for the `plaintext_value` as a workaround.
    /// </summary>
    [GithubResourceType("github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret")]
    public partial class ActionsEnvironmentSecret : Pulumi.CustomResource
    {
        /// <summary>
        /// Date of actions_environment_secret creation.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Encrypted value of the secret using the Github public key in Base64 format.
        /// </summary>
        [Output("encryptedValue")]
        public Output<string?> EncryptedValue { get; private set; } = null!;

        /// <summary>
        /// Name of the environment.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// Plaintext value of the secret to be encrypted.
        /// </summary>
        [Output("plaintextValue")]
        public Output<string?> PlaintextValue { get; private set; } = null!;

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// Name of the secret.
        /// </summary>
        [Output("secretName")]
        public Output<string> SecretName { get; private set; } = null!;

        /// <summary>
        /// Date of actions_environment_secret update.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ActionsEnvironmentSecret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionsEnvironmentSecret(string name, ActionsEnvironmentSecretArgs args, CustomResourceOptions? options = null)
            : base("github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret", name, args ?? new ActionsEnvironmentSecretArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionsEnvironmentSecret(string name, Input<string> id, ActionsEnvironmentSecretState? state = null, CustomResourceOptions? options = null)
            : base("github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActionsEnvironmentSecret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionsEnvironmentSecret Get(string name, Input<string> id, ActionsEnvironmentSecretState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionsEnvironmentSecret(name, id, state, options);
        }
    }

    public sealed class ActionsEnvironmentSecretArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encrypted value of the secret using the Github public key in Base64 format.
        /// </summary>
        [Input("encryptedValue")]
        public Input<string>? EncryptedValue { get; set; }

        /// <summary>
        /// Name of the environment.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// Plaintext value of the secret to be encrypted.
        /// </summary>
        [Input("plaintextValue")]
        public Input<string>? PlaintextValue { get; set; }

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// Name of the secret.
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        public ActionsEnvironmentSecretArgs()
        {
        }
    }

    public sealed class ActionsEnvironmentSecretState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date of actions_environment_secret creation.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Encrypted value of the secret using the Github public key in Base64 format.
        /// </summary>
        [Input("encryptedValue")]
        public Input<string>? EncryptedValue { get; set; }

        /// <summary>
        /// Name of the environment.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// Plaintext value of the secret to be encrypted.
        /// </summary>
        [Input("plaintextValue")]
        public Input<string>? PlaintextValue { get; set; }

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// Name of the secret.
        /// </summary>
        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        /// <summary>
        /// Date of actions_environment_secret update.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ActionsEnvironmentSecretState()
        {
        }
    }
}
