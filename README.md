# Gitlab CI

## Introduction

Many companies in Iran use [Gitlab](https://about.gitlab.com/) as their VCS (Version control software), so it is essential for us
to have a way for managing Gitlab CI over different projects with ease.

## Do more?

Yes, you can even do more with Gitlab. I always define tasks as issues on Gitlab and then create merge request from them.
This way, you always know which Merge request is related to which issue and then you can find out the desciption of the task.

You can review MRs on Gitlab and provides comments there which is a good thing because you can know about the approvement
and etc.

## Why we need pipelines?

It is a great idea to have coding style and structure shared with you team and then you can define pipelines to validate at
least those styles and for the next step you can have tests and validation over project.

## Need to know more?

- [Predefined variables reference](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html)
- [Services](https://docs.gitlab.com/ee/ci/services/)
- [`.gitlab-ci.yml` keyword reference](https://docs.gitlab.com/ee/ci/yaml/)

## Manage your `gitlab-ci`

When you start with `.gitlab-ci.yml` in the root of your project, you may not know about the future but I will told you,
in future it would be a huge file that you cannot find anything on it.
So, it is better to start by spliting and put files in a directory (for example named it `.gitlab/ci/`).

Let's do some hands-on experiece, consider you have a project named `nostrodumos` which is written in Go
and you want to enable pipelines for it.

### Define pipeline for Go

Lets review our requirements:

1. We need to have a `lint` stage based on `golangci-lint`.
2. We need to have a `test` stage.
