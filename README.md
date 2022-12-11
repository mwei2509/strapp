# wip - rewriting strapp cli (previously typescript) in go

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

Use `v<major>.<minor>` for versioning.  CircleCI creates the binary and updates the [homebrew repo](https://github.com/mwei2509/homebrew-taps).

### notes
cobra: `go get -u github.com/spf13/cobra`
cobra cli: `go install github.com/spf13/cobra-cli@latest`
pflag doc: https://pkg.go.dev/github.com/spf13/pflag#section-readme

`strapp init`

types:

- `backend`
- `frontend`
- `monorepo`
  `strapp app:create --type api`
  `strapp app:create --type web`
  `strapp app:create --type api web`

`strapp app create`
`strapp app create`
`strapp database create`
`strapp orm create`

`strapp add database`
`strapp add orm`
`strapp add`
