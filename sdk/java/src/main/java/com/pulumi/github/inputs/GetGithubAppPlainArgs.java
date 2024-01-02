// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetGithubAppPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGithubAppPlainArgs Empty = new GetGithubAppPlainArgs();

    /**
     * The URL-friendly name of your GitHub App.
     * 
     */
    @Import(name="slug", required=true)
    private String slug;

    /**
     * @return The URL-friendly name of your GitHub App.
     * 
     */
    public String slug() {
        return this.slug;
    }

    private GetGithubAppPlainArgs() {}

    private GetGithubAppPlainArgs(GetGithubAppPlainArgs $) {
        this.slug = $.slug;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGithubAppPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGithubAppPlainArgs $;

        public Builder() {
            $ = new GetGithubAppPlainArgs();
        }

        public Builder(GetGithubAppPlainArgs defaults) {
            $ = new GetGithubAppPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param slug The URL-friendly name of your GitHub App.
         * 
         * @return builder
         * 
         */
        public Builder slug(String slug) {
            $.slug = slug;
            return this;
        }

        public GetGithubAppPlainArgs build() {
            if ($.slug == null) {
                throw new MissingRequiredPropertyException("GetGithubAppPlainArgs", "slug");
            }
            return $;
        }
    }

}
