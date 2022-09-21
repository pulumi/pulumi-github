// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.TeamArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.TeamState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="github:index/team:Team")
public class Team extends com.pulumi.resources.CustomResource {
    @Export(name="createDefaultMaintainer", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> createDefaultMaintainer;

    public Output<Optional<Boolean>> createDefaultMaintainer() {
        return Codegen.optional(this.createDefaultMaintainer);
    }
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="ldapDn", type=String.class, parameters={})
    private Output</* @Nullable */ String> ldapDn;

    public Output<Optional<String>> ldapDn() {
        return Codegen.optional(this.ldapDn);
    }
    @Export(name="membersCount", type=Integer.class, parameters={})
    private Output<Integer> membersCount;

    public Output<Integer> membersCount() {
        return this.membersCount;
    }
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="nodeId", type=String.class, parameters={})
    private Output<String> nodeId;

    public Output<String> nodeId() {
        return this.nodeId;
    }
    @Export(name="parentTeamId", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> parentTeamId;

    public Output<Optional<Integer>> parentTeamId() {
        return Codegen.optional(this.parentTeamId);
    }
    @Export(name="privacy", type=String.class, parameters={})
    private Output</* @Nullable */ String> privacy;

    public Output<Optional<String>> privacy() {
        return Codegen.optional(this.privacy);
    }
    @Export(name="slug", type=String.class, parameters={})
    private Output<String> slug;

    public Output<String> slug() {
        return this.slug;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Team(String name) {
        this(name, TeamArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Team(String name, @Nullable TeamArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Team(String name, @Nullable TeamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/team:Team", name, args == null ? TeamArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Team(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/team:Team", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Team get(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Team(name, id, state, options);
    }
}
