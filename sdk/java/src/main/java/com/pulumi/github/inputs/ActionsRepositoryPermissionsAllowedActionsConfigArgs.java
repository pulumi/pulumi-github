// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ActionsRepositoryPermissionsAllowedActionsConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActionsRepositoryPermissionsAllowedActionsConfigArgs Empty = new ActionsRepositoryPermissionsAllowedActionsConfigArgs();

    /**
     * Whether GitHub-owned actions are allowed in the repository.
     * 
     */
    @Import(name="githubOwnedAllowed", required=true)
    private Output<Boolean> githubOwnedAllowed;

    /**
     * @return Whether GitHub-owned actions are allowed in the repository.
     * 
     */
    public Output<Boolean> githubOwnedAllowed() {
        return this.githubOwnedAllowed;
    }

    /**
     * Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*.&#34;
     * 
     */
    @Import(name="patternsAlloweds")
    private @Nullable Output<List<String>> patternsAlloweds;

    /**
     * @return Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*.&#34;
     * 
     */
    public Optional<Output<List<String>>> patternsAlloweds() {
        return Optional.ofNullable(this.patternsAlloweds);
    }

    /**
     * Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
     * 
     */
    @Import(name="verifiedAllowed")
    private @Nullable Output<Boolean> verifiedAllowed;

    /**
     * @return Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
     * 
     */
    public Optional<Output<Boolean>> verifiedAllowed() {
        return Optional.ofNullable(this.verifiedAllowed);
    }

    private ActionsRepositoryPermissionsAllowedActionsConfigArgs() {}

    private ActionsRepositoryPermissionsAllowedActionsConfigArgs(ActionsRepositoryPermissionsAllowedActionsConfigArgs $) {
        this.githubOwnedAllowed = $.githubOwnedAllowed;
        this.patternsAlloweds = $.patternsAlloweds;
        this.verifiedAllowed = $.verifiedAllowed;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActionsRepositoryPermissionsAllowedActionsConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActionsRepositoryPermissionsAllowedActionsConfigArgs $;

        public Builder() {
            $ = new ActionsRepositoryPermissionsAllowedActionsConfigArgs();
        }

        public Builder(ActionsRepositoryPermissionsAllowedActionsConfigArgs defaults) {
            $ = new ActionsRepositoryPermissionsAllowedActionsConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param githubOwnedAllowed Whether GitHub-owned actions are allowed in the repository.
         * 
         * @return builder
         * 
         */
        public Builder githubOwnedAllowed(Output<Boolean> githubOwnedAllowed) {
            $.githubOwnedAllowed = githubOwnedAllowed;
            return this;
        }

        /**
         * @param githubOwnedAllowed Whether GitHub-owned actions are allowed in the repository.
         * 
         * @return builder
         * 
         */
        public Builder githubOwnedAllowed(Boolean githubOwnedAllowed) {
            return githubOwnedAllowed(Output.of(githubOwnedAllowed));
        }

        /**
         * @param patternsAlloweds Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*.&#34;
         * 
         * @return builder
         * 
         */
        public Builder patternsAlloweds(@Nullable Output<List<String>> patternsAlloweds) {
            $.patternsAlloweds = patternsAlloweds;
            return this;
        }

        /**
         * @param patternsAlloweds Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*.&#34;
         * 
         * @return builder
         * 
         */
        public Builder patternsAlloweds(List<String> patternsAlloweds) {
            return patternsAlloweds(Output.of(patternsAlloweds));
        }

        /**
         * @param patternsAlloweds Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*.&#34;
         * 
         * @return builder
         * 
         */
        public Builder patternsAlloweds(String... patternsAlloweds) {
            return patternsAlloweds(List.of(patternsAlloweds));
        }

        /**
         * @param verifiedAllowed Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
         * 
         * @return builder
         * 
         */
        public Builder verifiedAllowed(@Nullable Output<Boolean> verifiedAllowed) {
            $.verifiedAllowed = verifiedAllowed;
            return this;
        }

        /**
         * @param verifiedAllowed Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
         * 
         * @return builder
         * 
         */
        public Builder verifiedAllowed(Boolean verifiedAllowed) {
            return verifiedAllowed(Output.of(verifiedAllowed));
        }

        public ActionsRepositoryPermissionsAllowedActionsConfigArgs build() {
            $.githubOwnedAllowed = Objects.requireNonNull($.githubOwnedAllowed, "expected parameter 'githubOwnedAllowed' to be non-null");
            return $;
        }
    }

}