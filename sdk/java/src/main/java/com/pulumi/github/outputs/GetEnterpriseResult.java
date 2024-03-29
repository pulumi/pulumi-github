// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetEnterpriseResult {
    /**
     * @return The time the enterprise was created.
     * 
     */
    private String createdAt;
    /**
     * @return The database ID of the enterprise.
     * 
     */
    private Integer databaseId;
    /**
     * @return The description of the enterprise.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the enterprise.
     * 
     */
    private String name;
    /**
     * @return The URL slug identifying the enterprise.
     * 
     */
    private String slug;
    /**
     * @return The url for the enterprise.
     * 
     */
    private String url;

    private GetEnterpriseResult() {}
    /**
     * @return The time the enterprise was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The database ID of the enterprise.
     * 
     */
    public Integer databaseId() {
        return this.databaseId;
    }
    /**
     * @return The description of the enterprise.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the enterprise.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The URL slug identifying the enterprise.
     * 
     */
    public String slug() {
        return this.slug;
    }
    /**
     * @return The url for the enterprise.
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEnterpriseResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private Integer databaseId;
        private String description;
        private String id;
        private String name;
        private String slug;
        private String url;
        public Builder() {}
        public Builder(GetEnterpriseResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.databaseId = defaults.databaseId;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.slug = defaults.slug;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetEnterpriseResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder databaseId(Integer databaseId) {
            if (databaseId == null) {
              throw new MissingRequiredPropertyException("GetEnterpriseResult", "databaseId");
            }
            this.databaseId = databaseId;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetEnterpriseResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetEnterpriseResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetEnterpriseResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder slug(String slug) {
            if (slug == null) {
              throw new MissingRequiredPropertyException("GetEnterpriseResult", "slug");
            }
            this.slug = slug;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetEnterpriseResult", "url");
            }
            this.url = url;
            return this;
        }
        public GetEnterpriseResult build() {
            final var _resultValue = new GetEnterpriseResult();
            _resultValue.createdAt = createdAt;
            _resultValue.databaseId = databaseId;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.slug = slug;
            _resultValue.url = url;
            return _resultValue;
        }
    }
}
