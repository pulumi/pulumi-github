// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.TeamSyncGroupMappingArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.TeamSyncGroupMappingState;
import com.pulumi.github.outputs.TeamSyncGroupMappingGroup;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage Identity Provider (IdP) group connections within your GitHub teams.
 * You must have team synchronization enabled for organizations owned by enterprise accounts.
 * 
 * To learn more about team synchronization between IdPs and GitHub, please refer to:
 * &lt;https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github&gt;
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitHub Team Sync Group Mappings can be imported using the GitHub team `slug` e.g.
 * 
 * ```sh
 * $ pulumi import github:index/teamSyncGroupMapping:TeamSyncGroupMapping example some_team
 * ```
 * 
 */
@ResourceType(type="github:index/teamSyncGroupMapping:TeamSyncGroupMapping")
public class TeamSyncGroupMapping extends com.pulumi.resources.CustomResource {
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
     * 
     */
    @Export(name="groups", refs={List.class,TeamSyncGroupMappingGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TeamSyncGroupMappingGroup>> groups;

    /**
     * @return An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
     * 
     */
    public Output<Optional<List<TeamSyncGroupMappingGroup>>> groups() {
        return Codegen.optional(this.groups);
    }
    /**
     * Slug of the team
     * 
     */
    @Export(name="teamSlug", refs={String.class}, tree="[0]")
    private Output<String> teamSlug;

    /**
     * @return Slug of the team
     * 
     */
    public Output<String> teamSlug() {
        return this.teamSlug;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TeamSyncGroupMapping(java.lang.String name) {
        this(name, TeamSyncGroupMappingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamSyncGroupMapping(java.lang.String name, TeamSyncGroupMappingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamSyncGroupMapping(java.lang.String name, TeamSyncGroupMappingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamSyncGroupMapping:TeamSyncGroupMapping", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TeamSyncGroupMapping(java.lang.String name, Output<java.lang.String> id, @Nullable TeamSyncGroupMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamSyncGroupMapping:TeamSyncGroupMapping", name, state, makeResourceOptions(options, id), false);
    }

    private static TeamSyncGroupMappingArgs makeArgs(TeamSyncGroupMappingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TeamSyncGroupMappingArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static TeamSyncGroupMapping get(java.lang.String name, Output<java.lang.String> id, @Nullable TeamSyncGroupMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamSyncGroupMapping(name, id, state, options);
    }
}
