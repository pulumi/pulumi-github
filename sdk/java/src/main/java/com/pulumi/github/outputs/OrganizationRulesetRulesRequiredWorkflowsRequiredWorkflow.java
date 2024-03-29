// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow {
    /**
     * @return (String) The path to the YAML definition file of the workflow.
     * 
     */
    private String path;
    /**
     * @return (String) The optional ref from which to fetch the workflow. Defaults to `master`.
     * 
     */
    private @Nullable String ref;
    /**
     * @return The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
     * 
     */
    private Integer repositoryId;

    private OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow() {}
    /**
     * @return (String) The path to the YAML definition file of the workflow.
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return (String) The optional ref from which to fetch the workflow. Defaults to `master`.
     * 
     */
    public Optional<String> ref() {
        return Optional.ofNullable(this.ref);
    }
    /**
     * @return The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
     * 
     */
    public Integer repositoryId() {
        return this.repositoryId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String path;
        private @Nullable String ref;
        private Integer repositoryId;
        public Builder() {}
        public Builder(OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.path = defaults.path;
    	      this.ref = defaults.ref;
    	      this.repositoryId = defaults.repositoryId;
        }

        @CustomType.Setter
        public Builder path(String path) {
            if (path == null) {
              throw new MissingRequiredPropertyException("OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow", "path");
            }
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder ref(@Nullable String ref) {

            this.ref = ref;
            return this;
        }
        @CustomType.Setter
        public Builder repositoryId(Integer repositoryId) {
            if (repositoryId == null) {
              throw new MissingRequiredPropertyException("OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow", "repositoryId");
            }
            this.repositoryId = repositoryId;
            return this;
        }
        public OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow build() {
            final var _resultValue = new OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow();
            _resultValue.path = path;
            _resultValue.ref = ref;
            _resultValue.repositoryId = repositoryId;
            return _resultValue;
        }
    }
}
