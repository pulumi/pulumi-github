import pulumi
import pulumi_github as github

repo = github.Repository("demo-repo",
    description="Generated from automated Python test",
    visibility="private"
)

pulumi.export('repo_name', repo.name)


