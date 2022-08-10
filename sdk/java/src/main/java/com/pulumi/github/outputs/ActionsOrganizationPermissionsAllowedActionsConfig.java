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
public final class ActionsOrganizationPermissionsAllowedActionsConfig {
    /**
     * @return Whether GitHub-owned actions are allowed in the organization.
     * 
     */
    private final Boolean githubOwnedAllowed;
    /**
     * @return Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*.&#34;
     * 
     */
    private final @Nullable List<String> patternsAlloweds;
    /**
     * @return Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
     * 
     */
    private final @Nullable Boolean verifiedAllowed;

    @CustomType.Constructor
    private ActionsOrganizationPermissionsAllowedActionsConfig(
        @CustomType.Parameter("githubOwnedAllowed") Boolean githubOwnedAllowed,
        @CustomType.Parameter("patternsAlloweds") @Nullable List<String> patternsAlloweds,
        @CustomType.Parameter("verifiedAllowed") @Nullable Boolean verifiedAllowed) {
        this.githubOwnedAllowed = githubOwnedAllowed;
        this.patternsAlloweds = patternsAlloweds;
        this.verifiedAllowed = verifiedAllowed;
    }

    /**
     * @return Whether GitHub-owned actions are allowed in the organization.
     * 
     */
    public Boolean githubOwnedAllowed() {
        return this.githubOwnedAllowed;
    }
    /**
     * @return Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*.&#34;
     * 
     */
    public List<String> patternsAlloweds() {
        return this.patternsAlloweds == null ? List.of() : this.patternsAlloweds;
    }
    /**
     * @return Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
     * 
     */
    public Optional<Boolean> verifiedAllowed() {
        return Optional.ofNullable(this.verifiedAllowed);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ActionsOrganizationPermissionsAllowedActionsConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Boolean githubOwnedAllowed;
        private @Nullable List<String> patternsAlloweds;
        private @Nullable Boolean verifiedAllowed;

        public Builder() {
    	      // Empty
        }

        public Builder(ActionsOrganizationPermissionsAllowedActionsConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.githubOwnedAllowed = defaults.githubOwnedAllowed;
    	      this.patternsAlloweds = defaults.patternsAlloweds;
    	      this.verifiedAllowed = defaults.verifiedAllowed;
        }

        public Builder githubOwnedAllowed(Boolean githubOwnedAllowed) {
            this.githubOwnedAllowed = Objects.requireNonNull(githubOwnedAllowed);
            return this;
        }
        public Builder patternsAlloweds(@Nullable List<String> patternsAlloweds) {
            this.patternsAlloweds = patternsAlloweds;
            return this;
        }
        public Builder patternsAlloweds(String... patternsAlloweds) {
            return patternsAlloweds(List.of(patternsAlloweds));
        }
        public Builder verifiedAllowed(@Nullable Boolean verifiedAllowed) {
            this.verifiedAllowed = verifiedAllowed;
            return this;
        }        public ActionsOrganizationPermissionsAllowedActionsConfig build() {
            return new ActionsOrganizationPermissionsAllowedActionsConfig(githubOwnedAllowed, patternsAlloweds, verifiedAllowed);
        }
    }
}