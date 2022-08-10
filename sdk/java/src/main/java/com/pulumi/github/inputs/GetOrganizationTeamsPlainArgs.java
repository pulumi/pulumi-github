// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOrganizationTeamsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOrganizationTeamsPlainArgs Empty = new GetOrganizationTeamsPlainArgs();

    /**
     * Only return teams that are at the organization&#39;s root, i.e. no nested teams. Defaults to `false`.
     * 
     */
    @Import(name="rootTeamsOnly")
    private @Nullable Boolean rootTeamsOnly;

    /**
     * @return Only return teams that are at the organization&#39;s root, i.e. no nested teams. Defaults to `false`.
     * 
     */
    public Optional<Boolean> rootTeamsOnly() {
        return Optional.ofNullable(this.rootTeamsOnly);
    }

    private GetOrganizationTeamsPlainArgs() {}

    private GetOrganizationTeamsPlainArgs(GetOrganizationTeamsPlainArgs $) {
        this.rootTeamsOnly = $.rootTeamsOnly;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOrganizationTeamsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOrganizationTeamsPlainArgs $;

        public Builder() {
            $ = new GetOrganizationTeamsPlainArgs();
        }

        public Builder(GetOrganizationTeamsPlainArgs defaults) {
            $ = new GetOrganizationTeamsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param rootTeamsOnly Only return teams that are at the organization&#39;s root, i.e. no nested teams. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder rootTeamsOnly(@Nullable Boolean rootTeamsOnly) {
            $.rootTeamsOnly = rootTeamsOnly;
            return this;
        }

        public GetOrganizationTeamsPlainArgs build() {
            return $;
        }
    }

}