// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicy {
    /**
     * @return Id of the policy.
     * 
     */
    private String id;
    /**
     * @return The name pattern that branches must match in order to deploy to the environment.
     * 
     */
    private String name;

    private GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicy() {}
    /**
     * @return Id of the policy.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name pattern that branches must match in order to deploy to the environment.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicy", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicy", "name");
            }
            this.name = name;
            return this;
        }
        public GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicy build() {
            final var _resultValue = new GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicy();
            _resultValue.id = id;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
