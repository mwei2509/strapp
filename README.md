# WIP for strapp.sh

- wip - still in concept / experiment stage\*

# STRAPP, a cli for boot(strapp)ing apps

Strapp is a command line interface that scaffolds your application from source code structure to deployments, taking care of containerizations, datastores, continuous integration and everything in between.

The simplest way to set up an application:
`strapp app:create my-new-app --type api --framework koa --db postgres`

Or initialize a directory with a `.strapprc` file that looks like this:

```
name: my-app
type: hello-world

# simplest single app
services:
  - name: my-api
    type: rest-api
    language: typescript
    framework: koa
    orm: sequelize
    datastores:
      - my-db
  - name: my-web
    type: frontend
    language: typescript
    framework: react
    css: tailwind
    state_management: redux

datastores:
  my-db:
    type: postgresql-db

cicd: circleci
```

Strapp will create the app directory and everything needed to set up a simple `Hello, World` application with a github repo, REST endpoints, an initial migration, Docker, and CICD

To bring up your app locally:
`docker compose up`

## Installation

```
brew install mwei2509/taps/strapp
```

## Development

### Local

Clone this repository to `$GOPATH/github.com/src/mwei2509/strapp`
Run: `go run main.go <args>`

### Deploy

New package is deployed via tag push

```
git tag -a <tag> -m "<message>"
git push origin <tag>
```

Use `v<major>.<minor>` for versioning. CircleCI creates the binary and updates the [homebrew repo](https://github.com/mwei2509/homebrew-taps).
