import pulumi
import pulumi_github as github

repo = github.Repository("demo-repo",
    description="Generated from automated test",
    private="true"
)

pulumi.export('repo_name', repo.name)


