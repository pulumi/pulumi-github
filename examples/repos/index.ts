// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as github from "@pulumi/github";

const name = "repos";

// Add a repository
const repo = new github.repos.Repository(name, {
    name: "some-repo",
    description: "Some repository",
    autoInit: true
});

// Add a collaborator to the repository
const collaborator = new github.repos.Collaborator(name, {
    repository: repo.name,
    username: "e38ce0e-collaborator",
    permission: "pull"
});

// Add a deploy key
const key = new github.repos.DeployKey(name, {
    title: "Some key",
    repository: repo.name,
    key: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC+uT30nCMa0fe5YcFy8eptiOlD2hjB+NNeaqOq2nOC7Hj3slk2HwLtonv2xCKDB1VstSUDYU/btd+7SKfq3m8nCHhcYmeqmLHqEhyZDA5Gdnr0XRtbQ020M5ClWkNLxxQ8lvZ6ovkEIOJF1vrdsBXkI342yuzFDfQy7yl0DJpP5PagYqXVgBKtl/argO11oryDpF7j71HjziPvBzndsqbC2RmxzHiGii7t97vHzM2f2J63QiSupvY393EAbR0j+FDRZ48xoFr5xIclk0jJIDx3wtv4xLJ0eJkebOelf63wE4h7cUyvwTLFN+veB+fvvI/WmYS5omsWH3Xq4huuuz4aa+fXljaBfs66DXAAJkY+4H2rBh99chJGmzybT4upfxovLBxHtCBSejbeJ+4Q5Wulns2kxB6iEXFB9Rtk8nIdPa8aWA4eIPqUKwZCy7fdVPt7QBjkydeVmhsbiTLw90QfXELcLGBDp01FYl3+F8LySdUicuoe2a6AVuaNN8vYVO+r1SujX1QMVOVLFAmQh8jI4wtVc+861sOGedMp5RHQoiTMgdqzkXAvNS4ON2DSYQKdOkfpeGD/YszxHRogIzOQN+NhNx0JnZqrkVxU0mOfAUdrwNGuZJ9h7Qkilca77pwCDzu9ZzRSf/GfLwtK5WTKl6Cr282oWrLTA6BjkGAz9Q== your_email@example.com"
});

// Add a label
const label = new github.repos.Label(name, {
    repository: repo.name,
    name: "enhancement",
    color: "3F0E44"
});

// Add a repository webhook
const webhook = new github.repos.Webhook(name, {
    repository: repo.name,
    configuration: github.WebhookConfiguration = {
        url: "https://google.com",
        content_type: github.ContentTypes.Form,
    },
    events: ["issues"],
    active: false
});

// Add a branch protection
const protection = new github.repos.BranchProtection(name, {
    repository: repo.name,
    branch: "master",
    enforceAdmins: true
});
