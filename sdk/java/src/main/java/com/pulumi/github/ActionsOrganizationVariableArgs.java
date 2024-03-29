// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ActionsOrganizationVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActionsOrganizationVariableArgs Empty = new ActionsOrganizationVariableArgs();

    /**
     * An array of repository ids that can access the organization variable.
     * 
     */
    @Import(name="selectedRepositoryIds")
    private @Nullable Output<List<Integer>> selectedRepositoryIds;

    /**
     * @return An array of repository ids that can access the organization variable.
     * 
     */
    public Optional<Output<List<Integer>>> selectedRepositoryIds() {
        return Optional.ofNullable(this.selectedRepositoryIds);
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
     * Name of the variable
     * 
     */
    @Import(name="variableName", required=true)
    private Output<String> variableName;

    /**
     * @return Name of the variable
     * 
     */
    public Output<String> variableName() {
        return this.variableName;
    }

    /**
     * Configures the access that repositories have to the organization variable.
     * Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
     * 
     */
    @Import(name="visibility", required=true)
    private Output<String> visibility;

    /**
     * @return Configures the access that repositories have to the organization variable.
     * Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }

    private ActionsOrganizationVariableArgs() {}

    private ActionsOrganizationVariableArgs(ActionsOrganizationVariableArgs $) {
        this.selectedRepositoryIds = $.selectedRepositoryIds;
        this.value = $.value;
        this.variableName = $.variableName;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActionsOrganizationVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActionsOrganizationVariableArgs $;

        public Builder() {
            $ = new ActionsOrganizationVariableArgs();
        }

        public Builder(ActionsOrganizationVariableArgs defaults) {
            $ = new ActionsOrganizationVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param selectedRepositoryIds An array of repository ids that can access the organization variable.
         * 
         * @return builder
         * 
         */
        public Builder selectedRepositoryIds(@Nullable Output<List<Integer>> selectedRepositoryIds) {
            $.selectedRepositoryIds = selectedRepositoryIds;
            return this;
        }

        /**
         * @param selectedRepositoryIds An array of repository ids that can access the organization variable.
         * 
         * @return builder
         * 
         */
        public Builder selectedRepositoryIds(List<Integer> selectedRepositoryIds) {
            return selectedRepositoryIds(Output.of(selectedRepositoryIds));
        }

        /**
         * @param selectedRepositoryIds An array of repository ids that can access the organization variable.
         * 
         * @return builder
         * 
         */
        public Builder selectedRepositoryIds(Integer... selectedRepositoryIds) {
            return selectedRepositoryIds(List.of(selectedRepositoryIds));
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
         * @param variableName Name of the variable
         * 
         * @return builder
         * 
         */
        public Builder variableName(Output<String> variableName) {
            $.variableName = variableName;
            return this;
        }

        /**
         * @param variableName Name of the variable
         * 
         * @return builder
         * 
         */
        public Builder variableName(String variableName) {
            return variableName(Output.of(variableName));
        }

        /**
         * @param visibility Configures the access that repositories have to the organization variable.
         * Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
         * 
         * @return builder
         * 
         */
        public Builder visibility(Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        /**
         * @param visibility Configures the access that repositories have to the organization variable.
         * Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
         * 
         * @return builder
         * 
         */
        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        public ActionsOrganizationVariableArgs build() {
            if ($.value == null) {
                throw new MissingRequiredPropertyException("ActionsOrganizationVariableArgs", "value");
            }
            if ($.variableName == null) {
                throw new MissingRequiredPropertyException("ActionsOrganizationVariableArgs", "variableName");
            }
            if ($.visibility == null) {
                throw new MissingRequiredPropertyException("ActionsOrganizationVariableArgs", "visibility");
            }
            return $;
        }
    }

}
