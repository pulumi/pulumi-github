// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetRestApiPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRestApiPlainArgs Empty = new GetRestApiPlainArgs();

    /**
     * REST API endpoint to send the GET request to.
     * 
     */
    @Import(name="endpoint", required=true)
    private String endpoint;

    /**
     * @return REST API endpoint to send the GET request to.
     * 
     */
    public String endpoint() {
        return this.endpoint;
    }

    private GetRestApiPlainArgs() {}

    private GetRestApiPlainArgs(GetRestApiPlainArgs $) {
        this.endpoint = $.endpoint;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRestApiPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRestApiPlainArgs $;

        public Builder() {
            $ = new GetRestApiPlainArgs();
        }

        public Builder(GetRestApiPlainArgs defaults) {
            $ = new GetRestApiPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param endpoint REST API endpoint to send the GET request to.
         * 
         * @return builder
         * 
         */
        public Builder endpoint(String endpoint) {
            $.endpoint = endpoint;
            return this;
        }

        public GetRestApiPlainArgs build() {
            if ($.endpoint == null) {
                throw new MissingRequiredPropertyException("GetRestApiPlainArgs", "endpoint");
            }
            return $;
        }
    }

}
