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


public final class RepositoryAutolinkReferenceArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryAutolinkReferenceArgs Empty = new RepositoryAutolinkReferenceArgs();

    /**
     * Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric
     * characters.
     * 
     */
    @Import(name="isAlphanumeric")
    private @Nullable Output<Boolean> isAlphanumeric;

    /**
     * @return Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric
     * characters.
     * 
     */
    public Optional<Output<Boolean>> isAlphanumeric() {
        return Optional.ofNullable(this.isAlphanumeric);
    }

    /**
     * This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
     * 
     */
    @Import(name="keyPrefix", required=true)
    private Output<String> keyPrefix;

    /**
     * @return This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
     * 
     */
    public Output<String> keyPrefix() {
        return this.keyPrefix;
    }

    /**
     * The repository name
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The repository name
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     * The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
     * 
     */
    @Import(name="targetUrlTemplate", required=true)
    private Output<String> targetUrlTemplate;

    /**
     * @return The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
     * 
     */
    public Output<String> targetUrlTemplate() {
        return this.targetUrlTemplate;
    }

    private RepositoryAutolinkReferenceArgs() {}

    private RepositoryAutolinkReferenceArgs(RepositoryAutolinkReferenceArgs $) {
        this.isAlphanumeric = $.isAlphanumeric;
        this.keyPrefix = $.keyPrefix;
        this.repository = $.repository;
        this.targetUrlTemplate = $.targetUrlTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryAutolinkReferenceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryAutolinkReferenceArgs $;

        public Builder() {
            $ = new RepositoryAutolinkReferenceArgs();
        }

        public Builder(RepositoryAutolinkReferenceArgs defaults) {
            $ = new RepositoryAutolinkReferenceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param isAlphanumeric Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric
         * characters.
         * 
         * @return builder
         * 
         */
        public Builder isAlphanumeric(@Nullable Output<Boolean> isAlphanumeric) {
            $.isAlphanumeric = isAlphanumeric;
            return this;
        }

        /**
         * @param isAlphanumeric Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric
         * characters.
         * 
         * @return builder
         * 
         */
        public Builder isAlphanumeric(Boolean isAlphanumeric) {
            return isAlphanumeric(Output.of(isAlphanumeric));
        }

        /**
         * @param keyPrefix This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
         * 
         * @return builder
         * 
         */
        public Builder keyPrefix(Output<String> keyPrefix) {
            $.keyPrefix = keyPrefix;
            return this;
        }

        /**
         * @param keyPrefix This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
         * 
         * @return builder
         * 
         */
        public Builder keyPrefix(String keyPrefix) {
            return keyPrefix(Output.of(keyPrefix));
        }

        /**
         * @param repository The repository name
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The repository name
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param targetUrlTemplate The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
         * 
         * @return builder
         * 
         */
        public Builder targetUrlTemplate(Output<String> targetUrlTemplate) {
            $.targetUrlTemplate = targetUrlTemplate;
            return this;
        }

        /**
         * @param targetUrlTemplate The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
         * 
         * @return builder
         * 
         */
        public Builder targetUrlTemplate(String targetUrlTemplate) {
            return targetUrlTemplate(Output.of(targetUrlTemplate));
        }

        public RepositoryAutolinkReferenceArgs build() {
            $.keyPrefix = Objects.requireNonNull($.keyPrefix, "expected parameter 'keyPrefix' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            $.targetUrlTemplate = Objects.requireNonNull($.targetUrlTemplate, "expected parameter 'targetUrlTemplate' to be non-null");
            return $;
        }
    }

}
