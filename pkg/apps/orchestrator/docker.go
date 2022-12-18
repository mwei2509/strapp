package orchestrator

// if there are datastores, add them

// for each service
var serviceTmpl string = `
{{.Name}}:
	container_name: {{.Name}}
	build:
		context: .
		dockerfile: {{.Directory}}/Dockerfile
	environment:
		NODE_ENV: development
		PORT: {{.Port}}
	volumes:
		- ./{{.Directory}}:/app:delegated
	commands: sh -c "{{.StartCommand}}"
`
