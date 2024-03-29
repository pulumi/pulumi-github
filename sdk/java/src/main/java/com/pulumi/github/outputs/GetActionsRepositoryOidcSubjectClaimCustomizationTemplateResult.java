// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult {
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
    private String name;
    /**
     * @return Whether the repository uses the default template.
     * 
     */
    private Boolean useDefault;

    private GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult() {}
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
    public String name() {
        return this.name;
    }
    /**
     * @return Whether the repository uses the default template.
     * 
     */
    public Boolean useDefault() {
        return this.useDefault;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> includeClaimKeys;
        private String name;
        private Boolean useDefault;
        public Builder() {}
        public Builder(GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.includeClaimKeys = defaults.includeClaimKeys;
    	      this.name = defaults.name;
    	      this.useDefault = defaults.useDefault;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder includeClaimKeys(List<String> includeClaimKeys) {
            if (includeClaimKeys == null) {
              throw new MissingRequiredPropertyException("GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult", "includeClaimKeys");
            }
            this.includeClaimKeys = includeClaimKeys;
            return this;
        }
        public Builder includeClaimKeys(String... includeClaimKeys) {
            return includeClaimKeys(List.of(includeClaimKeys));
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder useDefault(Boolean useDefault) {
            if (useDefault == null) {
              throw new MissingRequiredPropertyException("GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult", "useDefault");
            }
            this.useDefault = useDefault;
            return this;
        }
        public GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult build() {
            final var _resultValue = new GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult();
            _resultValue.id = id;
            _resultValue.includeClaimKeys = includeClaimKeys;
            _resultValue.name = name;
            _resultValue.useDefault = useDefault;
            return _resultValue;
        }
    }
}
