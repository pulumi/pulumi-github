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


public final class RepositoryTemplateArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryTemplateArgs Empty = new RepositoryTemplateArgs();

    /**
     * Whether the new repository should include all the branches from the template repository (defaults to false, which includes only the default branch from the template).
     * 
     */
    @Import(name="includeAllBranches")
    private @Nullable Output<Boolean> includeAllBranches;

    /**
     * @return Whether the new repository should include all the branches from the template repository (defaults to false, which includes only the default branch from the template).
     * 
     */
    public Optional<Output<Boolean>> includeAllBranches() {
        return Optional.ofNullable(this.includeAllBranches);
    }

    /**
     * The GitHub organization or user the template repository is owned by.
     * 
     */
    @Import(name="owner", required=true)
    private Output<String> owner;

    /**
     * @return The GitHub organization or user the template repository is owned by.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }

    /**
     * The name of the template repository.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The name of the template repository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private RepositoryTemplateArgs() {}

    private RepositoryTemplateArgs(RepositoryTemplateArgs $) {
        this.includeAllBranches = $.includeAllBranches;
        this.owner = $.owner;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryTemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryTemplateArgs $;

        public Builder() {
            $ = new RepositoryTemplateArgs();
        }

        public Builder(RepositoryTemplateArgs defaults) {
            $ = new RepositoryTemplateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param includeAllBranches Whether the new repository should include all the branches from the template repository (defaults to false, which includes only the default branch from the template).
         * 
         * @return builder
         * 
         */
        public Builder includeAllBranches(@Nullable Output<Boolean> includeAllBranches) {
            $.includeAllBranches = includeAllBranches;
            return this;
        }

        /**
         * @param includeAllBranches Whether the new repository should include all the branches from the template repository (defaults to false, which includes only the default branch from the template).
         * 
         * @return builder
         * 
         */
        public Builder includeAllBranches(Boolean includeAllBranches) {
            return includeAllBranches(Output.of(includeAllBranches));
        }

        /**
         * @param owner The GitHub organization or user the template repository is owned by.
         * 
         * @return builder
         * 
         */
        public Builder owner(Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The GitHub organization or user the template repository is owned by.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param repository The name of the template repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The name of the template repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public RepositoryTemplateArgs build() {
            if ($.owner == null) {
                throw new MissingRequiredPropertyException("RepositoryTemplateArgs", "owner");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("RepositoryTemplateArgs", "repository");
            }
            return $;
        }
    }

}