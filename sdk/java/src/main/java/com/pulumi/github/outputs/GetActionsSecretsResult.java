// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.outputs.GetActionsSecretsSecret;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetActionsSecretsResult {
    private String fullName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Secret name
     * 
     */
    private String name;
    /**
     * @return list of secrets for the repository
     * 
     */
    private List<GetActionsSecretsSecret> secrets;

    private GetActionsSecretsResult() {}
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
     * @return Secret name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return list of secrets for the repository
     * 
     */
    public List<GetActionsSecretsSecret> secrets() {
        return this.secrets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetActionsSecretsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fullName;
        private String id;
        private String name;
        private List<GetActionsSecretsSecret> secrets;
        public Builder() {}
        public Builder(GetActionsSecretsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fullName = defaults.fullName;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.secrets = defaults.secrets;
        }

        @CustomType.Setter
        public Builder fullName(String fullName) {
            if (fullName == null) {
              throw new MissingRequiredPropertyException("GetActionsSecretsResult", "fullName");
            }
            this.fullName = fullName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetActionsSecretsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetActionsSecretsResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder secrets(List<GetActionsSecretsSecret> secrets) {
            if (secrets == null) {
              throw new MissingRequiredPropertyException("GetActionsSecretsResult", "secrets");
            }
            this.secrets = secrets;
            return this;
        }
        public Builder secrets(GetActionsSecretsSecret... secrets) {
            return secrets(List.of(secrets));
        }
        public GetActionsSecretsResult build() {
            final var _resultValue = new GetActionsSecretsResult();
            _resultValue.fullName = fullName;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.secrets = secrets;
            return _resultValue;
        }
    }
}
