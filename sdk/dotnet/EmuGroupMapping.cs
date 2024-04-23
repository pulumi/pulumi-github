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
    /// This resource manages mappings between external groups for enterprise managed users and GitHub teams. It wraps the API detailed [here](https://docs.github.com/en/rest/reference/teams#external-groups). Note that this is a distinct resource from `github.TeamSyncGroupMapping`. `github.EmuGroupMapping` is special to the Enterprise Managed User (EMU) external group feature, whereas `github.TeamSyncGroupMapping` is specific to Identity Provider Groups.
    /// 
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
    ///     var exampleEmuGroupMapping = new Github.EmuGroupMapping("example_emu_group_mapping", new()
    ///     {
    ///         TeamSlug = "emu-test-team",
    ///         GroupId = 28836,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitHub EMU External Group Mappings can be imported using the external `group_id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/emuGroupMapping:EmuGroupMapping example_emu_group_mapping 28836
    /// ```
    /// </summary>
    [GithubResourceType("github:index/emuGroupMapping:EmuGroupMapping")]
    public partial class EmuGroupMapping : global::Pulumi.CustomResource
    {
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Integer corresponding to the external group ID to be linked
        /// </summary>
        [Output("groupId")]
        public Output<int> GroupId { get; private set; } = null!;

        /// <summary>
        /// Slug of the GitHub team
        /// </summary>
        [Output("teamSlug")]
        public Output<string> TeamSlug { get; private set; } = null!;


        /// <summary>
        /// Create a EmuGroupMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EmuGroupMapping(string name, EmuGroupMappingArgs args, CustomResourceOptions? options = null)
            : base("github:index/emuGroupMapping:EmuGroupMapping", name, args ?? new EmuGroupMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EmuGroupMapping(string name, Input<string> id, EmuGroupMappingState? state = null, CustomResourceOptions? options = null)
            : base("github:index/emuGroupMapping:EmuGroupMapping", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EmuGroupMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EmuGroupMapping Get(string name, Input<string> id, EmuGroupMappingState? state = null, CustomResourceOptions? options = null)
        {
            return new EmuGroupMapping(name, id, state, options);
        }
    }

    public sealed class EmuGroupMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer corresponding to the external group ID to be linked
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        /// <summary>
        /// Slug of the GitHub team
        /// </summary>
        [Input("teamSlug", required: true)]
        public Input<string> TeamSlug { get; set; } = null!;

        public EmuGroupMappingArgs()
        {
        }
        public static new EmuGroupMappingArgs Empty => new EmuGroupMappingArgs();
    }

    public sealed class EmuGroupMappingState : global::Pulumi.ResourceArgs
    {
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Integer corresponding to the external group ID to be linked
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// Slug of the GitHub team
        /// </summary>
        [Input("teamSlug")]
        public Input<string>? TeamSlug { get; set; }

        public EmuGroupMappingState()
        {
        }
        public static new EmuGroupMappingState Empty => new EmuGroupMappingState();
    }
}
