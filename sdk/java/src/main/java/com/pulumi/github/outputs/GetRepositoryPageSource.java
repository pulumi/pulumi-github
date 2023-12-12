// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryPageSource {
    private String branch;
    private String path;

    private GetRepositoryPageSource() {}
    public String branch() {
        return this.branch;
    }
    public String path() {
        return this.path;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryPageSource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String branch;
        private String path;
        public Builder() {}
        public Builder(GetRepositoryPageSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branch = defaults.branch;
    	      this.path = defaults.path;
        }

        @CustomType.Setter
        public Builder branch(String branch) {
            this.branch = Objects.requireNonNull(branch);
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public GetRepositoryPageSource build() {
            final var _resultValue = new GetRepositoryPageSource();
            _resultValue.branch = branch;
            _resultValue.path = path;
            return _resultValue;
        }
    }
}
