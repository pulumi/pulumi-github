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
public final class BranchProtectionRestrictPush {
    /**
     * @return Boolean, setting this to `false` allows people, teams, or apps to create new branches matching this rule. Defaults to `true`.
     * 
     */
    private @Nullable Boolean blocksCreations;
    /**
     * @return A list of actor Names/IDs that may push to the branch. Actor names must either begin with a &#34;/&#34; for users or the organization name followed by a &#34;/&#34; for teams. Organization administrators, repository administrators, and users with the Maintain role on the repository can always push when all other requirements have passed.
     * 
     */
    private @Nullable List<String> pushAllowances;

    private BranchProtectionRestrictPush() {}
    /**
     * @return Boolean, setting this to `false` allows people, teams, or apps to create new branches matching this rule. Defaults to `true`.
     * 
     */
    public Optional<Boolean> blocksCreations() {
        return Optional.ofNullable(this.blocksCreations);
    }
    /**
     * @return A list of actor Names/IDs that may push to the branch. Actor names must either begin with a &#34;/&#34; for users or the organization name followed by a &#34;/&#34; for teams. Organization administrators, repository administrators, and users with the Maintain role on the repository can always push when all other requirements have passed.
     * 
     */
    public List<String> pushAllowances() {
        return this.pushAllowances == null ? List.of() : this.pushAllowances;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchProtectionRestrictPush defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean blocksCreations;
        private @Nullable List<String> pushAllowances;
        public Builder() {}
        public Builder(BranchProtectionRestrictPush defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.blocksCreations = defaults.blocksCreations;
    	      this.pushAllowances = defaults.pushAllowances;
        }

        @CustomType.Setter
        public Builder blocksCreations(@Nullable Boolean blocksCreations) {

            this.blocksCreations = blocksCreations;
            return this;
        }
        @CustomType.Setter
        public Builder pushAllowances(@Nullable List<String> pushAllowances) {

            this.pushAllowances = pushAllowances;
            return this;
        }
        public Builder pushAllowances(String... pushAllowances) {
            return pushAllowances(List.of(pushAllowances));
        }
        public BranchProtectionRestrictPush build() {
            final var _resultValue = new BranchProtectionRestrictPush();
            _resultValue.blocksCreations = blocksCreations;
            _resultValue.pushAllowances = pushAllowances;
            return _resultValue;
        }
    }
}
