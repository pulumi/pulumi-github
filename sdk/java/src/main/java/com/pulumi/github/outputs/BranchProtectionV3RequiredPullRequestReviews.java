// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BranchProtectionV3RequiredPullRequestReviews {
    private final @Nullable Boolean dismissStaleReviews;
    private final @Nullable List<String> dismissalTeams;
    private final @Nullable List<String> dismissalUsers;
    /**
     * @deprecated
     * Use enforce_admins instead
     * 
     */
    @Deprecated /* Use enforce_admins instead */
    private final @Nullable Boolean includeAdmins;
    private final @Nullable Boolean requireCodeOwnerReviews;
    private final @Nullable Integer requiredApprovingReviewCount;

    @CustomType.Constructor
    private BranchProtectionV3RequiredPullRequestReviews(
        @CustomType.Parameter("dismissStaleReviews") @Nullable Boolean dismissStaleReviews,
        @CustomType.Parameter("dismissalTeams") @Nullable List<String> dismissalTeams,
        @CustomType.Parameter("dismissalUsers") @Nullable List<String> dismissalUsers,
        @CustomType.Parameter("includeAdmins") @Nullable Boolean includeAdmins,
        @CustomType.Parameter("requireCodeOwnerReviews") @Nullable Boolean requireCodeOwnerReviews,
        @CustomType.Parameter("requiredApprovingReviewCount") @Nullable Integer requiredApprovingReviewCount) {
        this.dismissStaleReviews = dismissStaleReviews;
        this.dismissalTeams = dismissalTeams;
        this.dismissalUsers = dismissalUsers;
        this.includeAdmins = includeAdmins;
        this.requireCodeOwnerReviews = requireCodeOwnerReviews;
        this.requiredApprovingReviewCount = requiredApprovingReviewCount;
    }

    public Optional<Boolean> dismissStaleReviews() {
        return Optional.ofNullable(this.dismissStaleReviews);
    }
    public List<String> dismissalTeams() {
        return this.dismissalTeams == null ? List.of() : this.dismissalTeams;
    }
    public List<String> dismissalUsers() {
        return this.dismissalUsers == null ? List.of() : this.dismissalUsers;
    }
    /**
     * @deprecated
     * Use enforce_admins instead
     * 
     */
    @Deprecated /* Use enforce_admins instead */
    public Optional<Boolean> includeAdmins() {
        return Optional.ofNullable(this.includeAdmins);
    }
    public Optional<Boolean> requireCodeOwnerReviews() {
        return Optional.ofNullable(this.requireCodeOwnerReviews);
    }
    public Optional<Integer> requiredApprovingReviewCount() {
        return Optional.ofNullable(this.requiredApprovingReviewCount);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchProtectionV3RequiredPullRequestReviews defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean dismissStaleReviews;
        private @Nullable List<String> dismissalTeams;
        private @Nullable List<String> dismissalUsers;
        private @Nullable Boolean includeAdmins;
        private @Nullable Boolean requireCodeOwnerReviews;
        private @Nullable Integer requiredApprovingReviewCount;

        public Builder() {
    	      // Empty
        }

        public Builder(BranchProtectionV3RequiredPullRequestReviews defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dismissStaleReviews = defaults.dismissStaleReviews;
    	      this.dismissalTeams = defaults.dismissalTeams;
    	      this.dismissalUsers = defaults.dismissalUsers;
    	      this.includeAdmins = defaults.includeAdmins;
    	      this.requireCodeOwnerReviews = defaults.requireCodeOwnerReviews;
    	      this.requiredApprovingReviewCount = defaults.requiredApprovingReviewCount;
        }

        public Builder dismissStaleReviews(@Nullable Boolean dismissStaleReviews) {
            this.dismissStaleReviews = dismissStaleReviews;
            return this;
        }
        public Builder dismissalTeams(@Nullable List<String> dismissalTeams) {
            this.dismissalTeams = dismissalTeams;
            return this;
        }
        public Builder dismissalTeams(String... dismissalTeams) {
            return dismissalTeams(List.of(dismissalTeams));
        }
        public Builder dismissalUsers(@Nullable List<String> dismissalUsers) {
            this.dismissalUsers = dismissalUsers;
            return this;
        }
        public Builder dismissalUsers(String... dismissalUsers) {
            return dismissalUsers(List.of(dismissalUsers));
        }
        public Builder includeAdmins(@Nullable Boolean includeAdmins) {
            this.includeAdmins = includeAdmins;
            return this;
        }
        public Builder requireCodeOwnerReviews(@Nullable Boolean requireCodeOwnerReviews) {
            this.requireCodeOwnerReviews = requireCodeOwnerReviews;
            return this;
        }
        public Builder requiredApprovingReviewCount(@Nullable Integer requiredApprovingReviewCount) {
            this.requiredApprovingReviewCount = requiredApprovingReviewCount;
            return this;
        }        public BranchProtectionV3RequiredPullRequestReviews build() {
            return new BranchProtectionV3RequiredPullRequestReviews(dismissStaleReviews, dismissalTeams, dismissalUsers, includeAdmins, requireCodeOwnerReviews, requiredApprovingReviewCount);
        }
    }
}
