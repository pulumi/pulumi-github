// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.outputs.GetActionsEnvironmentVariablesVariable;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetActionsEnvironmentVariablesResult {
    private String environment;
    private String fullName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Name of the variable
     * 
     */
    private String name;
    /**
     * @return list of variables for the environment
     * 
     */
    private List<GetActionsEnvironmentVariablesVariable> variables;

    private GetActionsEnvironmentVariablesResult() {}
    public String environment() {
        return this.environment;
    }
    public String fullName() {
        return this.fullName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Name of the variable
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return list of variables for the environment
     * 
     */
    public List<GetActionsEnvironmentVariablesVariable> variables() {
        return this.variables;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetActionsEnvironmentVariablesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String environment;
        private String fullName;
        private String id;
        private String name;
        private List<GetActionsEnvironmentVariablesVariable> variables;
        public Builder() {}
        public Builder(GetActionsEnvironmentVariablesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.environment = defaults.environment;
    	      this.fullName = defaults.fullName;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.variables = defaults.variables;
        }

        @CustomType.Setter
        public Builder environment(String environment) {
            if (environment == null) {
              throw new MissingRequiredPropertyException("GetActionsEnvironmentVariablesResult", "environment");
            }
            this.environment = environment;
            return this;
        }
        @CustomType.Setter
        public Builder fullName(String fullName) {
            if (fullName == null) {
              throw new MissingRequiredPropertyException("GetActionsEnvironmentVariablesResult", "fullName");
            }
            this.fullName = fullName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetActionsEnvironmentVariablesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetActionsEnvironmentVariablesResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder variables(List<GetActionsEnvironmentVariablesVariable> variables) {
            if (variables == null) {
              throw new MissingRequiredPropertyException("GetActionsEnvironmentVariablesResult", "variables");
            }
            this.variables = variables;
            return this;
        }
        public Builder variables(GetActionsEnvironmentVariablesVariable... variables) {
            return variables(List.of(variables));
        }
        public GetActionsEnvironmentVariablesResult build() {
            final var _resultValue = new GetActionsEnvironmentVariablesResult();
            _resultValue.environment = environment;
            _resultValue.fullName = fullName;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.variables = variables;
            return _resultValue;
        }
    }
}
