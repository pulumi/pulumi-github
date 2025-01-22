// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoryPullRequestsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoryPullRequestsArgs Empty = new GetRepositoryPullRequestsArgs();

    /**
     * If set, filters Pull Requests by base branch name.
     * 
     */
    @Import(name="baseRef")
    private @Nullable Output<String> baseRef;

    /**
     * @return If set, filters Pull Requests by base branch name.
     * 
     */
    public Optional<Output<String>> baseRef() {
        return Optional.ofNullable(this.baseRef);
    }

    /**
     * Name of the base repository to retrieve the Pull Requests from.
     * 
     */
    @Import(name="baseRepository", required=true)
    private Output<String> baseRepository;

    /**
     * @return Name of the base repository to retrieve the Pull Requests from.
     * 
     */
    public Output<String> baseRepository() {
        return this.baseRepository;
    }

    /**
     * If set, filters Pull Requests by head user or head organization and branch name in the format of &#34;user:ref-name&#34; or &#34;organization:ref-name&#34;. For example: &#34;github:new-script-format&#34; or &#34;octocat:test-branch&#34;.
     * 
     */
    @Import(name="headRef")
    private @Nullable Output<String> headRef;

    /**
     * @return If set, filters Pull Requests by head user or head organization and branch name in the format of &#34;user:ref-name&#34; or &#34;organization:ref-name&#34;. For example: &#34;github:new-script-format&#34; or &#34;octocat:test-branch&#34;.
     * 
     */
    public Optional<Output<String>> headRef() {
        return Optional.ofNullable(this.headRef);
    }

    /**
     * Owner of the repository. If not provided, the provider&#39;s default owner is used.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return Owner of the repository. If not provided, the provider&#39;s default owner is used.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * If set, indicates what to sort results by. Can be either &#34;created&#34;, &#34;updated&#34;, &#34;popularity&#34; (comment count) or &#34;long-running&#34; (age, filtering by pulls updated in the last month). Default: &#34;created&#34;.
     * 
     */
    @Import(name="sortBy")
    private @Nullable Output<String> sortBy;

    /**
     * @return If set, indicates what to sort results by. Can be either &#34;created&#34;, &#34;updated&#34;, &#34;popularity&#34; (comment count) or &#34;long-running&#34; (age, filtering by pulls updated in the last month). Default: &#34;created&#34;.
     * 
     */
    public Optional<Output<String>> sortBy() {
        return Optional.ofNullable(this.sortBy);
    }

    /**
     * If set, controls the direction of the sort. Can be either &#34;asc&#34; or &#34;desc&#34;. Default: &#34;asc&#34;.
     * 
     */
    @Import(name="sortDirection")
    private @Nullable Output<String> sortDirection;

    /**
     * @return If set, controls the direction of the sort. Can be either &#34;asc&#34; or &#34;desc&#34;. Default: &#34;asc&#34;.
     * 
     */
    public Optional<Output<String>> sortDirection() {
        return Optional.ofNullable(this.sortDirection);
    }

    /**
     * If set, filters Pull Requests by state. Can be &#34;open&#34;, &#34;closed&#34;, or &#34;all&#34;. Default: &#34;open&#34;.
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return If set, filters Pull Requests by state. Can be &#34;open&#34;, &#34;closed&#34;, or &#34;all&#34;. Default: &#34;open&#34;.
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    private GetRepositoryPullRequestsArgs() {}

    private GetRepositoryPullRequestsArgs(GetRepositoryPullRequestsArgs $) {
        this.baseRef = $.baseRef;
        this.baseRepository = $.baseRepository;
        this.headRef = $.headRef;
        this.owner = $.owner;
        this.sortBy = $.sortBy;
        this.sortDirection = $.sortDirection;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoryPullRequestsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoryPullRequestsArgs $;

        public Builder() {
            $ = new GetRepositoryPullRequestsArgs();
        }

        public Builder(GetRepositoryPullRequestsArgs defaults) {
            $ = new GetRepositoryPullRequestsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param baseRef If set, filters Pull Requests by base branch name.
         * 
         * @return builder
         * 
         */
        public Builder baseRef(@Nullable Output<String> baseRef) {
            $.baseRef = baseRef;
            return this;
        }

        /**
         * @param baseRef If set, filters Pull Requests by base branch name.
         * 
         * @return builder
         * 
         */
        public Builder baseRef(String baseRef) {
            return baseRef(Output.of(baseRef));
        }

        /**
         * @param baseRepository Name of the base repository to retrieve the Pull Requests from.
         * 
         * @return builder
         * 
         */
        public Builder baseRepository(Output<String> baseRepository) {
            $.baseRepository = baseRepository;
            return this;
        }

        /**
         * @param baseRepository Name of the base repository to retrieve the Pull Requests from.
         * 
         * @return builder
         * 
         */
        public Builder baseRepository(String baseRepository) {
            return baseRepository(Output.of(baseRepository));
        }

        /**
         * @param headRef If set, filters Pull Requests by head user or head organization and branch name in the format of &#34;user:ref-name&#34; or &#34;organization:ref-name&#34;. For example: &#34;github:new-script-format&#34; or &#34;octocat:test-branch&#34;.
         * 
         * @return builder
         * 
         */
        public Builder headRef(@Nullable Output<String> headRef) {
            $.headRef = headRef;
            return this;
        }

        /**
         * @param headRef If set, filters Pull Requests by head user or head organization and branch name in the format of &#34;user:ref-name&#34; or &#34;organization:ref-name&#34;. For example: &#34;github:new-script-format&#34; or &#34;octocat:test-branch&#34;.
         * 
         * @return builder
         * 
         */
        public Builder headRef(String headRef) {
            return headRef(Output.of(headRef));
        }

        /**
         * @param owner Owner of the repository. If not provided, the provider&#39;s default owner is used.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner Owner of the repository. If not provided, the provider&#39;s default owner is used.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param sortBy If set, indicates what to sort results by. Can be either &#34;created&#34;, &#34;updated&#34;, &#34;popularity&#34; (comment count) or &#34;long-running&#34; (age, filtering by pulls updated in the last month). Default: &#34;created&#34;.
         * 
         * @return builder
         * 
         */
        public Builder sortBy(@Nullable Output<String> sortBy) {
            $.sortBy = sortBy;
            return this;
        }

        /**
         * @param sortBy If set, indicates what to sort results by. Can be either &#34;created&#34;, &#34;updated&#34;, &#34;popularity&#34; (comment count) or &#34;long-running&#34; (age, filtering by pulls updated in the last month). Default: &#34;created&#34;.
         * 
         * @return builder
         * 
         */
        public Builder sortBy(String sortBy) {
            return sortBy(Output.of(sortBy));
        }

        /**
         * @param sortDirection If set, controls the direction of the sort. Can be either &#34;asc&#34; or &#34;desc&#34;. Default: &#34;asc&#34;.
         * 
         * @return builder
         * 
         */
        public Builder sortDirection(@Nullable Output<String> sortDirection) {
            $.sortDirection = sortDirection;
            return this;
        }

        /**
         * @param sortDirection If set, controls the direction of the sort. Can be either &#34;asc&#34; or &#34;desc&#34;. Default: &#34;asc&#34;.
         * 
         * @return builder
         * 
         */
        public Builder sortDirection(String sortDirection) {
            return sortDirection(Output.of(sortDirection));
        }

        /**
         * @param state If set, filters Pull Requests by state. Can be &#34;open&#34;, &#34;closed&#34;, or &#34;all&#34;. Default: &#34;open&#34;.
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state If set, filters Pull Requests by state. Can be &#34;open&#34;, &#34;closed&#34;, or &#34;all&#34;. Default: &#34;open&#34;.
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public GetRepositoryPullRequestsArgs build() {
            if ($.baseRepository == null) {
                throw new MissingRequiredPropertyException("GetRepositoryPullRequestsArgs", "baseRepository");
            }
            return $;
        }
    }

}