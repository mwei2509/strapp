package koa

import (
	"text/template"
)

var packageJson string = `
{
  "name": "{{.Name}}",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "build": "./node_modules/typescript/bin/tsc",
    "start": "node index.js",
    "start:dev": "./node_modules/nodemon/bin/nodemon.js --watch './**/*.ts' --exec 'ts-node' src/app.ts -- --inspect=0.0.0.0:{{.DebuggerPort}}"
  },
  "author": "",
  "license": "ISC",
  "dependencies": {
    "dotenv": "^16.0.3",
    "koa": "^2.14.1",
    "koa-bodyparser": "^4.3.0",
    "koa-json": "^2.0.2",
    "koa-logger": "^3.2.1",
    "koa-router": "^12.0.0"
  },
  "devDependencies": {
    "@types/node": "^18.11.12",
    "@types/validator": "^13.7.10",
    "nodemon": "^2.0.20",
    "ts-node": "^10.9.1",
    "typescript": "^4.9.3"
  }
}
`

var app string = `
import * as Koa from 'koa';
import * as logger from 'koa-logger';
import * as json from 'koa-json';
import router from './routes';

require('dotenv').config();

const app = new Koa();

// middleware
app.use(json());
app.use(logger());

// routes
app.use(router.routes()).use(router.allowedMethods());

app.listen(process.env.PORT, () => {
    console.log('API started!')
});
`

var routes string = `
import * as Router from "koa-router";

const router = new Router();

router.get('/healthcheck', async (ctx) => {
    ctx.body = "ok";
});

export default router;
`

// if i need to create console.log(`something ${whatever}`)
// `console.log(` + "`" + `something ${{"{"}}{{.Whatever}}{{"}"}}` + "`" + `);`

func (k *KoaApp) CreateTypescriptAppFiles() error {
	// package.json
	tmpl, err := template.New("").Parse(packageJson)
	if err != nil {
		return err
	}
	k.createFileFromTemplate(tmpl, "/package.json")

	// src/app.ts
	tmpl, err = template.New("").Parse(app)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/src/app.ts"); err != nil {
		return err
	}

	// src/routes.ts
	tmpl, err = template.New("").Parse(routes)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/src/routes.ts"); err != nil {
		return err
	}
	return err
}
