// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetEnterprisePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEnterprisePlainArgs Empty = new GetEnterprisePlainArgs();

    /**
     * The URL slug identifying the enterprise.
     * 
     */
    @Import(name="slug", required=true)
    private String slug;

    /**
     * @return The URL slug identifying the enterprise.
     * 
     */
    public String slug() {
        return this.slug;
    }

    private GetEnterprisePlainArgs() {}

    private GetEnterprisePlainArgs(GetEnterprisePlainArgs $) {
        this.slug = $.slug;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEnterprisePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEnterprisePlainArgs $;

        public Builder() {
            $ = new GetEnterprisePlainArgs();
        }

        public Builder(GetEnterprisePlainArgs defaults) {
            $ = new GetEnterprisePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param slug The URL slug identifying the enterprise.
         * 
         * @return builder
         * 
         */
        public Builder slug(String slug) {
            $.slug = slug;
            return this;
        }

        public GetEnterprisePlainArgs build() {
            if ($.slug == null) {
                throw new MissingRequiredPropertyException("GetEnterprisePlainArgs", "slug");
            }
            return $;
        }
    }

}
