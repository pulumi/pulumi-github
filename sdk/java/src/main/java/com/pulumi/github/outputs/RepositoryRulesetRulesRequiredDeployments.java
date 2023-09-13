// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class RepositoryRulesetRulesRequiredDeployments {
    /**
     * @return (List of String) The environments that must be successfully deployed to before branches can be merged.
     * 
     */
    private List<String> requiredDeploymentEnvironments;

    private RepositoryRulesetRulesRequiredDeployments() {}
    /**
     * @return (List of String) The environments that must be successfully deployed to before branches can be merged.
     * 
     */
    public List<String> requiredDeploymentEnvironments() {
        return this.requiredDeploymentEnvironments;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryRulesetRulesRequiredDeployments defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> requiredDeploymentEnvironments;
        public Builder() {}
        public Builder(RepositoryRulesetRulesRequiredDeployments defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.requiredDeploymentEnvironments = defaults.requiredDeploymentEnvironments;
        }

        @CustomType.Setter
        public Builder requiredDeploymentEnvironments(List<String> requiredDeploymentEnvironments) {
            this.requiredDeploymentEnvironments = Objects.requireNonNull(requiredDeploymentEnvironments);
            return this;
        }
        public Builder requiredDeploymentEnvironments(String... requiredDeploymentEnvironments) {
            return requiredDeploymentEnvironments(List.of(requiredDeploymentEnvironments));
        }
        public RepositoryRulesetRulesRequiredDeployments build() {
            final var o = new RepositoryRulesetRulesRequiredDeployments();
            o.requiredDeploymentEnvironments = requiredDeploymentEnvironments;
            return o;
        }
    }
}
