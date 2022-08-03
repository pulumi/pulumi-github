// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.TeamMembersMemberArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamMembersState extends com.pulumi.resources.ResourceArgs {

    public static final TeamMembersState Empty = new TeamMembersState();

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * List of team members. See Members below for details.
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<TeamMembersMemberArgs>> members;

    /**
     * @return List of team members. See Members below for details.
     * 
     */
    public Optional<Output<List<TeamMembersMemberArgs>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The GitHub team id
     * 
     */
    @Import(name="teamId")
    private @Nullable Output<String> teamId;

    /**
     * @return The GitHub team id
     * 
     */
    public Optional<Output<String>> teamId() {
        return Optional.ofNullable(this.teamId);
    }

    private TeamMembersState() {}

    private TeamMembersState(TeamMembersState $) {
        this.etag = $.etag;
        this.members = $.members;
        this.teamId = $.teamId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamMembersState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamMembersState $;

        public Builder() {
            $ = new TeamMembersState();
        }

        public Builder(TeamMembersState defaults) {
            $ = new TeamMembersState(Objects.requireNonNull(defaults));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param members List of team members. See Members below for details.
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<TeamMembersMemberArgs>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members List of team members. See Members below for details.
         * 
         * @return builder
         * 
         */
        public Builder members(List<TeamMembersMemberArgs> members) {
            return members(Output.of(members));
        }

        /**
         * @param members List of team members. See Members below for details.
         * 
         * @return builder
         * 
         */
        public Builder members(TeamMembersMemberArgs... members) {
            return members(List.of(members));
        }

        /**
         * @param teamId The GitHub team id
         * 
         * @return builder
         * 
         */
        public Builder teamId(@Nullable Output<String> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId The GitHub team id
         * 
         * @return builder
         * 
         */
        public Builder teamId(String teamId) {
            return teamId(Output.of(teamId));
        }

        public TeamMembersState build() {
            return $;
        }
    }

}
