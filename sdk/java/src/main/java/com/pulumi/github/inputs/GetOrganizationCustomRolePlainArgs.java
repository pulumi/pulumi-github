// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetOrganizationCustomRolePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOrganizationCustomRolePlainArgs Empty = new GetOrganizationCustomRolePlainArgs();

    /**
     * The name of the custom role.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the custom role.
     * 
     */
    public String name() {
        return this.name;
    }

    private GetOrganizationCustomRolePlainArgs() {}

    private GetOrganizationCustomRolePlainArgs(GetOrganizationCustomRolePlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOrganizationCustomRolePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOrganizationCustomRolePlainArgs $;

        public Builder() {
            $ = new GetOrganizationCustomRolePlainArgs();
        }

        public Builder(GetOrganizationCustomRolePlainArgs defaults) {
            $ = new GetOrganizationCustomRolePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the custom role.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetOrganizationCustomRolePlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetOrganizationCustomRolePlainArgs", "name");
            }
            return $;
        }
    }

}
