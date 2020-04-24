import * as github from "@pulumi/github";

const repo = new github.Repository("demo-repo", {
    description: "Generated from automated test",
    private: true,
});

export const repoName = repo.name;
