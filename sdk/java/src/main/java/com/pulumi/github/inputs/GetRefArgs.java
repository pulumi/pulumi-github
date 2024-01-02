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


public final class GetRefArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRefArgs Empty = new GetRefArgs();

    /**
     * Owner of the repository.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return Owner of the repository.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * The repository ref to look up. Must be formatted `heads/&lt;ref&gt;` for branches, and `tags/&lt;ref&gt;` for tags.
     * 
     */
    @Import(name="ref", required=true)
    private Output<String> ref;

    /**
     * @return The repository ref to look up. Must be formatted `heads/&lt;ref&gt;` for branches, and `tags/&lt;ref&gt;` for tags.
     * 
     */
    public Output<String> ref() {
        return this.ref;
    }

    /**
     * The GitHub repository name.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The GitHub repository name.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private GetRefArgs() {}

    private GetRefArgs(GetRefArgs $) {
        this.owner = $.owner;
        this.ref = $.ref;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRefArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRefArgs $;

        public Builder() {
            $ = new GetRefArgs();
        }

        public Builder(GetRefArgs defaults) {
            $ = new GetRefArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param owner Owner of the repository.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner Owner of the repository.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param ref The repository ref to look up. Must be formatted `heads/&lt;ref&gt;` for branches, and `tags/&lt;ref&gt;` for tags.
         * 
         * @return builder
         * 
         */
        public Builder ref(Output<String> ref) {
            $.ref = ref;
            return this;
        }

        /**
         * @param ref The repository ref to look up. Must be formatted `heads/&lt;ref&gt;` for branches, and `tags/&lt;ref&gt;` for tags.
         * 
         * @return builder
         * 
         */
        public Builder ref(String ref) {
            return ref(Output.of(ref));
        }

        /**
         * @param repository The GitHub repository name.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The GitHub repository name.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public GetRefArgs build() {
            if ($.ref == null) {
                throw new MissingRequiredPropertyException("GetRefArgs", "ref");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("GetRefArgs", "repository");
            }
            return $;
        }
    }

}
