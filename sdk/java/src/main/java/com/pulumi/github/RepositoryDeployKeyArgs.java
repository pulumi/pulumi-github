// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryDeployKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryDeployKeyArgs Empty = new RepositoryDeployKeyArgs();

    /**
     * A SSH key.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return A SSH key.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * A boolean qualifying the key to be either read only or read/write.
     * 
     */
    @Import(name="readOnly")
    private @Nullable Output<Boolean> readOnly;

    /**
     * @return A boolean qualifying the key to be either read only or read/write.
     * 
     */
    public Optional<Output<Boolean>> readOnly() {
        return Optional.ofNullable(this.readOnly);
    }

    /**
     * Name of the GitHub repository.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return Name of the GitHub repository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     * A title.
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return A title.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    private RepositoryDeployKeyArgs() {}

    private RepositoryDeployKeyArgs(RepositoryDeployKeyArgs $) {
        this.key = $.key;
        this.readOnly = $.readOnly;
        this.repository = $.repository;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryDeployKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryDeployKeyArgs $;

        public Builder() {
            $ = new RepositoryDeployKeyArgs();
        }

        public Builder(RepositoryDeployKeyArgs defaults) {
            $ = new RepositoryDeployKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key A SSH key.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key A SSH key.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param readOnly A boolean qualifying the key to be either read only or read/write.
         * 
         * @return builder
         * 
         */
        public Builder readOnly(@Nullable Output<Boolean> readOnly) {
            $.readOnly = readOnly;
            return this;
        }

        /**
         * @param readOnly A boolean qualifying the key to be either read only or read/write.
         * 
         * @return builder
         * 
         */
        public Builder readOnly(Boolean readOnly) {
            return readOnly(Output.of(readOnly));
        }

        /**
         * @param repository Name of the GitHub repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository Name of the GitHub repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param title A title.
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title A title.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public RepositoryDeployKeyArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            $.title = Objects.requireNonNull($.title, "expected parameter 'title' to be non-null");
            return $;
        }
    }

}
