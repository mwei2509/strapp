# wip - rewriting strapp cli (previously typescript) in go

MVP:

- assumptions: `-y` yes for everything, otherwise prompt
  - Github (need to install `gh` client if does not exist)
  - Create app directory
  - Docker (install docker if not exists)

# STRAPP, a cli for boot(strapp)ing apps

Strapp is a command line interface that scaffolds your application from source code structure to deployments, taking care of containerizations, datastores, continuous integration in between.

The simplest way to set up an application:
`strapp app:create my-new-app --type api --framework koa --db postgres`

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

### notes

cobra: `go get -u github.com/spf13/cobra`
cobra cli: `go install github.com/spf13/cobra-cli@latest`
pflag doc: https://pkg.go.dev/github.com/spf13/pflag#section-readme

`strapp init`

types:

# Usage

## Create New App

```
strapp app:create <app-name> \
  --type <api|web> \
  --language <language dependent on type> \
  --framework <framework dependent on language> \
  --database <postgres|none> \
  --orm <orm dependent on framework and database>
```

This will generate directory `app-name` scaffolded with a .strapp-config file in the root.
If you ever need to change a config, you can update the file

## Config File

### monorepo

There's 2 ways to create a monorepo.

You can initialize a monorepo by feeding multiple lines
You can create a monorepo by feeding multiple types in order

`strapp app:create my-app`

- If you don't indicate a type, it will start an empty app
- This creates the directory `my-app` with a default `.strapp-config.yml` and `.strapp-config.templates.yml`
- At this point, you can update the .strapp-config to the configuration you would like, then run, in project directory, run

`strapp app:init`

To manually add a new project run in project directory
`strapp app:add backend-project --type api`
`strapp app:add frontend-project --type web`

`strapp app:create my-app --type api`
`strapp app:create my-app --type web`
`strapp app:create my-app --type monorepo`
`strapp app:create my-app --project `

`strapp app create`
`strapp app create`
`strapp database create`
`strapp orm create`

`strapp add database`
`strapp add orm`
`strapp add`

```
strapp app:create \
  -t api -l go -d postgres \
  -t web -l typescript -f react
```

## TYPES of apps

1. _Foo Bar Api_, `foo-bar-api` - basic api that has one endpoint that returns a single endpoint that returns `{ "foo": "bar" }`.

2. _Get Weather App_, `get-weather-app` - basic api that has one endpoint that returns the weather (from https://openweathermap.org/api) and a web app that displays the weather

3. _Hello World Web_, `hello-world-web` - basic web app that has one view that says "Hello World"

4. _Create Users Web App_, `create-users-with-login` - basic api + web that has a way to create users and for users to login. Comes with a db and orm.

5. _Create Users API_, `create-users-api` - basic MVC type application that includes a database and orm

scenarios:

- if an API application has no ORM or database and no web, create a `foo-bar-api`
- if an API application has a database and no web, create a `create-users-api` (use sequelize if no ORM is stated)
- if an API application has a database AND web, create a `create-users-with-login` app
- If an API application has no ORM or database but has a web, create a `get-weather-app`
- If a Web application has no API, create `hello-world-app`

# sample .strapprc

single application

```yaml
name: my-app
services:
  - name: my-app # same as top level name
    type: rest-api
    language: typescript
    framework: koa
    orm: sequelize
    datastores:
      - my-db
    deployment: heroku
datastores:
  my-db:
    type: postgresql-db
    deployment: heroku
cicd: circleci
```

monorepo

```yaml
name: my-app
type: monorepo | microservice
# microservice will break services into separate repos within main directory
# monorepo will keep it all in one place
services:
  - name: backend-service
    type: rest-api
    language: typescript
    framework: koa
    orm: sequelize
    deployment: aws
    datastores:
      - my-db
  - name: frontend-service
    type: frontend
    language: typescript
    framework: react
    css: tailwind
    state_management: redux
    deployment: surge
datastores:
  my-db:
    type: postgresql-db
    deployment: aws
cicd: circleci
# maybe include terraform stuff
```

microservice

```yaml

```
