// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.outputs.GetCodespacesSecretsSecret;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCodespacesSecretsResult {
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
     * @return list of codespaces secrets for the repository
     * 
     */
    private List<GetCodespacesSecretsSecret> secrets;

    private GetCodespacesSecretsResult() {}
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
     * @return list of codespaces secrets for the repository
     * 
     */
    public List<GetCodespacesSecretsSecret> secrets() {
        return this.secrets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCodespacesSecretsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fullName;
        private String id;
        private String name;
        private List<GetCodespacesSecretsSecret> secrets;
        public Builder() {}
        public Builder(GetCodespacesSecretsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fullName = defaults.fullName;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.secrets = defaults.secrets;
        }

        @CustomType.Setter
        public Builder fullName(String fullName) {
            if (fullName == null) {
              throw new MissingRequiredPropertyException("GetCodespacesSecretsResult", "fullName");
            }
            this.fullName = fullName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetCodespacesSecretsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetCodespacesSecretsResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder secrets(List<GetCodespacesSecretsSecret> secrets) {
            if (secrets == null) {
              throw new MissingRequiredPropertyException("GetCodespacesSecretsResult", "secrets");
            }
            this.secrets = secrets;
            return this;
        }
        public Builder secrets(GetCodespacesSecretsSecret... secrets) {
            return secrets(List.of(secrets));
        }
        public GetCodespacesSecretsResult build() {
            final var _resultValue = new GetCodespacesSecretsResult();
            _resultValue.fullName = fullName;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.secrets = secrets;
            return _resultValue;
        }
    }
}
