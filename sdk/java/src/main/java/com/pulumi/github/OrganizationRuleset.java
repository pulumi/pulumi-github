// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.OrganizationRulesetArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.OrganizationRulesetState;
import com.pulumi.github.outputs.OrganizationRulesetBypassActor;
import com.pulumi.github.outputs.OrganizationRulesetConditions;
import com.pulumi.github.outputs.OrganizationRulesetRules;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a GitHub organization ruleset.
 * 
 * This resource allows you to create and manage rulesets on the organization level. When applied, a new ruleset will be created. When destroyed, that ruleset will be removed.
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.OrganizationRuleset;
 * import com.pulumi.github.OrganizationRulesetArgs;
 * import com.pulumi.github.inputs.OrganizationRulesetBypassActorArgs;
 * import com.pulumi.github.inputs.OrganizationRulesetConditionsArgs;
 * import com.pulumi.github.inputs.OrganizationRulesetConditionsRefNameArgs;
 * import com.pulumi.github.inputs.OrganizationRulesetRulesArgs;
 * import com.pulumi.github.inputs.OrganizationRulesetRulesBranchNamePatternArgs;
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
 *         var example = new OrganizationRuleset(&#34;example&#34;, OrganizationRulesetArgs.builder()        
 *             .bypassActors(OrganizationRulesetBypassActorArgs.builder()
 *                 .actorId(13473)
 *                 .actorType(&#34;Integration&#34;)
 *                 .bypassMode(&#34;always&#34;)
 *                 .build())
 *             .conditions(OrganizationRulesetConditionsArgs.builder()
 *                 .refName(OrganizationRulesetConditionsRefNameArgs.builder()
 *                     .exclude()
 *                     .include(&#34;~ALL&#34;)
 *                     .build())
 *                 .build())
 *             .enforcement(&#34;active&#34;)
 *             .rules(OrganizationRulesetRulesArgs.builder()
 *                 .branchNamePattern(OrganizationRulesetRulesBranchNamePatternArgs.builder()
 *                     .name(&#34;example&#34;)
 *                     .negate(false)
 *                     .operator(&#34;starts_with&#34;)
 *                     .pattern(&#34;ex&#34;)
 *                     .build())
 *                 .creation(true)
 *                 .deletion(true)
 *                 .requiredLinearHistory(true)
 *                 .requiredSignatures(true)
 *                 .update(true)
 *                 .build())
 *             .target(&#34;branch&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Organization Rulesets can be imported using the GitHub ruleset ID e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/organizationRuleset:OrganizationRuleset example 12345`
 * ```
 * 
 */
@ResourceType(type="github:index/organizationRuleset:OrganizationRuleset")
public class OrganizationRuleset extends com.pulumi.resources.CustomResource {
    /**
     * (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
     * 
     */
    @Export(name="bypassActors", refs={List.class,OrganizationRulesetBypassActor.class}, tree="[0,1]")
    private Output</* @Nullable */ List<OrganizationRulesetBypassActor>> bypassActors;

    /**
     * @return (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
     * 
     */
    public Output<Optional<List<OrganizationRulesetBypassActor>>> bypassActors() {
        return Codegen.optional(this.bypassActors);
    }
    /**
     * (Block List, Max: 1) Parameters for an organization ruleset ref name condition. (see below for nested schema)
     * 
     */
    @Export(name="conditions", refs={OrganizationRulesetConditions.class}, tree="[0]")
    private Output</* @Nullable */ OrganizationRulesetConditions> conditions;

    /**
     * @return (Block List, Max: 1) Parameters for an organization ruleset ref name condition. (see below for nested schema)
     * 
     */
    public Output<Optional<OrganizationRulesetConditions>> conditions() {
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
     * (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
     * 
     */
    @Export(name="rules", refs={OrganizationRulesetRules.class}, tree="[0]")
    private Output<OrganizationRulesetRules> rules;

    /**
     * @return (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
     * 
     */
    public Output<OrganizationRulesetRules> rules() {
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
    public OrganizationRuleset(String name) {
        this(name, OrganizationRulesetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrganizationRuleset(String name, OrganizationRulesetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrganizationRuleset(String name, OrganizationRulesetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/organizationRuleset:OrganizationRuleset", name, args == null ? OrganizationRulesetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OrganizationRuleset(String name, Output<String> id, @Nullable OrganizationRulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/organizationRuleset:OrganizationRuleset", name, state, makeResourceOptions(options, id));
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
    public static OrganizationRuleset get(String name, Output<String> id, @Nullable OrganizationRulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrganizationRuleset(name, id, state, options);
    }
}
