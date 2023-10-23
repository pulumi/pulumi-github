// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryRulesetArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryRulesetState;
import com.pulumi.github.outputs.RepositoryRulesetBypassActor;
import com.pulumi.github.outputs.RepositoryRulesetConditions;
import com.pulumi.github.outputs.RepositoryRulesetRules;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a GitHub repository ruleset.
 * 
 * This resource allows you to create and manage rulesets on the repository level. When applied, a new ruleset will be created. When destroyed, that ruleset will be removed.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.RepositoryRuleset;
 * import com.pulumi.github.RepositoryRulesetArgs;
 * import com.pulumi.github.inputs.RepositoryRulesetConditionsArgs;
 * import com.pulumi.github.inputs.RepositoryRulesetConditionsRefNameArgs;
 * import com.pulumi.github.inputs.RepositoryRulesetBypassActorArgs;
 * import com.pulumi.github.inputs.RepositoryRulesetRulesArgs;
 * import com.pulumi.github.inputs.RepositoryRulesetRulesRequiredDeploymentsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleRepository = new Repository(&#34;exampleRepository&#34;, RepositoryArgs.builder()        
 *             .description(&#34;Example repository&#34;)
 *             .build());
 * 
 *         var exampleRepositoryRuleset = new RepositoryRuleset(&#34;exampleRepositoryRuleset&#34;, RepositoryRulesetArgs.builder()        
 *             .repository(exampleRepository.name())
 *             .target(&#34;branch&#34;)
 *             .enforcement(&#34;active&#34;)
 *             .conditions(RepositoryRulesetConditionsArgs.builder()
 *                 .refName(RepositoryRulesetConditionsRefNameArgs.builder()
 *                     .includes(&#34;~ALL&#34;)
 *                     .excludes()
 *                     .build())
 *                 .build())
 *             .bypassActors(RepositoryRulesetBypassActorArgs.builder()
 *                 .actorId(13473)
 *                 .actorType(&#34;Integration&#34;)
 *                 .bypassMode(&#34;always&#34;)
 *                 .build())
 *             .rules(RepositoryRulesetRulesArgs.builder()
 *                 .creation(true)
 *                 .update(true)
 *                 .deletion(true)
 *                 .requiredLinearHistory(true)
 *                 .requiredSignatures(true)
 *                 .requiredDeployments(RepositoryRulesetRulesRequiredDeploymentsArgs.builder()
 *                     .requiredDeploymentEnvironments(&#34;test&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Repository Rulesets can be imported using the GitHub repository name and ruleset ID e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/repositoryRuleset:RepositoryRuleset example example:12345`
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryRuleset:RepositoryRuleset")
public class RepositoryRuleset extends com.pulumi.resources.CustomResource {
    /**
     * (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
     * 
     */
    @Export(name="bypassActors", refs={List.class,RepositoryRulesetBypassActor.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RepositoryRulesetBypassActor>> bypassActors;

    /**
     * @return (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
     * 
     */
    public Output<Optional<List<RepositoryRulesetBypassActor>>> bypassActors() {
        return Codegen.optional(this.bypassActors);
    }
    /**
     * (Block List, Max: 1) Parameters for a repository ruleset ref name condition. (see below for nested schema)
     * 
     */
    @Export(name="conditions", refs={RepositoryRulesetConditions.class}, tree="[0]")
    private Output</* @Nullable */ RepositoryRulesetConditions> conditions;

    /**
     * @return (Block List, Max: 1) Parameters for a repository ruleset ref name condition. (see below for nested schema)
     * 
     */
    public Output<Optional<RepositoryRulesetConditions>> conditions() {
        return Codegen.optional(this.conditions);
    }
    /**
     * (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
     * 
     */
    @Export(name="enforcement", refs={String.class}, tree="[0]")
    private Output<String> enforcement;

    /**
     * @return (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
     * 
     */
    public Output<String> enforcement() {
        return this.enforcement;
    }
    /**
     * (String)
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (String)
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * (String) The name of the ruleset.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return (String) The name of the ruleset.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * (String) GraphQL global node id for use with v4 API.
     * 
     */
    @Export(name="nodeId", refs={String.class}, tree="[0]")
    private Output<String> nodeId;

    /**
     * @return (String) GraphQL global node id for use with v4 API.
     * 
     */
    public Output<String> nodeId() {
        return this.nodeId;
    }
    /**
     * (String) Name of the repository to apply rulset to.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repository;

    /**
     * @return (String) Name of the repository to apply rulset to.
     * 
     */
    public Output<Optional<String>> repository() {
        return Codegen.optional(this.repository);
    }
    /**
     * (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
     * 
     */
    @Export(name="rules", refs={RepositoryRulesetRules.class}, tree="[0]")
    private Output<RepositoryRulesetRules> rules;

    /**
     * @return (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
     * 
     */
    public Output<RepositoryRulesetRules> rules() {
        return this.rules;
    }
    /**
     * (Number) GitHub ID for the ruleset.
     * 
     */
    @Export(name="rulesetId", refs={Integer.class}, tree="[0]")
    private Output<Integer> rulesetId;

    /**
     * @return (Number) GitHub ID for the ruleset.
     * 
     */
    public Output<Integer> rulesetId() {
        return this.rulesetId;
    }
    /**
     * (String) Possible values are `branch` and `tag`.
     * 
     */
    @Export(name="target", refs={String.class}, tree="[0]")
    private Output<String> target;

    /**
     * @return (String) Possible values are `branch` and `tag`.
     * 
     */
    public Output<String> target() {
        return this.target;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryRuleset(String name) {
        this(name, RepositoryRulesetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryRuleset(String name, RepositoryRulesetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryRuleset(String name, RepositoryRulesetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryRuleset:RepositoryRuleset", name, args == null ? RepositoryRulesetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryRuleset(String name, Output<String> id, @Nullable RepositoryRulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryRuleset:RepositoryRuleset", name, state, makeResourceOptions(options, id));
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
    public static RepositoryRuleset get(String name, Output<String> id, @Nullable RepositoryRulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryRuleset(name, id, state, options);
    }
}