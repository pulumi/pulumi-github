// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class OrganizationRulesetConditionsRefNameArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationRulesetConditionsRefNameArgs Empty = new OrganizationRulesetConditionsRefNameArgs();

    /**
     * Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
     * 
     */
    @Import(name="excludes", required=true)
    private Output<List<String>> excludes;

    /**
     * @return Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
     * 
     */
    public Output<List<String>> excludes() {
        return this.excludes;
    }

    /**
     * Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
     * 
     */
    @Import(name="includes", required=true)
    private Output<List<String>> includes;

    /**
     * @return Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
     * 
     */
    public Output<List<String>> includes() {
        return this.includes;
    }

    private OrganizationRulesetConditionsRefNameArgs() {}

    private OrganizationRulesetConditionsRefNameArgs(OrganizationRulesetConditionsRefNameArgs $) {
        this.excludes = $.excludes;
        this.includes = $.includes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationRulesetConditionsRefNameArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationRulesetConditionsRefNameArgs $;

        public Builder() {
            $ = new OrganizationRulesetConditionsRefNameArgs();
        }

        public Builder(OrganizationRulesetConditionsRefNameArgs defaults) {
            $ = new OrganizationRulesetConditionsRefNameArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param excludes Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
         * 
         * @return builder
         * 
         */
        public Builder excludes(Output<List<String>> excludes) {
            $.excludes = excludes;
            return this;
        }

        /**
         * @param excludes Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
         * 
         * @return builder
         * 
         */
        public Builder excludes(List<String> excludes) {
            return excludes(Output.of(excludes));
        }

        /**
         * @param excludes Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
         * 
         * @return builder
         * 
         */
        public Builder excludes(String... excludes) {
            return excludes(List.of(excludes));
        }

        /**
         * @param includes Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
         * 
         * @return builder
         * 
         */
        public Builder includes(Output<List<String>> includes) {
            $.includes = includes;
            return this;
        }

        /**
         * @param includes Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
         * 
         * @return builder
         * 
         */
        public Builder includes(List<String> includes) {
            return includes(Output.of(includes));
        }

        /**
         * @param includes Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
         * 
         * @return builder
         * 
         */
        public Builder includes(String... includes) {
            return includes(List.of(includes));
        }

        public OrganizationRulesetConditionsRefNameArgs build() {
            if ($.excludes == null) {
                throw new MissingRequiredPropertyException("OrganizationRulesetConditionsRefNameArgs", "excludes");
            }
            if ($.includes == null) {
                throw new MissingRequiredPropertyException("OrganizationRulesetConditionsRefNameArgs", "includes");
            }
            return $;
        }
    }

}
