// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetGithubAppArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGithubAppArgs Empty = new GetGithubAppArgs();

    /**
     * The URL-friendly name of your GitHub App.
     * 
     */
    @Import(name="slug", required=true)
    private Output<String> slug;

    /**
     * @return The URL-friendly name of your GitHub App.
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }

    private GetGithubAppArgs() {}

    private GetGithubAppArgs(GetGithubAppArgs $) {
        this.slug = $.slug;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGithubAppArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGithubAppArgs $;

        public Builder() {
            $ = new GetGithubAppArgs();
        }

        public Builder(GetGithubAppArgs defaults) {
            $ = new GetGithubAppArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param slug The URL-friendly name of your GitHub App.
         * 
         * @return builder
         * 
         */
        public Builder slug(Output<String> slug) {
            $.slug = slug;
            return this;
        }

        /**
         * @param slug The URL-friendly name of your GitHub App.
         * 
         * @return builder
         * 
         */
        public Builder slug(String slug) {
            return slug(Output.of(slug));
        }

        public GetGithubAppArgs build() {
            if ($.slug == null) {
                throw new MissingRequiredPropertyException("GetGithubAppArgs", "slug");
            }
            return $;
        }
    }

}
