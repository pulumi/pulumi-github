// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrganizationProjectArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationProjectArgs Empty = new OrganizationProjectArgs();

    /**
     * The body of the project.
     * 
     */
    @Import(name="body")
    private @Nullable Output<String> body;

    /**
     * @return The body of the project.
     * 
     */
    public Optional<Output<String>> body() {
        return Optional.ofNullable(this.body);
    }

    /**
     * The name of the project.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the project.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private OrganizationProjectArgs() {}

    private OrganizationProjectArgs(OrganizationProjectArgs $) {
        this.body = $.body;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationProjectArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationProjectArgs $;

        public Builder() {
            $ = new OrganizationProjectArgs();
        }

        public Builder(OrganizationProjectArgs defaults) {
            $ = new OrganizationProjectArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param body The body of the project.
         * 
         * @return builder
         * 
         */
        public Builder body(@Nullable Output<String> body) {
            $.body = body;
            return this;
        }

        /**
         * @param body The body of the project.
         * 
         * @return builder
         * 
         */
        public Builder body(String body) {
            return body(Output.of(body));
        }

        /**
         * @param name The name of the project.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the project.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public OrganizationProjectArgs build() {
            return $;
        }
    }

}
