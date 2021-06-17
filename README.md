# merged-prs

merged-prs is a go tool to assist in determining differences between git hashes
based upon the GitHub PR as the vehicle for change.

## Requirements

* Go 1.16 (lower may work)
* Dep
* Git
* GitHub

## Features

* Built in Go
* Given a set of git hashes, a list of merged GitHub Pull Requests will be
  retrieved and parsed.
* List will be output to console containing the PR #, Author, Summary, and a
  link to the PR

## Configuration

Use environment variables `GITHUB_TOKEN` and `GITHUB_OWNER` to respectively specify the
[Personal access token](https://github.com/settings/tokens) and Github owner.


## Installation

In order to use the `merged-prs` tool one should use `go get`

```bash
go get github.com/brianmhunt/merged-prs
```

## Usage

Calling the `merged-prs` tool will act in the current directory's context.

### Flags

```bash
# Runtime Flags
  -path       <Specify path to repo in order to use outside the context of a repo>
  -dev=false  <Ignore the git-flow dev branch paradigm (loose comparison "...")>
```

### Example

```bash
# merged-prs [-test] [-c @lucas] [-dev=false] [-path /path/to/repo] [Previous Git Hash] [Current Git Hash]
# User should specify the older revision first ie. merging `dev` into `master` would necessitate that `master` is the older commit, and `dev` is the newer

$ merged-prs master dev
Determining merged branches between the following hashes: master dev

REPO: Merged GitHub PRs between the following refs: master dev
PR   Author    Description              URL
#55  @lucas    Typo 100 vs 1000         http://github.com/promiseofcake/foo/pull/55
#54  @lucas    LRU Cache Store Results  http://github.com/promiseofcake/foo/pull/54
```

## History

2021-06-17 
- Forked from [promiseofcake/merged-prs](http://github.com/promiseofcake/merged-prs)
- Switch to Go Modules
- Remove Slack integration
- User `GITHUB_TOKEN` and `GITHUB_OWNER` environment variables for config