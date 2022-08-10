// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetMembershipArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMembershipArgs Empty = new GetMembershipArgs();

    /**
     * The organization to check for the above username.
     * 
     */
    @Import(name="organization")
    private @Nullable Output<String> organization;

    /**
     * @return The organization to check for the above username.
     * 
     */
    public Optional<Output<String>> organization() {
        return Optional.ofNullable(this.organization);
    }

    /**
     * The username to lookup in the organization.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username to lookup in the organization.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private GetMembershipArgs() {}

    private GetMembershipArgs(GetMembershipArgs $) {
        this.organization = $.organization;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMembershipArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMembershipArgs $;

        public Builder() {
            $ = new GetMembershipArgs();
        }

        public Builder(GetMembershipArgs defaults) {
            $ = new GetMembershipArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param organization The organization to check for the above username.
         * 
         * @return builder
         * 
         */
        public Builder organization(@Nullable Output<String> organization) {
            $.organization = organization;
            return this;
        }

        /**
         * @param organization The organization to check for the above username.
         * 
         * @return builder
         * 
         */
        public Builder organization(String organization) {
            return organization(Output.of(organization));
        }

        /**
         * @param username The username to lookup in the organization.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username to lookup in the organization.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public GetMembershipArgs build() {
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}