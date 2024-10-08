// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCollaboratorsCollaborator {
    /**
     * @return The GitHub API URL for the collaborator&#39;s events.
     * 
     */
    private String eventsUrl;
    /**
     * @return The GitHub API URL for the collaborator&#39;s followers.
     * 
     */
    private String followersUrl;
    /**
     * @return The GitHub API URL for those following the collaborator.
     * 
     */
    private String followingUrl;
    /**
     * @return The GitHub API URL for the collaborator&#39;s gists.
     * 
     */
    private String gistsUrl;
    /**
     * @return The GitHub HTML URL for the collaborator.
     * 
     */
    private String htmlUrl;
    /**
     * @return The ID of the collaborator.
     * 
     */
    private Integer id;
    /**
     * @return The collaborator&#39;s login.
     * 
     */
    private String login;
    /**
     * @return The GitHub API URL for the collaborator&#39;s organizations.
     * 
     */
    private String organizationsUrl;
    /**
     * @return Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
     * 
     */
    private String permission;
    /**
     * @return The GitHub API URL for the collaborator&#39;s received events.
     * 
     */
    private String receivedEventsUrl;
    /**
     * @return The GitHub API URL for the collaborator&#39;s repositories.
     * 
     */
    private String reposUrl;
    /**
     * @return Whether the user is a GitHub admin.
     * 
     */
    private Boolean siteAdmin;
    /**
     * @return The GitHub API URL for the collaborator&#39;s starred repositories.
     * 
     */
    private String starredUrl;
    /**
     * @return The GitHub API URL for the collaborator&#39;s subscribed repositories.
     * 
     */
    private String subscriptionsUrl;
    /**
     * @return The type of the collaborator (ex. `user`).
     * 
     */
    private String type;
    /**
     * @return The GitHub API URL for the collaborator.
     * 
     */
    private String url;

    private GetCollaboratorsCollaborator() {}
    /**
     * @return The GitHub API URL for the collaborator&#39;s events.
     * 
     */
    public String eventsUrl() {
        return this.eventsUrl;
    }
    /**
     * @return The GitHub API URL for the collaborator&#39;s followers.
     * 
     */
    public String followersUrl() {
        return this.followersUrl;
    }
    /**
     * @return The GitHub API URL for those following the collaborator.
     * 
     */
    public String followingUrl() {
        return this.followingUrl;
    }
    /**
     * @return The GitHub API URL for the collaborator&#39;s gists.
     * 
     */
    public String gistsUrl() {
        return this.gistsUrl;
    }
    /**
     * @return The GitHub HTML URL for the collaborator.
     * 
     */
    public String htmlUrl() {
        return this.htmlUrl;
    }
    /**
     * @return The ID of the collaborator.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return The collaborator&#39;s login.
     * 
     */
    public String login() {
        return this.login;
    }
    /**
     * @return The GitHub API URL for the collaborator&#39;s organizations.
     * 
     */
    public String organizationsUrl() {
        return this.organizationsUrl;
    }
    /**
     * @return Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
     * 
     */
    public String permission() {
        return this.permission;
    }
    /**
     * @return The GitHub API URL for the collaborator&#39;s received events.
     * 
     */
    public String receivedEventsUrl() {
        return this.receivedEventsUrl;
    }
    /**
     * @return The GitHub API URL for the collaborator&#39;s repositories.
     * 
     */
    public String reposUrl() {
        return this.reposUrl;
    }
    /**
     * @return Whether the user is a GitHub admin.
     * 
     */
    public Boolean siteAdmin() {
        return this.siteAdmin;
    }
    /**
     * @return The GitHub API URL for the collaborator&#39;s starred repositories.
     * 
     */
    public String starredUrl() {
        return this.starredUrl;
    }
    /**
     * @return The GitHub API URL for the collaborator&#39;s subscribed repositories.
     * 
     */
    public String subscriptionsUrl() {
        return this.subscriptionsUrl;
    }
    /**
     * @return The type of the collaborator (ex. `user`).
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return The GitHub API URL for the collaborator.
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCollaboratorsCollaborator defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String eventsUrl;
        private String followersUrl;
        private String followingUrl;
        private String gistsUrl;
        private String htmlUrl;
        private Integer id;
        private String login;
        private String organizationsUrl;
        private String permission;
        private String receivedEventsUrl;
        private String reposUrl;
        private Boolean siteAdmin;
        private String starredUrl;
        private String subscriptionsUrl;
        private String type;
        private String url;
        public Builder() {}
        public Builder(GetCollaboratorsCollaborator defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.eventsUrl = defaults.eventsUrl;
    	      this.followersUrl = defaults.followersUrl;
    	      this.followingUrl = defaults.followingUrl;
    	      this.gistsUrl = defaults.gistsUrl;
    	      this.htmlUrl = defaults.htmlUrl;
    	      this.id = defaults.id;
    	      this.login = defaults.login;
    	      this.organizationsUrl = defaults.organizationsUrl;
    	      this.permission = defaults.permission;
    	      this.receivedEventsUrl = defaults.receivedEventsUrl;
    	      this.reposUrl = defaults.reposUrl;
    	      this.siteAdmin = defaults.siteAdmin;
    	      this.starredUrl = defaults.starredUrl;
    	      this.subscriptionsUrl = defaults.subscriptionsUrl;
    	      this.type = defaults.type;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder eventsUrl(String eventsUrl) {
            if (eventsUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "eventsUrl");
            }
            this.eventsUrl = eventsUrl;
            return this;
        }
        @CustomType.Setter
        public Builder followersUrl(String followersUrl) {
            if (followersUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "followersUrl");
            }
            this.followersUrl = followersUrl;
            return this;
        }
        @CustomType.Setter
        public Builder followingUrl(String followingUrl) {
            if (followingUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "followingUrl");
            }
            this.followingUrl = followingUrl;
            return this;
        }
        @CustomType.Setter
        public Builder gistsUrl(String gistsUrl) {
            if (gistsUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "gistsUrl");
            }
            this.gistsUrl = gistsUrl;
            return this;
        }
        @CustomType.Setter
        public Builder htmlUrl(String htmlUrl) {
            if (htmlUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "htmlUrl");
            }
            this.htmlUrl = htmlUrl;
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder login(String login) {
            if (login == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "login");
            }
            this.login = login;
            return this;
        }
        @CustomType.Setter
        public Builder organizationsUrl(String organizationsUrl) {
            if (organizationsUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "organizationsUrl");
            }
            this.organizationsUrl = organizationsUrl;
            return this;
        }
        @CustomType.Setter
        public Builder permission(String permission) {
            if (permission == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "permission");
            }
            this.permission = permission;
            return this;
        }
        @CustomType.Setter
        public Builder receivedEventsUrl(String receivedEventsUrl) {
            if (receivedEventsUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "receivedEventsUrl");
            }
            this.receivedEventsUrl = receivedEventsUrl;
            return this;
        }
        @CustomType.Setter
        public Builder reposUrl(String reposUrl) {
            if (reposUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "reposUrl");
            }
            this.reposUrl = reposUrl;
            return this;
        }
        @CustomType.Setter
        public Builder siteAdmin(Boolean siteAdmin) {
            if (siteAdmin == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "siteAdmin");
            }
            this.siteAdmin = siteAdmin;
            return this;
        }
        @CustomType.Setter
        public Builder starredUrl(String starredUrl) {
            if (starredUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "starredUrl");
            }
            this.starredUrl = starredUrl;
            return this;
        }
        @CustomType.Setter
        public Builder subscriptionsUrl(String subscriptionsUrl) {
            if (subscriptionsUrl == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "subscriptionsUrl");
            }
            this.subscriptionsUrl = subscriptionsUrl;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "type");
            }
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetCollaboratorsCollaborator", "url");
            }
            this.url = url;
            return this;
        }
        public GetCollaboratorsCollaborator build() {
            final var _resultValue = new GetCollaboratorsCollaborator();
            _resultValue.eventsUrl = eventsUrl;
            _resultValue.followersUrl = followersUrl;
            _resultValue.followingUrl = followingUrl;
            _resultValue.gistsUrl = gistsUrl;
            _resultValue.htmlUrl = htmlUrl;
            _resultValue.id = id;
            _resultValue.login = login;
            _resultValue.organizationsUrl = organizationsUrl;
            _resultValue.permission = permission;
            _resultValue.receivedEventsUrl = receivedEventsUrl;
            _resultValue.reposUrl = reposUrl;
            _resultValue.siteAdmin = siteAdmin;
            _resultValue.starredUrl = starredUrl;
            _resultValue.subscriptionsUrl = subscriptionsUrl;
            _resultValue.type = type;
            _resultValue.url = url;
            return _resultValue;
        }
    }
}
