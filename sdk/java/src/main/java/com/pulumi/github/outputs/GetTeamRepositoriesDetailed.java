// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetTeamRepositoriesDetailed {
    private Integer repoId;
    private String roleName;

    private GetTeamRepositoriesDetailed() {}
    public Integer repoId() {
        return this.repoId;
    }
    public String roleName() {
        return this.roleName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTeamRepositoriesDetailed defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer repoId;
        private String roleName;
        public Builder() {}
        public Builder(GetTeamRepositoriesDetailed defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.repoId = defaults.repoId;
    	      this.roleName = defaults.roleName;
        }

        @CustomType.Setter
        public Builder repoId(Integer repoId) {
            if (repoId == null) {
              throw new MissingRequiredPropertyException("GetTeamRepositoriesDetailed", "repoId");
            }
            this.repoId = repoId;
            return this;
        }
        @CustomType.Setter
        public Builder roleName(String roleName) {
            if (roleName == null) {
              throw new MissingRequiredPropertyException("GetTeamRepositoriesDetailed", "roleName");
            }
            this.roleName = roleName;
            return this;
        }
        public GetTeamRepositoriesDetailed build() {
            final var _resultValue = new GetTeamRepositoriesDetailed();
            _resultValue.repoId = repoId;
            _resultValue.roleName = roleName;
            return _resultValue;
        }
    }
}