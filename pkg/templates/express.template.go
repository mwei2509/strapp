package templates

import (
	"os"
	"text/template"
)

var tempVanillaAppJs string = `
import express from "express";
import router from "./routes";

const app = express();
const port = {{.Port}};

app.use(express.json());
app.use(router);

app.listen(port, () => console.log(` + "`" + `application started at ${{"{"}}{{.Port}}{{"}"}}` + "`" + `));
`

// if i need to create console.log(`something ${whatever}`)
// `console.log(` + "`" + `something ${{"{"}}{{.Whatever}}{{"}"}}` + "`" + `);`
type AppJs struct {
	Port int64
}

func CreateVanillaAppJs(directory string, txt AppJs) error {
	tmpl, err := template.New("").Parse(tempVanillaAppJs)
	if err != nil {
		return err
	}

	file, err := os.Create(directory + "/app.js")
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, txt)
	if err != nil {
		return err
	}
	return nil
}
