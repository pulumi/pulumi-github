// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetOrganizationCustomRoleArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOrganizationCustomRoleArgs Empty = new GetOrganizationCustomRoleArgs();

    /**
     * The name of the custom role.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the custom role.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetOrganizationCustomRoleArgs() {}

    private GetOrganizationCustomRoleArgs(GetOrganizationCustomRoleArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOrganizationCustomRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOrganizationCustomRoleArgs $;

        public Builder() {
            $ = new GetOrganizationCustomRoleArgs();
        }

        public Builder(GetOrganizationCustomRoleArgs defaults) {
            $ = new GetOrganizationCustomRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the custom role.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the custom role.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetOrganizationCustomRoleArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetOrganizationCustomRoleArgs", "name");
            }
            return $;
        }
    }

}