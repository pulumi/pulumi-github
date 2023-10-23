// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RepositoryRulesetRulesCommitMessagePattern {
    /**
     * @return (String) The name of the ruleset.
     * 
     */
    private @Nullable String name;
    /**
     * @return (Boolean) If true, the rule will fail if the pattern matches.
     * 
     */
    private @Nullable Boolean negate;
    /**
     * @return (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.
     * 
     */
    private String operator;
    /**
     * @return (String) The pattern to match with.
     * 
     */
    private String pattern;

    private RepositoryRulesetRulesCommitMessagePattern() {}
    /**
     * @return (String) The name of the ruleset.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return (Boolean) If true, the rule will fail if the pattern matches.
     * 
     */
    public Optional<Boolean> negate() {
        return Optional.ofNullable(this.negate);
    }
    /**
     * @return (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.
     * 
     */
    public String operator() {
        return this.operator;
    }
    /**
     * @return (String) The pattern to match with.
     * 
     */
    public String pattern() {
        return this.pattern;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryRulesetRulesCommitMessagePattern defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable Boolean negate;
        private String operator;
        private String pattern;
        public Builder() {}
        public Builder(RepositoryRulesetRulesCommitMessagePattern defaults) {
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
            this.operator = Objects.requireNonNull(operator);
            return this;
        }
        @CustomType.Setter
        public Builder pattern(String pattern) {
            this.pattern = Objects.requireNonNull(pattern);
            return this;
        }
        public RepositoryRulesetRulesCommitMessagePattern build() {
            final var o = new RepositoryRulesetRulesCommitMessagePattern();
            o.name = name;
            o.negate = negate;
            o.operator = operator;
            o.pattern = pattern;
            return o;
        }
    }
}