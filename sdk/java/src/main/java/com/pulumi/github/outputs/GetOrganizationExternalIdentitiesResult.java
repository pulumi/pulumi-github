// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.outputs.GetOrganizationExternalIdentitiesIdentity;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOrganizationExternalIdentitiesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return An Array of identities returned from GitHub
     * 
     */
    private List<GetOrganizationExternalIdentitiesIdentity> identities;

    private GetOrganizationExternalIdentitiesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return An Array of identities returned from GitHub
     * 
     */
    public List<GetOrganizationExternalIdentitiesIdentity> identities() {
        return this.identities;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrganizationExternalIdentitiesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<GetOrganizationExternalIdentitiesIdentity> identities;
        public Builder() {}
        public Builder(GetOrganizationExternalIdentitiesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.identities = defaults.identities;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetOrganizationExternalIdentitiesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder identities(List<GetOrganizationExternalIdentitiesIdentity> identities) {
            if (identities == null) {
              throw new MissingRequiredPropertyException("GetOrganizationExternalIdentitiesResult", "identities");
            }
            this.identities = identities;
            return this;
        }
        public Builder identities(GetOrganizationExternalIdentitiesIdentity... identities) {
            return identities(List.of(identities));
        }
        public GetOrganizationExternalIdentitiesResult build() {
            final var _resultValue = new GetOrganizationExternalIdentitiesResult();
            _resultValue.id = id;
            _resultValue.identities = identities;
            return _resultValue;
        }
    }
}
