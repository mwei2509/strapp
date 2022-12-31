package templates

// if there are datastores, add them

// for each service
var nodeServiceTmpl string = `
{{.Name}}:
	container_name: {{.Name}}
	build:
		context: .
		dockerfile: {{.Directory}}/Dockerfile
	environment:
		NODE_ENV: development
		PORT: {{.Port}}
	ports:
		- {{.Port}}:{{.Port}}
		- {{.DebuggerPort}}:{{.DebuggerPort}}
	volumes:
		- ./{{.Directory}}:/app:delegated
	commands: sh -c "{{.StartCommand}}"
`
