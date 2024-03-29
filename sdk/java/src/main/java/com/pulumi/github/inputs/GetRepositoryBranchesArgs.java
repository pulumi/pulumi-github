// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoryBranchesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoryBranchesArgs Empty = new GetRepositoryBranchesArgs();

    /**
     * . If true, the `branches` attributes will be populated only with non protected branches. Default: `false`.
     * 
     */
    @Import(name="onlyNonProtectedBranches")
    private @Nullable Output<Boolean> onlyNonProtectedBranches;

    /**
     * @return . If true, the `branches` attributes will be populated only with non protected branches. Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> onlyNonProtectedBranches() {
        return Optional.ofNullable(this.onlyNonProtectedBranches);
    }

    /**
     * . If true, the `branches` attributes will be populated only with protected branches. Default: `false`.
     * 
     */
    @Import(name="onlyProtectedBranches")
    private @Nullable Output<Boolean> onlyProtectedBranches;

    /**
     * @return . If true, the `branches` attributes will be populated only with protected branches. Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> onlyProtectedBranches() {
        return Optional.ofNullable(this.onlyProtectedBranches);
    }

    /**
     * Name of the repository to retrieve the branches from.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return Name of the repository to retrieve the branches from.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private GetRepositoryBranchesArgs() {}

    private GetRepositoryBranchesArgs(GetRepositoryBranchesArgs $) {
        this.onlyNonProtectedBranches = $.onlyNonProtectedBranches;
        this.onlyProtectedBranches = $.onlyProtectedBranches;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoryBranchesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoryBranchesArgs $;

        public Builder() {
            $ = new GetRepositoryBranchesArgs();
        }

        public Builder(GetRepositoryBranchesArgs defaults) {
            $ = new GetRepositoryBranchesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param onlyNonProtectedBranches . If true, the `branches` attributes will be populated only with non protected branches. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder onlyNonProtectedBranches(@Nullable Output<Boolean> onlyNonProtectedBranches) {
            $.onlyNonProtectedBranches = onlyNonProtectedBranches;
            return this;
        }

        /**
         * @param onlyNonProtectedBranches . If true, the `branches` attributes will be populated only with non protected branches. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder onlyNonProtectedBranches(Boolean onlyNonProtectedBranches) {
            return onlyNonProtectedBranches(Output.of(onlyNonProtectedBranches));
        }

        /**
         * @param onlyProtectedBranches . If true, the `branches` attributes will be populated only with protected branches. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder onlyProtectedBranches(@Nullable Output<Boolean> onlyProtectedBranches) {
            $.onlyProtectedBranches = onlyProtectedBranches;
            return this;
        }

        /**
         * @param onlyProtectedBranches . If true, the `branches` attributes will be populated only with protected branches. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder onlyProtectedBranches(Boolean onlyProtectedBranches) {
            return onlyProtectedBranches(Output.of(onlyProtectedBranches));
        }

        /**
         * @param repository Name of the repository to retrieve the branches from.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository Name of the repository to retrieve the branches from.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public GetRepositoryBranchesArgs build() {
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("GetRepositoryBranchesArgs", "repository");
            }
            return $;
        }
    }

}
