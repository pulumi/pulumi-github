// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EnterpriseOrganizationArgs extends com.pulumi.resources.ResourceArgs {

    public static final EnterpriseOrganizationArgs Empty = new EnterpriseOrganizationArgs();

    @Import(name="adminLogins", required=true)
    private Output<List<String>> adminLogins;

    public Output<List<String>> adminLogins() {
        return this.adminLogins;
    }

    @Import(name="billingEmail", required=true)
    private Output<String> billingEmail;

    public Output<String> billingEmail() {
        return this.billingEmail;
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="enterpriseId", required=true)
    private Output<String> enterpriseId;

    public Output<String> enterpriseId() {
        return this.enterpriseId;
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private EnterpriseOrganizationArgs() {}

    private EnterpriseOrganizationArgs(EnterpriseOrganizationArgs $) {
        this.adminLogins = $.adminLogins;
        this.billingEmail = $.billingEmail;
        this.description = $.description;
        this.enterpriseId = $.enterpriseId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EnterpriseOrganizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EnterpriseOrganizationArgs $;

        public Builder() {
            $ = new EnterpriseOrganizationArgs();
        }

        public Builder(EnterpriseOrganizationArgs defaults) {
            $ = new EnterpriseOrganizationArgs(Objects.requireNonNull(defaults));
        }

        public Builder adminLogins(Output<List<String>> adminLogins) {
            $.adminLogins = adminLogins;
            return this;
        }

        public Builder adminLogins(List<String> adminLogins) {
            return adminLogins(Output.of(adminLogins));
        }

        public Builder adminLogins(String... adminLogins) {
            return adminLogins(List.of(adminLogins));
        }

        public Builder billingEmail(Output<String> billingEmail) {
            $.billingEmail = billingEmail;
            return this;
        }

        public Builder billingEmail(String billingEmail) {
            return billingEmail(Output.of(billingEmail));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder enterpriseId(Output<String> enterpriseId) {
            $.enterpriseId = enterpriseId;
            return this;
        }

        public Builder enterpriseId(String enterpriseId) {
            return enterpriseId(Output.of(enterpriseId));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public EnterpriseOrganizationArgs build() {
            $.adminLogins = Objects.requireNonNull($.adminLogins, "expected parameter 'adminLogins' to be non-null");
            $.billingEmail = Objects.requireNonNull($.billingEmail, "expected parameter 'billingEmail' to be non-null");
            $.enterpriseId = Objects.requireNonNull($.enterpriseId, "expected parameter 'enterpriseId' to be non-null");
            return $;
        }
    }

}
