// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryBranchesBranch {
    /**
     * @return Name of the branch.
     * 
     */
    private String name;
    /**
     * @return Whether the branch is protected.
     * 
     */
    private Boolean protected_;

    private GetRepositoryBranchesBranch() {}
    /**
     * @return Name of the branch.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Whether the branch is protected.
     * 
     */
    public Boolean protected_() {
        return this.protected_;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryBranchesBranch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private Boolean protected_;
        public Builder() {}
        public Builder(GetRepositoryBranchesBranch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.protected_ = defaults.protected_;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetRepositoryBranchesBranch", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter("protected")
        public Builder protected_(Boolean protected_) {
            if (protected_ == null) {
              throw new MissingRequiredPropertyException("GetRepositoryBranchesBranch", "protected_");
            }
            this.protected_ = protected_;
            return this;
        }
        public GetRepositoryBranchesBranch build() {
            final var _resultValue = new GetRepositoryBranchesBranch();
            _resultValue.name = name;
            _resultValue.protected_ = protected_;
            return _resultValue;
        }
    }
}