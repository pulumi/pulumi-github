// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ActionsEnvironmentVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActionsEnvironmentVariableArgs Empty = new ActionsEnvironmentVariableArgs();

    /**
     * Name of the environment.
     * 
     */
    @Import(name="environment", required=true)
    private Output<String> environment;

    /**
     * @return Name of the environment.
     * 
     */
    public Output<String> environment() {
        return this.environment;
    }

    /**
     * Name of the repository.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return Name of the repository.
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
     * Name of the variable.
     * 
     */
    @Import(name="variableName", required=true)
    private Output<String> variableName;

    /**
     * @return Name of the variable.
     * 
     */
    public Output<String> variableName() {
        return this.variableName;
    }

    private ActionsEnvironmentVariableArgs() {}

    private ActionsEnvironmentVariableArgs(ActionsEnvironmentVariableArgs $) {
        this.environment = $.environment;
        this.repository = $.repository;
        this.value = $.value;
        this.variableName = $.variableName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActionsEnvironmentVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActionsEnvironmentVariableArgs $;

        public Builder() {
            $ = new ActionsEnvironmentVariableArgs();
        }

        public Builder(ActionsEnvironmentVariableArgs defaults) {
            $ = new ActionsEnvironmentVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param environment Name of the environment.
         * 
         * @return builder
         * 
         */
        public Builder environment(Output<String> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment Name of the environment.
         * 
         * @return builder
         * 
         */
        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param repository Name of the repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository Name of the repository.
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
         * @param variableName Name of the variable.
         * 
         * @return builder
         * 
         */
        public Builder variableName(Output<String> variableName) {
            $.variableName = variableName;
            return this;
        }

        /**
         * @param variableName Name of the variable.
         * 
         * @return builder
         * 
         */
        public Builder variableName(String variableName) {
            return variableName(Output.of(variableName));
        }

        public ActionsEnvironmentVariableArgs build() {
            if ($.environment == null) {
                throw new MissingRequiredPropertyException("ActionsEnvironmentVariableArgs", "environment");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("ActionsEnvironmentVariableArgs", "repository");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("ActionsEnvironmentVariableArgs", "value");
            }
            if ($.variableName == null) {
                throw new MissingRequiredPropertyException("ActionsEnvironmentVariableArgs", "variableName");
            }
            return $;
        }
    }

}
