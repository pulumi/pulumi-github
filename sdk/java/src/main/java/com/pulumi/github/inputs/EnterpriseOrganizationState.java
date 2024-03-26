// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EnterpriseOrganizationState extends com.pulumi.resources.ResourceArgs {

    public static final EnterpriseOrganizationState Empty = new EnterpriseOrganizationState();

    /**
     * List of organization owner usernames.
     * 
     */
    @Import(name="adminLogins")
    private @Nullable Output<List<String>> adminLogins;

    /**
     * @return List of organization owner usernames.
     * 
     */
    public Optional<Output<List<String>>> adminLogins() {
        return Optional.ofNullable(this.adminLogins);
    }

    /**
     * The billing email address.
     * 
     */
    @Import(name="billingEmail")
    private @Nullable Output<String> billingEmail;

    /**
     * @return The billing email address.
     * 
     */
    public Optional<Output<String>> billingEmail() {
        return Optional.ofNullable(this.billingEmail);
    }

    /**
     * The ID of the organization.
     * 
     */
    @Import(name="databaseId")
    private @Nullable Output<Integer> databaseId;

    /**
     * @return The ID of the organization.
     * 
     */
    public Optional<Output<Integer>> databaseId() {
        return Optional.ofNullable(this.databaseId);
    }

    /**
     * The description of the organization.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the organization.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The display name of the organization.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The display name of the organization.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The ID of the enterprise.
     * 
     */
    @Import(name="enterpriseId")
    private @Nullable Output<String> enterpriseId;

    /**
     * @return The ID of the enterprise.
     * 
     */
    public Optional<Output<String>> enterpriseId() {
        return Optional.ofNullable(this.enterpriseId);
    }

    /**
     * The name of the organization.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the organization.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private EnterpriseOrganizationState() {}

    private EnterpriseOrganizationState(EnterpriseOrganizationState $) {
        this.adminLogins = $.adminLogins;
        this.billingEmail = $.billingEmail;
        this.databaseId = $.databaseId;
        this.description = $.description;
        this.displayName = $.displayName;
        this.enterpriseId = $.enterpriseId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EnterpriseOrganizationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EnterpriseOrganizationState $;

        public Builder() {
            $ = new EnterpriseOrganizationState();
        }

        public Builder(EnterpriseOrganizationState defaults) {
            $ = new EnterpriseOrganizationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminLogins List of organization owner usernames.
         * 
         * @return builder
         * 
         */
        public Builder adminLogins(@Nullable Output<List<String>> adminLogins) {
            $.adminLogins = adminLogins;
            return this;
        }

        /**
         * @param adminLogins List of organization owner usernames.
         * 
         * @return builder
         * 
         */
        public Builder adminLogins(List<String> adminLogins) {
            return adminLogins(Output.of(adminLogins));
        }

        /**
         * @param adminLogins List of organization owner usernames.
         * 
         * @return builder
         * 
         */
        public Builder adminLogins(String... adminLogins) {
            return adminLogins(List.of(adminLogins));
        }

        /**
         * @param billingEmail The billing email address.
         * 
         * @return builder
         * 
         */
        public Builder billingEmail(@Nullable Output<String> billingEmail) {
            $.billingEmail = billingEmail;
            return this;
        }

        /**
         * @param billingEmail The billing email address.
         * 
         * @return builder
         * 
         */
        public Builder billingEmail(String billingEmail) {
            return billingEmail(Output.of(billingEmail));
        }

        /**
         * @param databaseId The ID of the organization.
         * 
         * @return builder
         * 
         */
        public Builder databaseId(@Nullable Output<Integer> databaseId) {
            $.databaseId = databaseId;
            return this;
        }

        /**
         * @param databaseId The ID of the organization.
         * 
         * @return builder
         * 
         */
        public Builder databaseId(Integer databaseId) {
            return databaseId(Output.of(databaseId));
        }

        /**
         * @param description The description of the organization.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the organization.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName The display name of the organization.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name of the organization.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param enterpriseId The ID of the enterprise.
         * 
         * @return builder
         * 
         */
        public Builder enterpriseId(@Nullable Output<String> enterpriseId) {
            $.enterpriseId = enterpriseId;
            return this;
        }

        /**
         * @param enterpriseId The ID of the enterprise.
         * 
         * @return builder
         * 
         */
        public Builder enterpriseId(String enterpriseId) {
            return enterpriseId(Output.of(enterpriseId));
        }

        /**
         * @param name The name of the organization.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the organization.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public EnterpriseOrganizationState build() {
            return $;
        }
    }

}
