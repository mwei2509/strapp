# wip - rewriting strapp cli (previously typescript) in go

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
