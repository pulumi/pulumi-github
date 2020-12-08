import * as github from "@pulumi/github";

const repo = new github.Repository("demo-repo", {
    description: "Generated from automated test",
    visibility: "private",
});

export const repoName = repo.name;
