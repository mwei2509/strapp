package koa

import (
	"text/template"
)

var tsconfig string = `
{
    "version": "2.0.2",
    "compilerOptions": {
        "outDir": "./dist",
        "lib": ["es5", "es6"],
        "target": "es6",
        "module": "commonjs",
        "moduleResolution": "node",
        "emitDecoratorMetadata": true,
        "experimentalDecorators": true,
        "strictNullChecks": true,
        "esModuleInterop": true
    },
    "include": ["./src/**/*"]
}
`

// var packageJson string = `
// {
//   "name": "foo-bar-api",
//   "version": "1.0.0",
//   "description": "a foo bar api",
//   "main": "dist/server.js",
//   "scripts": {
//     "test": "jest",
//     "lint": "npx eslint .",
//     "build": "npx tsc",
//     "start": "node dist/server.js",
//     "start:dev": "./node_modules/nodemon/bin/nodemon.js --watch './**/*.ts' --exec 'ts-node' src/server.ts -- --inspect=0.0.0.0:{{.DebuggerPort}}"
//   },
//   "author": "",
//   "license": "ISC",
//   "dependencies": {
//     "@koa/router": "^12.0.0",
//     "dotenv": "^16.0.3",
//     "koa": "^2.14.1",
//     "koa-bodyparser": "^4.3.0",
//     "koa-json": "^2.0.2",
//     "koa-logger": "^3.2.1",
//     "koa-router": "^12.0.0"
//   },
//   "devDependencies": {
//     "@types/jest": "^29.2.4",
//     "@types/koa__router": "^12.0.0",
//     "@types/koa-json": "^2.0.20",
//     "@types/koa-logger": "^3.1.2",
//     "@types/koa-router": "^7.4.4",
//     "@types/supertest": "^2.0.12",
//     "@typescript-eslint/eslint-plugin": "^5.46.1",
//     "@typescript-eslint/parser": "^5.46.1",
//     "eslint": "^8.29.0",
//     "eslint-config-airbnb-typescript": "^17.0.0",
//     "eslint-config-standard-with-typescript": "^23.0.0",
//     "eslint-plugin-import": "^2.26.0",
//     "eslint-plugin-n": "^15.6.0",
//     "eslint-plugin-promise": "^6.1.1",
//     "jest": "^29.3.1",
//     "nodemon": "^2.0.20",
//     "supertest": "^6.3.3",
//     "ts-jest": "^29.0.3",
//     "ts-node": "^10.9.1",
//     "typescript": "^4.9.4"
//   }
// }
// `

var jestConfig string = `
{
  "verbose": true,
  "transform": {
    "^.+\\.tsx?$": "ts-jest"
  }
}
`

var prettierRc string = `
{
  "semi": true,
  "trailingComma": "es5",
  "singleQuote": true,
  "printWidth": 100
}
`

var prettierIgnore string = `
dist
node_modules
`

var gitIgnore string = `
dist
node_modules

.env
`

var eslintJson string = `
{
  "env": {
    "browser": true,
    "es2021": true
  },
  "extends": [
    "airbnb-base",
    "airbnb-typescript/base"
  ],
  "parserOptions": {
    "project": "./tsconfig.json"
  },
  "rules": {}
}
`

var env string = `
PORT={{.Port}}
`

var srcRoutes string = `
import Router from '@koa/router';

const router = new Router();

router.get('/foo-bar', (ctx) => {
  ctx.body = {
    foo: 'bar',
  };
});

export default router;
`

var srcServer string = `
import Koa from 'koa';
import logger from 'koa-logger';
import json from 'koa-json';
import dotenv from 'dotenv';
import router from './routes';

dotenv.config();

const app = new Koa();

// middleware
app.use(json());
app.use(logger());

// routes
app.use(router.routes()).use(router.allowedMethods());

const server = app.listen(process.env.PORT, () => {
  // eslint-disable-next-line
  console.log('API started!');
});

export default server;
`

var srcServerSpec string = `
import request from 'supertest';
import server from './server';

describe('Foo Bar Test Suite', () => {
  afterAll(() => {
    server.close();
  });
  it('tests /foo-bar endpoint', async () => {
    const response = await request(server).get('/foo-bar');
    expect(response.body).toEqual({ foo: 'bar' });
    expect(response.statusCode).toBe(200);
  });
});
`

// if i need to create console.log(`something ${whatever}`)
// `console.log(` + "`" + `something ${{"{"}}{{.Whatever}}{{"}"}}` + "`" + `);`

func (k *KoaApp) CreateFooBarFiles() error {
	// package.json
	// tmpl, err := template.New("").Parse(packageJson)
	// if err != nil {
	// 	return err
	// }

	// exPath, err := exec.CommandContext("api_base")
	// if err != nil {
	// 	return err
	// }
	// buff, err := ioutil.ReadFile(exPath + "/pkg/apps/node/templates/koa/api_base/template.package.json")
	// if err != nil {
	// 	return err
	// }
	// packageJson := string(buff)
	tmpl, err := template.New("").ParseFiles("api_base/template.package.json")
	if err != nil {
		return err
	}

	if err = k.createFileFromTemplate(tmpl, "/package.json"); err != nil {
		return err
	}

	// tsconfig.json
	tmpl, err = template.New("").Parse(tsconfig)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/tsconfig.json"); err != nil {
		return err
	}

	// jest.config.json
	tmpl, err = template.New("").Parse(jestConfig)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/jest.config.json"); err != nil {
		return err
	}

	// .prettierrc
	tmpl, err = template.New("").Parse(prettierRc)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/.prettierrc"); err != nil {
		return err
	}

	// .prettierignore
	tmpl, err = template.New("").Parse(prettierIgnore)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/.prettierignore"); err != nil {
		return err
	}

	// .gitignore
	tmpl, err = template.New("").Parse(gitIgnore)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/.gitignore"); err != nil {
		return err
	}

	// .eslintrc.json
	tmpl, err = template.New("").Parse(eslintJson)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/.eslintrc.json"); err != nil {
		return err
	}

	// .env
	tmpl, err = template.New("").Parse(env)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/.env"); err != nil {
		return err
	}

	// src/routes.ts
	tmpl, err = template.New("").Parse(srcRoutes)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/src/routes.ts"); err != nil {
		return err
	}

	// src/server.ts
	tmpl, err = template.New("").Parse(srcServer)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/src/server.ts"); err != nil {
		return err
	}

	// src/server.spec.ts
	tmpl, err = template.New("").Parse(srcServerSpec)
	if err != nil {
		return err
	}
	if err = k.createFileFromTemplate(tmpl, "/src/server.spec.ts"); err != nil {
		return err
	}
	return nil
}
