// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDependabotPublicKeyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDependabotPublicKeyPlainArgs Empty = new GetDependabotPublicKeyPlainArgs();

    /**
     * Name of the repository to get public key from.
     * 
     */
    @Import(name="repository", required=true)
    private String repository;

    /**
     * @return Name of the repository to get public key from.
     * 
     */
    public String repository() {
        return this.repository;
    }

    private GetDependabotPublicKeyPlainArgs() {}

    private GetDependabotPublicKeyPlainArgs(GetDependabotPublicKeyPlainArgs $) {
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDependabotPublicKeyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDependabotPublicKeyPlainArgs $;

        public Builder() {
            $ = new GetDependabotPublicKeyPlainArgs();
        }

        public Builder(GetDependabotPublicKeyPlainArgs defaults) {
            $ = new GetDependabotPublicKeyPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param repository Name of the repository to get public key from.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            $.repository = repository;
            return this;
        }

        public GetDependabotPublicKeyPlainArgs build() {
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("GetDependabotPublicKeyPlainArgs", "repository");
            }
            return $;
        }
    }

}
