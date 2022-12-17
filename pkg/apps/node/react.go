package node

// "github.com/mwei2509/strapp/pkg/templates"

/* REACT */
// func (n *Node) installReact() error {
// 	cmd := exec.Command("npm", "install", "react")
// 	cmd.Dir = n.Directory
// 	out, err := cmd.Output()
// 	if err != nil {
// 		return err
// 	}
// 	n.Log(fmt.Sprintf("%s\n", out))
// 	return nil
// }

func (n *Node) setupReact() error {
	// n.Log("initializing with create-react-app")
	// cmd := exec.Command("npx", "create-react-app", ".", "--template", "redux-typescript")
	// cmd.Dir = n.Directory
	// out, err := cmd.Output()
	// if err != nil {
	// 	return err
	// }
	// n.Log(fmt.Sprintf("%s\n", out))
	// uncomment above!
	// !

	// npx create-react-app my-app --template redux-typescript
	// if err := n.installReact(); err != nil {
	// 	return err
	// }

	// appjs := templates.AppJs{Port: 3000}
	// if err := templates.CreateVanillaAppJs(n.Directory, appjs); err != nil {
	// 	return err
	// }
	return nil

}
