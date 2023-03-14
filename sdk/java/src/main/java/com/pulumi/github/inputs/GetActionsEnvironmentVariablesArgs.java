// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetActionsEnvironmentVariablesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetActionsEnvironmentVariablesArgs Empty = new GetActionsEnvironmentVariablesArgs();

    @Import(name="environment", required=true)
    private Output<String> environment;

    public Output<String> environment() {
        return this.environment;
    }

    @Import(name="fullName")
    private @Nullable Output<String> fullName;

    public Optional<Output<String>> fullName() {
        return Optional.ofNullable(this.fullName);
    }

    /**
     * Name of the variable
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the variable
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetActionsEnvironmentVariablesArgs() {}

    private GetActionsEnvironmentVariablesArgs(GetActionsEnvironmentVariablesArgs $) {
        this.environment = $.environment;
        this.fullName = $.fullName;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetActionsEnvironmentVariablesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetActionsEnvironmentVariablesArgs $;

        public Builder() {
            $ = new GetActionsEnvironmentVariablesArgs();
        }

        public Builder(GetActionsEnvironmentVariablesArgs defaults) {
            $ = new GetActionsEnvironmentVariablesArgs(Objects.requireNonNull(defaults));
        }

        public Builder environment(Output<String> environment) {
            $.environment = environment;
            return this;
        }

        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        public Builder fullName(@Nullable Output<String> fullName) {
            $.fullName = fullName;
            return this;
        }

        public Builder fullName(String fullName) {
            return fullName(Output.of(fullName));
        }

        /**
         * @param name Name of the variable
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the variable
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetActionsEnvironmentVariablesArgs build() {
            $.environment = Objects.requireNonNull($.environment, "expected parameter 'environment' to be non-null");
            return $;
        }
    }

}
