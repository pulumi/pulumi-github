// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetUserExternalIdentityArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserExternalIdentityArgs Empty = new GetUserExternalIdentityArgs();

    /**
     * The username of the member to fetch external identity for.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the member to fetch external identity for.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private GetUserExternalIdentityArgs() {}

    private GetUserExternalIdentityArgs(GetUserExternalIdentityArgs $) {
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserExternalIdentityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserExternalIdentityArgs $;

        public Builder() {
            $ = new GetUserExternalIdentityArgs();
        }

        public Builder(GetUserExternalIdentityArgs defaults) {
            $ = new GetUserExternalIdentityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param username The username of the member to fetch external identity for.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the member to fetch external identity for.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public GetUserExternalIdentityArgs build() {
            if ($.username == null) {
                throw new MissingRequiredPropertyException("GetUserExternalIdentityArgs", "username");
            }
            return $;
        }
    }

}