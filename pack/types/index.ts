// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

// An Asset is data related to a release.  See https://developer.github.com/v3/repos/releases/ for details.
export interface Asset {
    url: string;
    browserDownloadUrl: string;
    id: number;
    name: string;
    label: string;
    state: string;
    contentType: string;
    size: number;
    downloadCount: number;
    createdAt: string;
    updatedAt: string;
    uploader: User;
}

// A Comment is text written by a user that pertains to a repo artifact at a specific commit ID.  See
// https://developer.github.com/v3/repos/comments/ for more information.
export interface Comment {
    htmlUrl: string;
    url: string;
    id: number;
    body: string;
    path: string;
    position: number;
    line: number;
    commitId: string;
    user: User;
    createdAt: string;
    updatedAt: string;
}

// A Deployment is a request for a specific ref (branch, SHA, tag) to be deployed.  GitHub then dispatches deployment
// events that external services can listen for and act on.  This enables developers and organizations to build
// loosely-coupled tooling around deployments, without having to worry about implementation details of delivering
// different types of applications (e.g., web, native).  See https://developer.github.com/v3/repos/deployments/.
export interface Deployment {
    url: string;
    id: number;
    // sha is the commit SHA for which this deployment was created.
    sha: string;
    ref: string;
    task: string; // TODO: enum.
    /// payload is optional extra information for this deployment.
    payload?: string;
    // environment is the optional environment to deploy to.  Default: `"production"`.
    environment?: string;
    // description is an optional human-readable description added to the deployment.
    description?: string;
    creator: User;
    createdAt: string;
    updatedAt: string;
    type: string; // TODO: enum.
    siteAdmin: boolean;
}

// A Download is a file that may be downloaded from a repo.  See https://developer.github.com/v3/repos/downloads/
export interface Download {
    url: string;
    htmlUrl: string;
    id: number;
    name: string;
    description: string;
    size: number;
    downloadCount: number;
    contentType: string;
}

export interface Error {
    message?: string;
}

// A Gist is a collection of individually published code snippets.  See https://developer.github.com/v3/gists/.
export interface Gist {
    url: string;
    forksUrl: string;
    commitsUrl: string;
    id: string;
    description: string;
    public: boolean;
    owner: User;
    user: User;
    files: Map<string, GistFile>;
    comments: number;
    commentsUrl: string;
    htmlUrl: string;
    gitPullUrl: string;
    gitPushUrl: string;
    createdAt: string;
    updatedAt: string;
    forks: GistFork[];
    history: GistRevision[];
}

// A GistFile is a file associated with a Gist.
export interface GistFile {
    size: number;
    rawUrl: string;
    type: string;
    language: string;
    truncated: boolean;
    content: string;
}

// A GistFork is a forked Gist, in the usual Git sense.
export interface GistFork {
    user: User;
    url: string;
    // id is the SHA hash of the fork.
    id: string;
    createdAt: string;
    updatedAt: string;
}

// A GistRevision is a specific revision of a Gist.
export interface GistRevision {
    url: string;
    // version is the SHA hash of this revision.
    version: string;
    user: User;
    changeStatus: GistChangeStatus;
    committedAt: string;
}

// GistChangeStatus reflects the status of a specific Gist revision.
export interface GistChangeStatus {
    deletions: number;
    additions: number;
    total: number;
}

// An Issue tracks issues -- work items, bugs, project todos, etc. -- pertaining to a particular repo.  For more
// information, please see https://developer.github.com/v3/issues/.
export interface Issue {
    id: number;
    url: string;
    labelsUrl: string;
    commentsUrl: string;
    eventsUrl: string;
    htmlUrl: string
    number: number;
    state: string; // TODO: enum.
    title: string;
    body: string;
    user: User;
    labels: Label[];
    assignee?: User;
    milestone?: Milestone;
    locked: boolean;
    comments: number;
    pullRequest?: IssuePullRequest;
    closedAt?: string;
    createdAt: string;
    updatedAt: string;
    closedBy?: User;
}

// An IssuePullRequest is used when an issue is also a pull request.
export interface IssuePullRequest {
    url: string;
    htmlUrl: string;
    diffUrl: string;
    patchUrl: string;
}

// A Label is used to tag issues with informative labels.
export interface Label {
    url: string;
    name: string;
    color: string;
}

// A Milestone is used to classify an issue for purposes of project planning and tracking.
export interface Milestone {
    id: number;
    url: string;
    htmlUrl: string;
    labelsUrl: string;
    number: number;
    state: string; // TODO: enum.
    title: string;
    description: string;
    creator: User;
    openIssues: number;
    closedIssues: number;
    createdAt: string;
    updatedAt: string;
    closedAt: string;
    dueOn: string;
}

// An Organization is a company, team, or other special entity that a User and/or Repo may belong to.
export interface Organization {
    id: number;
    login: string;
    url: string;
    reposUrl: string;
    eventsUrl: string;
    membersUrl: string;
    publicMembersUrl: string;
    avatarUrl: string;
    description?: string;
}

// A PageBuild is a built GitHub Pages repo, at a particular commit, published at a given URL.  For more information
// about page builds, see https://developer.github.com/v3/repos/pages/#list-pages-builds.
export interface PageBuild {
    url: string;
    status: string;
    error: Error;
    pusher: User;
    commit: string;
    duration: number;
    createdAt: string;
    updatedAt: string;
}

// Permissions tracks the permissions granted on a given repo.
export interface Permissions {
    admin: boolean;
    push: boolean;
    pull: boolean;
}

// Plan describes the plan an account is currently subscribed to.
export interface Plan {
    name: string;
    space: number;
    privateRepos: number;
    collaborators: number;
}

// A PullRequest is an object representing a request to merge functionality into a branch by way of a Git pull command.
// For more information, see https://developer.github.com/v3/pulls/.
export interface PullRequest {
    id: number;
    url: string;
    htmlUrl: string;
    diffUrl: string;
    patchUrl: string;
    issueUrl: string;
    commitsUrl: string;
    reviewCommentsUrl: string;
    reviewCommentUrl: string;
    commentsUrl: string;
    statusesUrl: string;
    number: number;
    state: string; // TODO: enum.
    title: string;
    body: string;
    assignee: User;
    milestone: Milestone;
    locked: boolean;
    createdAt: string;
    updatedAt: string;
    closedAt: string;
    mergedAt: string;
    head: PullRequestCommit;
    base: PullRequestCommit;
    _links: PullRequestLinks[];
    user: User;
    mergeCommitSha: string;
    merged: boolean;
    mergeable: boolean;
    mergedBy: User;
    comments: number;
    commits: number;
    additions: number;
    deletions: number;
    changedFiles: number;
}

// PullRequestCommit references a specific Git label, ref, and commit, for a given User and Repo.
export interface PullRequestCommit {
    label: string;
    ref: string;
    sha: string;
    user: User;
    repo: Repository;
}

// PullRequestLinks contains all the pertinent links associated with a PullRequest object.
export interface PullRequestLinks {
    self: PullRequestLink;
    html: PullRequestLink;
    issue: PullRequestLink;
    comments: PullRequestLink;
    reviewComments: PullRequestLink
    reviewComment: PullRequestLink;
    commits: PullRequestLink;
    statuses: PullRequestLink;
}

// PullRequestLink contains information about a single PullRequest hyperlink.
export interface PullRequestLink {
    href: string;
}

// A Release is a named grouping of code and assets that represents a point in time snapshot of software in the given
// Repo.  For more information, refer to https://developer.github.com/v3/repos/releases/.
export interface Release {
    url: string;
    htmlUrl: string;
    assetsUrl: string;
    uploadUrl: string;
    tarballUrl: string;
    zipballUrl: string;
    id: number;
    tagName: string;
    targetCommitish: string;
    name: string;
    body: string;
    draft: boolean;
    prerelease: boolean;
    createdAt: string;
    publishedAt: string;
    author: User;
    assets: Asset[];
}

// A Repo is a repository where code, issues, and people come together on GitHub, and corresponds 1:1 with an actual
// Git repo.  Please see https://developer.github.com/v3/repos/.
export interface Repository {
    id: number;
    owner: User;
    name: string;
    fullName: string;
    description: string;
    private: boolean;
    fork: boolean;
    url: string;
    htmlUrl: string;
    archiveUrl: string;
    assigneesUrl: string;
    blobsUrl: string;
    branchesUrl: string;
    cloneUrl: string;
    collaboratorsUrl: string;
    commentsUrl: string;
    commitsUrl: string;
    compareUrl: string;
    contentsUrl: string;
    contributorsUrl: string;
    downloadsUrl: string;
    eventsUrl: string;
    forksUrl: string;
    gitCommitsUrl: string;
    gitRefsUrl: string;
    gitTagsUrl: string;
    gitUrl: string;
    hooksUrl: string;
    issueCommentUrl: string;
    issueEventsUrl: string;
    issuesUrl: string;
    keysUrl: string;
    labelsUrl: string;
    languagesUrl: string;
    mergesUrl: string;
    milestonesUrl: string;
    mirrorUrl: string;
    notificationsUrl: string;
    pullsUrl: string;
    releasesUrl: string;
    sshUrl: string;
    stargazersUrl: string;
    statusesUrl: string;
    subscribersUrl: string;
    subscriptionUrl: string;
    svnUrl: string;
    tagsUrl: string;
    teamsUrl: string;
    treesUrl: string;
    homepage: string;
    language?: string;
    forksCount: number;
    stargazersCount: number;
    watchersCount: number;
    size: number;
    defaultBranch: string;
    openIssuesCount: number;
    hasIssues: boolean;
    hasWiki: boolean;
    hasPages: boolean;
    hasDownloads: boolean;
    pushedAt: string;
    createdAt: string;
    updatedAt: string;
    permissions: Permissions;
    subscribersCount: number;
    organization: User;
    parent?: Repository;
    source?: Repository;
}

// A Team is a logical grouping of Users within a given Organization.  See https://developer.github.com/v3/orgs/teams/.
export interface Team {
    id: number;
    url: string;
    name: string;
    slug: string;
    description: string;
    privacy: string; // TODO: enum.
    permission: string; // TODO: enum.
    membersUrl: string;
    repositoriesUrl: string;
}

// A User is an authenticated GitHub user with an identity.  See https://developer.github.com/v3/users/.
export interface User {
    login: string;
    id: number;
    avatarUrl: string;
    gravatarId: string;
    url: string;
    htmlUrl: string;
    followersUrl: string;
    followingUrl: string;
    gistsUrl: string;
    starredUrl: string;
    subscriptionsUrl: string;
    organizationsUrl: string;
    reposUrl: string;
    eventsUrl: string;
    receivedEventsUrl: string;
    type: string;
    siteAdmin: boolean;

    // These aren't present in short user payloads:
    name?: string;
    company?: string;
    blog?: string;
    location?: string;
    email?: string;
    hireable?: boolean;
    bio?: string;
    publicRepos?: number;
    publicGists?: number;
    followers?: number;
    following?: number;
    createdAt?: string;
    updatedAt?: string;

    // These are only available on authenticated users:
    totalPrivateRepos?: number;
    ownedPrivateRepos?: number;
    privateGists?: number;
    diskUsage?: number;
    collaborators?: number;
    plan?: Plan;
}

