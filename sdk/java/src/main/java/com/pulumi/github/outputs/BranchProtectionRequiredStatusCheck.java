// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BranchProtectionRequiredStatusCheck {
    /**
     * @return The list of status checks to require in order to merge into this branch. No status checks are required by default.
     * 
     */
    private @Nullable List<String> contexts;
    /**
     * @return Require branches to be up to date before merging. Defaults to `false`.
     * 
     */
    private @Nullable Boolean strict;

    private BranchProtectionRequiredStatusCheck() {}
    /**
     * @return The list of status checks to require in order to merge into this branch. No status checks are required by default.
     * 
     */
    public List<String> contexts() {
        return this.contexts == null ? List.of() : this.contexts;
    }
    /**
     * @return Require branches to be up to date before merging. Defaults to `false`.
     * 
     */
    public Optional<Boolean> strict() {
        return Optional.ofNullable(this.strict);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchProtectionRequiredStatusCheck defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> contexts;
        private @Nullable Boolean strict;
        public Builder() {}
        public Builder(BranchProtectionRequiredStatusCheck defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contexts = defaults.contexts;
    	      this.strict = defaults.strict;
        }

        @CustomType.Setter
        public Builder contexts(@Nullable List<String> contexts) {
            this.contexts = contexts;
            return this;
        }
        public Builder contexts(String... contexts) {
            return contexts(List.of(contexts));
        }
        @CustomType.Setter
        public Builder strict(@Nullable Boolean strict) {
            this.strict = strict;
            return this;
        }
        public BranchProtectionRequiredStatusCheck build() {
            final var _resultValue = new BranchProtectionRequiredStatusCheck();
            _resultValue.contexts = contexts;
            _resultValue.strict = strict;
            return _resultValue;
        }
    }
}
