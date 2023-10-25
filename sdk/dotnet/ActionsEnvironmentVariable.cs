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
    /// This resource allows you to create and manage GitHub Actions variables within your GitHub repository environments.
    /// You must have write access to a repository to use this resource.
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an ID made up of the repository name, environment name, and variable name:
    /// 
    /// ```sh
    ///  $ pulumi import github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable test_variable myrepo:myenv:myvariable
    /// ```
    /// </summary>
    [GithubResourceType("github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable")]
    public partial class ActionsEnvironmentVariable : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date of actions_environment_secret creation.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the environment.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// Date of actions_environment_secret update.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Value of the variable
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// Name of the variable.
        /// </summary>
        [Output("variableName")]
        public Output<string> VariableName { get; private set; } = null!;


        /// <summary>
        /// Create a ActionsEnvironmentVariable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionsEnvironmentVariable(string name, ActionsEnvironmentVariableArgs args, CustomResourceOptions? options = null)
            : base("github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable", name, args ?? new ActionsEnvironmentVariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionsEnvironmentVariable(string name, Input<string> id, ActionsEnvironmentVariableState? state = null, CustomResourceOptions? options = null)
            : base("github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActionsEnvironmentVariable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionsEnvironmentVariable Get(string name, Input<string> id, ActionsEnvironmentVariableState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionsEnvironmentVariable(name, id, state, options);
        }
    }

    public sealed class ActionsEnvironmentVariableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the environment.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// Value of the variable
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// Name of the variable.
        /// </summary>
        [Input("variableName", required: true)]
        public Input<string> VariableName { get; set; } = null!;

        public ActionsEnvironmentVariableArgs()
        {
        }
        public static new ActionsEnvironmentVariableArgs Empty => new ActionsEnvironmentVariableArgs();
    }

    public sealed class ActionsEnvironmentVariableState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date of actions_environment_secret creation.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Name of the environment.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// Date of actions_environment_secret update.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Value of the variable
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Name of the variable.
        /// </summary>
        [Input("variableName")]
        public Input<string>? VariableName { get; set; }

        public ActionsEnvironmentVariableState()
        {
        }
        public static new ActionsEnvironmentVariableState Empty => new ActionsEnvironmentVariableState();
    }
}
