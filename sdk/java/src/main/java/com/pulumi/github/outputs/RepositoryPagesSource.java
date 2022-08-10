// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RepositoryPagesSource {
    /**
     * @return The repository branch used to publish the site&#39;s source files. (i.e. `main` or `gh-pages`.
     * 
     */
    private final String branch;
    /**
     * @return The repository directory from which the site publishes (Default: `/`).
     * 
     */
    private final @Nullable String path;

    @CustomType.Constructor
    private RepositoryPagesSource(
        @CustomType.Parameter("branch") String branch,
        @CustomType.Parameter("path") @Nullable String path) {
        this.branch = branch;
        this.path = path;
    }

    /**
     * @return The repository branch used to publish the site&#39;s source files. (i.e. `main` or `gh-pages`.
     * 
     */
    public String branch() {
        return this.branch;
    }
    /**
     * @return The repository directory from which the site publishes (Default: `/`).
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryPagesSource defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String branch;
        private @Nullable String path;

        public Builder() {
    	      // Empty
        }

        public Builder(RepositoryPagesSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branch = defaults.branch;
    	      this.path = defaults.path;
        }

        public Builder branch(String branch) {
            this.branch = Objects.requireNonNull(branch);
            return this;
        }
        public Builder path(@Nullable String path) {
            this.path = path;
            return this;
        }        public RepositoryPagesSource build() {
            return new RepositoryPagesSource(branch, path);
        }
    }
}