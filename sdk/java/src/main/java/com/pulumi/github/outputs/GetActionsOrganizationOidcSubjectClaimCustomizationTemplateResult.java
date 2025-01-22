// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of OpenID Connect claim keys.
     * 
     */
    private List<String> includeClaimKeys;

    private GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of OpenID Connect claim keys.
     * 
     */
    public List<String> includeClaimKeys() {
        return this.includeClaimKeys;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> includeClaimKeys;
        public Builder() {}
        public Builder(GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.includeClaimKeys = defaults.includeClaimKeys;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder includeClaimKeys(List<String> includeClaimKeys) {
            if (includeClaimKeys == null) {
              throw new MissingRequiredPropertyException("GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult", "includeClaimKeys");
            }
            this.includeClaimKeys = includeClaimKeys;
            return this;
        }
        public Builder includeClaimKeys(String... includeClaimKeys) {
            return includeClaimKeys(List.of(includeClaimKeys));
        }
        public GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult build() {
            final var _resultValue = new GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult();
            _resultValue.id = id;
            _resultValue.includeClaimKeys = includeClaimKeys;
            return _resultValue;
        }
    }
}