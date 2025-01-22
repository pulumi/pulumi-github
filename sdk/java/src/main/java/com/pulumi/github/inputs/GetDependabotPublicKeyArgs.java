// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDependabotPublicKeyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDependabotPublicKeyArgs Empty = new GetDependabotPublicKeyArgs();

    /**
     * Name of the repository to get public key from.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return Name of the repository to get public key from.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private GetDependabotPublicKeyArgs() {}

    private GetDependabotPublicKeyArgs(GetDependabotPublicKeyArgs $) {
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDependabotPublicKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDependabotPublicKeyArgs $;

        public Builder() {
            $ = new GetDependabotPublicKeyArgs();
        }

        public Builder(GetDependabotPublicKeyArgs defaults) {
            $ = new GetDependabotPublicKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param repository Name of the repository to get public key from.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository Name of the repository to get public key from.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public GetDependabotPublicKeyArgs build() {
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("GetDependabotPublicKeyArgs", "repository");
            }
            return $;
        }
    }

}