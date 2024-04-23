// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OrganizationRulesetRulesBranchNamePattern {
    /**
     * @return (String) The name of the ruleset.
     * 
     */
    private @Nullable String name;
    /**
     * @return If true, the rule will fail if the pattern matches.
     * 
     */
    private @Nullable Boolean negate;
    /**
     * @return The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.
     * 
     */
    private String operator;
    /**
     * @return The pattern to match with.
     * 
     */
    private String pattern;

    private OrganizationRulesetRulesBranchNamePattern() {}
    /**
     * @return (String) The name of the ruleset.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return If true, the rule will fail if the pattern matches.
     * 
     */
    public Optional<Boolean> negate() {
        return Optional.ofNullable(this.negate);
    }
    /**
     * @return The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.
     * 
     */
    public String operator() {
        return this.operator;
    }
    /**
     * @return The pattern to match with.
     * 
     */
    public String pattern() {
        return this.pattern;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrganizationRulesetRulesBranchNamePattern defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable Boolean negate;
        private String operator;
        private String pattern;
        public Builder() {}
        public Builder(OrganizationRulesetRulesBranchNamePattern defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.negate = defaults.negate;
    	      this.operator = defaults.operator;
    	      this.pattern = defaults.pattern;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder negate(@Nullable Boolean negate) {

            this.negate = negate;
            return this;
        }
        @CustomType.Setter
        public Builder operator(String operator) {
            if (operator == null) {
              throw new MissingRequiredPropertyException("OrganizationRulesetRulesBranchNamePattern", "operator");
            }
            this.operator = operator;
            return this;
        }
        @CustomType.Setter
        public Builder pattern(String pattern) {
            if (pattern == null) {
              throw new MissingRequiredPropertyException("OrganizationRulesetRulesBranchNamePattern", "pattern");
            }
            this.pattern = pattern;
            return this;
        }
        public OrganizationRulesetRulesBranchNamePattern build() {
            final var _resultValue = new OrganizationRulesetRulesBranchNamePattern();
            _resultValue.name = name;
            _resultValue.negate = negate;
            _resultValue.operator = operator;
            _resultValue.pattern = pattern;
            return _resultValue;
        }
    }
}
