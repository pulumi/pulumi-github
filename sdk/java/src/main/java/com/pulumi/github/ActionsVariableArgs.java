// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ActionsVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActionsVariableArgs Empty = new ActionsVariableArgs();

    /**
     * Name of the repository
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return Name of the repository
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     * Value of the variable
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Value of the variable
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     * Name of the variable
     * 
     */
    @Import(name="variableName", required=true)
    private Output<String> variableName;

    /**
     * @return Name of the variable
     * 
     */
    public Output<String> variableName() {
        return this.variableName;
    }

    private ActionsVariableArgs() {}

    private ActionsVariableArgs(ActionsVariableArgs $) {
        this.repository = $.repository;
        this.value = $.value;
        this.variableName = $.variableName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActionsVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActionsVariableArgs $;

        public Builder() {
            $ = new ActionsVariableArgs();
        }

        public Builder(ActionsVariableArgs defaults) {
            $ = new ActionsVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param repository Name of the repository
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository Name of the repository
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param value Value of the variable
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Value of the variable
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        /**
         * @param variableName Name of the variable
         * 
         * @return builder
         * 
         */
        public Builder variableName(Output<String> variableName) {
            $.variableName = variableName;
            return this;
        }

        /**
         * @param variableName Name of the variable
         * 
         * @return builder
         * 
         */
        public Builder variableName(String variableName) {
            return variableName(Output.of(variableName));
        }

        public ActionsVariableArgs build() {
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("ActionsVariableArgs", "repository");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("ActionsVariableArgs", "value");
            }
            if ($.variableName == null) {
                throw new MissingRequiredPropertyException("ActionsVariableArgs", "variableName");
            }
            return $;
        }
    }

}
