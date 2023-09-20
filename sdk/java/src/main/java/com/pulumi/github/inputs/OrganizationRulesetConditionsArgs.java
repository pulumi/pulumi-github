// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.OrganizationRulesetConditionsRefNameArgs;
import com.pulumi.github.inputs.OrganizationRulesetConditionsRepositoryNameArgs;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrganizationRulesetConditionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationRulesetConditionsArgs Empty = new OrganizationRulesetConditionsArgs();

    /**
     * (Block List, Min: 1, Max: 1) (see below for nested schema)
     * 
     */
    @Import(name="refName", required=true)
    private Output<OrganizationRulesetConditionsRefNameArgs> refName;

    /**
     * @return (Block List, Min: 1, Max: 1) (see below for nested schema)
     * 
     */
    public Output<OrganizationRulesetConditionsRefNameArgs> refName() {
        return this.refName;
    }

    /**
     * The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
     * 
     */
    @Import(name="repositoryId")
    private @Nullable Output<Integer> repositoryId;

    /**
     * @return The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
     * 
     */
    public Optional<Output<Integer>> repositoryId() {
        return Optional.ofNullable(this.repositoryId);
    }

    /**
     * Conflicts with `repository_id`. (see below for nested schema)
     * 
     * One of `repository_id` and `repository_name` must be set for the rule to target any repositories.
     * 
     */
    @Import(name="repositoryName")
    private @Nullable Output<OrganizationRulesetConditionsRepositoryNameArgs> repositoryName;

    /**
     * @return Conflicts with `repository_id`. (see below for nested schema)
     * 
     * One of `repository_id` and `repository_name` must be set for the rule to target any repositories.
     * 
     */
    public Optional<Output<OrganizationRulesetConditionsRepositoryNameArgs>> repositoryName() {
        return Optional.ofNullable(this.repositoryName);
    }

    private OrganizationRulesetConditionsArgs() {}

    private OrganizationRulesetConditionsArgs(OrganizationRulesetConditionsArgs $) {
        this.refName = $.refName;
        this.repositoryId = $.repositoryId;
        this.repositoryName = $.repositoryName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationRulesetConditionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationRulesetConditionsArgs $;

        public Builder() {
            $ = new OrganizationRulesetConditionsArgs();
        }

        public Builder(OrganizationRulesetConditionsArgs defaults) {
            $ = new OrganizationRulesetConditionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param refName (Block List, Min: 1, Max: 1) (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder refName(Output<OrganizationRulesetConditionsRefNameArgs> refName) {
            $.refName = refName;
            return this;
        }

        /**
         * @param refName (Block List, Min: 1, Max: 1) (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder refName(OrganizationRulesetConditionsRefNameArgs refName) {
            return refName(Output.of(refName));
        }

        /**
         * @param repositoryId The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(@Nullable Output<Integer> repositoryId) {
            $.repositoryId = repositoryId;
            return this;
        }

        /**
         * @param repositoryId The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(Integer repositoryId) {
            return repositoryId(Output.of(repositoryId));
        }

        /**
         * @param repositoryName Conflicts with `repository_id`. (see below for nested schema)
         * 
         * One of `repository_id` and `repository_name` must be set for the rule to target any repositories.
         * 
         * @return builder
         * 
         */
        public Builder repositoryName(@Nullable Output<OrganizationRulesetConditionsRepositoryNameArgs> repositoryName) {
            $.repositoryName = repositoryName;
            return this;
        }

        /**
         * @param repositoryName Conflicts with `repository_id`. (see below for nested schema)
         * 
         * One of `repository_id` and `repository_name` must be set for the rule to target any repositories.
         * 
         * @return builder
         * 
         */
        public Builder repositoryName(OrganizationRulesetConditionsRepositoryNameArgs repositoryName) {
            return repositoryName(Output.of(repositoryName));
        }

        public OrganizationRulesetConditionsArgs build() {
            $.refName = Objects.requireNonNull($.refName, "expected parameter 'refName' to be non-null");
            return $;
        }
    }

}
