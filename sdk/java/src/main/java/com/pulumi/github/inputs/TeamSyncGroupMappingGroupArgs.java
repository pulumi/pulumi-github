// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class TeamSyncGroupMappingGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeamSyncGroupMappingGroupArgs Empty = new TeamSyncGroupMappingGroupArgs();

    /**
     * The description of the IdP group.
     * 
     */
    @Import(name="groupDescription", required=true)
    private Output<String> groupDescription;

    /**
     * @return The description of the IdP group.
     * 
     */
    public Output<String> groupDescription() {
        return this.groupDescription;
    }

    /**
     * The ID of the IdP group.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<String> groupId;

    /**
     * @return The ID of the IdP group.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }

    /**
     * The name of the IdP group.
     * 
     */
    @Import(name="groupName", required=true)
    private Output<String> groupName;

    /**
     * @return The name of the IdP group.
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }

    private TeamSyncGroupMappingGroupArgs() {}

    private TeamSyncGroupMappingGroupArgs(TeamSyncGroupMappingGroupArgs $) {
        this.groupDescription = $.groupDescription;
        this.groupId = $.groupId;
        this.groupName = $.groupName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamSyncGroupMappingGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamSyncGroupMappingGroupArgs $;

        public Builder() {
            $ = new TeamSyncGroupMappingGroupArgs();
        }

        public Builder(TeamSyncGroupMappingGroupArgs defaults) {
            $ = new TeamSyncGroupMappingGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupDescription The description of the IdP group.
         * 
         * @return builder
         * 
         */
        public Builder groupDescription(Output<String> groupDescription) {
            $.groupDescription = groupDescription;
            return this;
        }

        /**
         * @param groupDescription The description of the IdP group.
         * 
         * @return builder
         * 
         */
        public Builder groupDescription(String groupDescription) {
            return groupDescription(Output.of(groupDescription));
        }

        /**
         * @param groupId The ID of the IdP group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the IdP group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param groupName The name of the IdP group.
         * 
         * @return builder
         * 
         */
        public Builder groupName(Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        /**
         * @param groupName The name of the IdP group.
         * 
         * @return builder
         * 
         */
        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        public TeamSyncGroupMappingGroupArgs build() {
            $.groupDescription = Objects.requireNonNull($.groupDescription, "expected parameter 'groupDescription' to be non-null");
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            $.groupName = Objects.requireNonNull($.groupName, "expected parameter 'groupName' to be non-null");
            return $;
        }
    }

}
