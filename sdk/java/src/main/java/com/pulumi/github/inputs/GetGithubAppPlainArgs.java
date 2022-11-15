// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetGithubAppPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGithubAppPlainArgs Empty = new GetGithubAppPlainArgs();

    @Import(name="slug", required=true)
    private String slug;

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

        public Builder slug(String slug) {
            $.slug = slug;
            return this;
        }

        public GetGithubAppPlainArgs build() {
            $.slug = Objects.requireNonNull($.slug, "expected parameter 'slug' to be non-null");
            return $;
        }
    }

}
