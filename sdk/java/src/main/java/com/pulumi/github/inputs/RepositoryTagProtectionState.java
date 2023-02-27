// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryTagProtectionState extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryTagProtectionState Empty = new RepositoryTagProtectionState();

    /**
     * The pattern of the tag to protect.
     * 
     */
    @Import(name="pattern")
    private @Nullable Output<String> pattern;

    /**
     * @return The pattern of the tag to protect.
     * 
     */
    public Optional<Output<String>> pattern() {
        return Optional.ofNullable(this.pattern);
    }

    /**
     * Name of the repository to add the tag protection to.
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return Name of the repository to add the tag protection to.
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * The ID of the tag protection.
     * 
     */
    @Import(name="tagProtectionId")
    private @Nullable Output<Integer> tagProtectionId;

    /**
     * @return The ID of the tag protection.
     * 
     */
    public Optional<Output<Integer>> tagProtectionId() {
        return Optional.ofNullable(this.tagProtectionId);
    }

    private RepositoryTagProtectionState() {}

    private RepositoryTagProtectionState(RepositoryTagProtectionState $) {
        this.pattern = $.pattern;
        this.repository = $.repository;
        this.tagProtectionId = $.tagProtectionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryTagProtectionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryTagProtectionState $;

        public Builder() {
            $ = new RepositoryTagProtectionState();
        }

        public Builder(RepositoryTagProtectionState defaults) {
            $ = new RepositoryTagProtectionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param pattern The pattern of the tag to protect.
         * 
         * @return builder
         * 
         */
        public Builder pattern(@Nullable Output<String> pattern) {
            $.pattern = pattern;
            return this;
        }

        /**
         * @param pattern The pattern of the tag to protect.
         * 
         * @return builder
         * 
         */
        public Builder pattern(String pattern) {
            return pattern(Output.of(pattern));
        }

        /**
         * @param repository Name of the repository to add the tag protection to.
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository Name of the repository to add the tag protection to.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param tagProtectionId The ID of the tag protection.
         * 
         * @return builder
         * 
         */
        public Builder tagProtectionId(@Nullable Output<Integer> tagProtectionId) {
            $.tagProtectionId = tagProtectionId;
            return this;
        }

        /**
         * @param tagProtectionId The ID of the tag protection.
         * 
         * @return builder
         * 
         */
        public Builder tagProtectionId(Integer tagProtectionId) {
            return tagProtectionId(Output.of(tagProtectionId));
        }

        public RepositoryTagProtectionState build() {
            return $;
        }
    }

}
