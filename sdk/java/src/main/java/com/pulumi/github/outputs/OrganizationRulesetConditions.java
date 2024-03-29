// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.outputs.OrganizationRulesetConditionsRefName;
import com.pulumi.github.outputs.OrganizationRulesetConditionsRepositoryName;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OrganizationRulesetConditions {
    /**
     * @return (Block List, Min: 1, Max: 1) (see below for nested schema)
     * 
     */
    private OrganizationRulesetConditionsRefName refName;
    /**
     * @return The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
     * 
     */
    private @Nullable List<Integer> repositoryIds;
    /**
     * @return Conflicts with `repository_id`. (see below for nested schema)
     * 
     * One of `repository_id` and `repository_name` must be set for the rule to target any repositories.
     * 
     */
    private @Nullable OrganizationRulesetConditionsRepositoryName repositoryName;

    private OrganizationRulesetConditions() {}
    /**
     * @return (Block List, Min: 1, Max: 1) (see below for nested schema)
     * 
     */
    public OrganizationRulesetConditionsRefName refName() {
        return this.refName;
    }
    /**
     * @return The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
     * 
     */
    public List<Integer> repositoryIds() {
        return this.repositoryIds == null ? List.of() : this.repositoryIds;
    }
    /**
     * @return Conflicts with `repository_id`. (see below for nested schema)
     * 
     * One of `repository_id` and `repository_name` must be set for the rule to target any repositories.
     * 
     */
    public Optional<OrganizationRulesetConditionsRepositoryName> repositoryName() {
        return Optional.ofNullable(this.repositoryName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrganizationRulesetConditions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private OrganizationRulesetConditionsRefName refName;
        private @Nullable List<Integer> repositoryIds;
        private @Nullable OrganizationRulesetConditionsRepositoryName repositoryName;
        public Builder() {}
        public Builder(OrganizationRulesetConditions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.refName = defaults.refName;
    	      this.repositoryIds = defaults.repositoryIds;
    	      this.repositoryName = defaults.repositoryName;
        }

        @CustomType.Setter
        public Builder refName(OrganizationRulesetConditionsRefName refName) {
            if (refName == null) {
              throw new MissingRequiredPropertyException("OrganizationRulesetConditions", "refName");
            }
            this.refName = refName;
            return this;
        }
        @CustomType.Setter
        public Builder repositoryIds(@Nullable List<Integer> repositoryIds) {

            this.repositoryIds = repositoryIds;
            return this;
        }
        public Builder repositoryIds(Integer... repositoryIds) {
            return repositoryIds(List.of(repositoryIds));
        }
        @CustomType.Setter
        public Builder repositoryName(@Nullable OrganizationRulesetConditionsRepositoryName repositoryName) {

            this.repositoryName = repositoryName;
            return this;
        }
        public OrganizationRulesetConditions build() {
            final var _resultValue = new OrganizationRulesetConditions();
            _resultValue.refName = refName;
            _resultValue.repositoryIds = repositoryIds;
            _resultValue.repositoryName = repositoryName;
            return _resultValue;
        }
    }
}
