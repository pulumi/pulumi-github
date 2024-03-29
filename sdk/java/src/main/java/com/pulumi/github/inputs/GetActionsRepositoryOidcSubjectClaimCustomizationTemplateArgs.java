// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs Empty = new GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs();

    /**
     * Name of the repository to get the OpenID Connect subject claim customization template for.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the repository to get the OpenID Connect subject claim customization template for.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs() {}

    private GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs(GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs $;

        public Builder() {
            $ = new GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs();
        }

        public Builder(GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs defaults) {
            $ = new GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the repository to get the OpenID Connect subject claim customization template for.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the repository to get the OpenID Connect subject claim customization template for.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs", "name");
            }
            return $;
        }
    }

}
