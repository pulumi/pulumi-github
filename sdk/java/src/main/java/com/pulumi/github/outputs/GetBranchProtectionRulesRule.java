// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBranchProtectionRulesRule {
    /**
     * @return Identifies the protection rule pattern.
     * 
     */
    private String pattern;

    private GetBranchProtectionRulesRule() {}
    /**
     * @return Identifies the protection rule pattern.
     * 
     */
    public String pattern() {
        return this.pattern;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBranchProtectionRulesRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String pattern;
        public Builder() {}
        public Builder(GetBranchProtectionRulesRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.pattern = defaults.pattern;
        }

        @CustomType.Setter
        public Builder pattern(String pattern) {
            if (pattern == null) {
              throw new MissingRequiredPropertyException("GetBranchProtectionRulesRule", "pattern");
            }
            this.pattern = pattern;
            return this;
        }
        public GetBranchProtectionRulesRule build() {
            final var _resultValue = new GetBranchProtectionRulesRule();
            _resultValue.pattern = pattern;
            return _resultValue;
        }
    }
}
