// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchDefaultArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchDefaultArgs Empty = new BranchDefaultArgs();

    /**
     * The branch (e.g. `main`)
     * 
     */
    @Import(name="branch", required=true)
    private Output<String> branch;

    /**
     * @return The branch (e.g. `main`)
     * 
     */
    public Output<String> branch() {
        return this.branch;
    }

    /**
     * Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
     * 
     */
    @Import(name="rename")
    private @Nullable Output<Boolean> rename;

    /**
     * @return Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> rename() {
        return Optional.ofNullable(this.rename);
    }

    /**
     * The GitHub repository
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The GitHub repository
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private BranchDefaultArgs() {}

    private BranchDefaultArgs(BranchDefaultArgs $) {
        this.branch = $.branch;
        this.rename = $.rename;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchDefaultArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchDefaultArgs $;

        public Builder() {
            $ = new BranchDefaultArgs();
        }

        public Builder(BranchDefaultArgs defaults) {
            $ = new BranchDefaultArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branch The branch (e.g. `main`)
         * 
         * @return builder
         * 
         */
        public Builder branch(Output<String> branch) {
            $.branch = branch;
            return this;
        }

        /**
         * @param branch The branch (e.g. `main`)
         * 
         * @return builder
         * 
         */
        public Builder branch(String branch) {
            return branch(Output.of(branch));
        }

        /**
         * @param rename Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder rename(@Nullable Output<Boolean> rename) {
            $.rename = rename;
            return this;
        }

        /**
         * @param rename Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder rename(Boolean rename) {
            return rename(Output.of(rename));
        }

        /**
         * @param repository The GitHub repository
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The GitHub repository
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public BranchDefaultArgs build() {
            if ($.branch == null) {
                throw new MissingRequiredPropertyException("BranchDefaultArgs", "branch");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("BranchDefaultArgs", "repository");
            }
            return $;
        }
    }

}
