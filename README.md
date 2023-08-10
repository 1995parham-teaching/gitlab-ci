<h1 align="center"> GitLab CI </h1>

## Introduction

Many companies in Iran use [GitLab](https://about.gitlab.com/) as their VCS (Version control software), so it is essential for us
to have a way for managing GitLab CI over different projects with ease.

## Do more?

Yes, you can even do more with GitLab. I always define tasks as issues on GitLab and then create merge request from them.
This way, you always know which Merge request is related to which issue, and then you can find out the description of the task.

You can review merge requests on GitLab and provides comments there which is a good thing because you can know about the
approval, etc.

## Why we need pipelines?

It is a great idea to have coding style and structure shared with you team, and then you can define pipelines to validate at
least those styles and for the next step you can have tests and validation over project.

## Need to know more?

- [Predefined variables reference](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html)
- [Services](https://docs.gitlab.com/ee/ci/services/)
- [`.gitlab-ci.yml` keyword reference](https://docs.gitlab.com/ee/ci/yaml/)

## Manage your `gitlab-ci`

When you start with `.gitlab-ci.yml` in the root of your project, you may not know about the future, but I will tell you,
in future it would be a huge file that you cannot find anything on it.
So, it is better to start by splitting and put files in a directory (for example named it `.gitlab/ci/`).

Let's do some hands-on experience, consider you have a project named `nostrodumos` which is written in Go
and you want to enable pipelines for it.

### Define pipeline for Go

Let's review our requirements:

1. We need to have a `lint` stage based on `golangci-lint`.
2. We need to have a `test` stage.
3. We need to have a `compile` stage which compiles the code.
4. We need to have a `build` stage which builds our Dockerfile.

First we are going to define a new repository name `easy-ci` which contains
templates that you can extend in your project. These templates will reduce a huge
number of code duplicates and make your `.gitlab-ci` cleaner and more readable.

Then we extend what we defined in `easy-ci` in our project repository.

These projects are defined on GitLab.

- [Easy CI](https://gitlab.com/1995parham-teaching/easy-ci)
- [Nostradamus](https://gitlab.com/1995parham-teaching/nostradamus)

```bash
git subtree pull --prefix nostradamus https://gitlab.com/1995parham-teaching/nostradamus.git  main --squash                                                                                                                                                                                                                                                                                      19:03   Linux 6.4.9-zen1-1-zen
git subtree pull --prefix easy-ci https://gitlab.com/1995parham-teaching/easy-ci.git  main --squash                                                                                                                                                                                                                                                                                      19:03   Linux 6.4.9-zen1-1-zen
```

## Variables

It is better to define variables on Organization scope when these variables
indicate team's secrets instead of defining them in each project and repeat
them many times (trust me, in future they may change).
