// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.outputs.GetCodespacesOrganizationSecretsSecret;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCodespacesOrganizationSecretsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return list of secrets for the repository
     * 
     */
    private List<GetCodespacesOrganizationSecretsSecret> secrets;

    private GetCodespacesOrganizationSecretsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return list of secrets for the repository
     * 
     */
    public List<GetCodespacesOrganizationSecretsSecret> secrets() {
        return this.secrets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCodespacesOrganizationSecretsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<GetCodespacesOrganizationSecretsSecret> secrets;
        public Builder() {}
        public Builder(GetCodespacesOrganizationSecretsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.secrets = defaults.secrets;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetCodespacesOrganizationSecretsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder secrets(List<GetCodespacesOrganizationSecretsSecret> secrets) {
            if (secrets == null) {
              throw new MissingRequiredPropertyException("GetCodespacesOrganizationSecretsResult", "secrets");
            }
            this.secrets = secrets;
            return this;
        }
        public Builder secrets(GetCodespacesOrganizationSecretsSecret... secrets) {
            return secrets(List.of(secrets));
        }
        public GetCodespacesOrganizationSecretsResult build() {
            final var _resultValue = new GetCodespacesOrganizationSecretsResult();
            _resultValue.id = id;
            _resultValue.secrets = secrets;
            return _resultValue;
        }
    }
}
