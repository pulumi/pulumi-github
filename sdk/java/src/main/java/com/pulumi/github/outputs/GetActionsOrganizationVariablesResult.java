// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.GetActionsOrganizationVariablesVariable;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetActionsOrganizationVariablesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return list of variables for the repository
     * 
     */
    private List<GetActionsOrganizationVariablesVariable> variables;

    private GetActionsOrganizationVariablesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return list of variables for the repository
     * 
     */
    public List<GetActionsOrganizationVariablesVariable> variables() {
        return this.variables;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetActionsOrganizationVariablesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<GetActionsOrganizationVariablesVariable> variables;
        public Builder() {}
        public Builder(GetActionsOrganizationVariablesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.variables = defaults.variables;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder variables(List<GetActionsOrganizationVariablesVariable> variables) {
            this.variables = Objects.requireNonNull(variables);
            return this;
        }
        public Builder variables(GetActionsOrganizationVariablesVariable... variables) {
            return variables(List.of(variables));
        }
        public GetActionsOrganizationVariablesResult build() {
            final var _resultValue = new GetActionsOrganizationVariablesResult();
            _resultValue.id = id;
            _resultValue.variables = variables;
            return _resultValue;
        }
    }
}
