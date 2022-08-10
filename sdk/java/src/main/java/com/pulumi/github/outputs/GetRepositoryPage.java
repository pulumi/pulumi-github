// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.GetRepositoryPageSource;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRepositoryPage {
    private final String cname;
    private final Boolean custom404;
    /**
     * @return URL to the repository on the web.
     * 
     */
    private final String htmlUrl;
    private final List<GetRepositoryPageSource> sources;
    private final String status;
    private final String url;

    @CustomType.Constructor
    private GetRepositoryPage(
        @CustomType.Parameter("cname") String cname,
        @CustomType.Parameter("custom404") Boolean custom404,
        @CustomType.Parameter("htmlUrl") String htmlUrl,
        @CustomType.Parameter("sources") List<GetRepositoryPageSource> sources,
        @CustomType.Parameter("status") String status,
        @CustomType.Parameter("url") String url) {
        this.cname = cname;
        this.custom404 = custom404;
        this.htmlUrl = htmlUrl;
        this.sources = sources;
        this.status = status;
        this.url = url;
    }

    public String cname() {
        return this.cname;
    }
    public Boolean custom404() {
        return this.custom404;
    }
    /**
     * @return URL to the repository on the web.
     * 
     */
    public String htmlUrl() {
        return this.htmlUrl;
    }
    public List<GetRepositoryPageSource> sources() {
        return this.sources;
    }
    public String status() {
        return this.status;
    }
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryPage defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String cname;
        private Boolean custom404;
        private String htmlUrl;
        private List<GetRepositoryPageSource> sources;
        private String status;
        private String url;

        public Builder() {
    	      // Empty
        }

        public Builder(GetRepositoryPage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cname = defaults.cname;
    	      this.custom404 = defaults.custom404;
    	      this.htmlUrl = defaults.htmlUrl;
    	      this.sources = defaults.sources;
    	      this.status = defaults.status;
    	      this.url = defaults.url;
        }

        public Builder cname(String cname) {
            this.cname = Objects.requireNonNull(cname);
            return this;
        }
        public Builder custom404(Boolean custom404) {
            this.custom404 = Objects.requireNonNull(custom404);
            return this;
        }
        public Builder htmlUrl(String htmlUrl) {
            this.htmlUrl = Objects.requireNonNull(htmlUrl);
            return this;
        }
        public Builder sources(List<GetRepositoryPageSource> sources) {
            this.sources = Objects.requireNonNull(sources);
            return this;
        }
        public Builder sources(GetRepositoryPageSource... sources) {
            return sources(List.of(sources));
        }
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }        public GetRepositoryPage build() {
            return new GetRepositoryPage(cname, custom404, htmlUrl, sources, status, url);
        }
    }
}