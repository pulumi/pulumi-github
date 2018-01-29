// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as github from "@pulumi/github";

const name = "teams";

// Add a team
const team = new github.teams.Team(name, {
    name: "some-team",
    description: "Some team"
});

// Add a membership
const teamMembership = new github.teams.Membership(name, {
    teamId: team.id,
    username: "e38ce0e-collaborator",
    role: "member"
});

// Add a repository
const repo = new github.repos.Repository(name, {
    name: "team-repo",
    description: "Test repository"
});

// Grant the team permission to pull from the repo
const teamRepo = new github.teams.Repository(name, {
    teamId: team.id,
    repository: repo.name,
    permission: "pull"
});
