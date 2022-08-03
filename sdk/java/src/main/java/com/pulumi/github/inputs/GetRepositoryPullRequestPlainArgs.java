// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoryPullRequestPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoryPullRequestPlainArgs Empty = new GetRepositoryPullRequestPlainArgs();

    /**
     * Name of the base repository to retrieve the Pull Request from.
     * 
     */
    @Import(name="baseRepository", required=true)
    private String baseRepository;

    /**
     * @return Name of the base repository to retrieve the Pull Request from.
     * 
     */
    public String baseRepository() {
        return this.baseRepository;
    }

    /**
     * The number of the Pull Request within the repository.
     * 
     */
    @Import(name="number", required=true)
    private Integer number;

    /**
     * @return The number of the Pull Request within the repository.
     * 
     */
    public Integer number() {
        return this.number;
    }

    /**
     * Owner of the repository. If not provided, the provider&#39;s default owner is used.
     * 
     */
    @Import(name="owner")
    private @Nullable String owner;

    /**
     * @return Owner of the repository. If not provided, the provider&#39;s default owner is used.
     * 
     */
    public Optional<String> owner() {
        return Optional.ofNullable(this.owner);
    }

    private GetRepositoryPullRequestPlainArgs() {}

    private GetRepositoryPullRequestPlainArgs(GetRepositoryPullRequestPlainArgs $) {
        this.baseRepository = $.baseRepository;
        this.number = $.number;
        this.owner = $.owner;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoryPullRequestPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoryPullRequestPlainArgs $;

        public Builder() {
            $ = new GetRepositoryPullRequestPlainArgs();
        }

        public Builder(GetRepositoryPullRequestPlainArgs defaults) {
            $ = new GetRepositoryPullRequestPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param baseRepository Name of the base repository to retrieve the Pull Request from.
         * 
         * @return builder
         * 
         */
        public Builder baseRepository(String baseRepository) {
            $.baseRepository = baseRepository;
            return this;
        }

        /**
         * @param number The number of the Pull Request within the repository.
         * 
         * @return builder
         * 
         */
        public Builder number(Integer number) {
            $.number = number;
            return this;
        }

        /**
         * @param owner Owner of the repository. If not provided, the provider&#39;s default owner is used.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable String owner) {
            $.owner = owner;
            return this;
        }

        public GetRepositoryPullRequestPlainArgs build() {
            $.baseRepository = Objects.requireNonNull($.baseRepository, "expected parameter 'baseRepository' to be non-null");
            $.number = Objects.requireNonNull($.number, "expected parameter 'number' to be non-null");
            return $;
        }
    }

}
