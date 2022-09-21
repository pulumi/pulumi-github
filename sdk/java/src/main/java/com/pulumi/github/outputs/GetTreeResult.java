// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.GetTreeEntry;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTreeResult {
    private List<GetTreeEntry> entries;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean recursive;
    private String repository;
    private String treeSha;

    private GetTreeResult() {}
    public List<GetTreeEntry> entries() {
        return this.entries;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> recursive() {
        return Optional.ofNullable(this.recursive);
    }
    public String repository() {
        return this.repository;
    }
    public String treeSha() {
        return this.treeSha;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTreeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetTreeEntry> entries;
        private String id;
        private @Nullable Boolean recursive;
        private String repository;
        private String treeSha;
        public Builder() {}
        public Builder(GetTreeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.entries = defaults.entries;
    	      this.id = defaults.id;
    	      this.recursive = defaults.recursive;
    	      this.repository = defaults.repository;
    	      this.treeSha = defaults.treeSha;
        }

        @CustomType.Setter
        public Builder entries(List<GetTreeEntry> entries) {
            this.entries = Objects.requireNonNull(entries);
            return this;
        }
        public Builder entries(GetTreeEntry... entries) {
            return entries(List.of(entries));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder recursive(@Nullable Boolean recursive) {
            this.recursive = recursive;
            return this;
        }
        @CustomType.Setter
        public Builder repository(String repository) {
            this.repository = Objects.requireNonNull(repository);
            return this;
        }
        @CustomType.Setter
        public Builder treeSha(String treeSha) {
            this.treeSha = Objects.requireNonNull(treeSha);
            return this;
        }
        public GetTreeResult build() {
            final var o = new GetTreeResult();
            o.entries = entries;
            o.id = id;
            o.recursive = recursive;
            o.repository = repository;
            o.treeSha = treeSha;
            return o;
        }
    }
}
