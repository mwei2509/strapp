package repo

import (
	"fmt"
	"os/exec"
	"strings"

	"sort"

	"github.com/masterminds/semver"
)

func CreateNewTag(increment string, message string) error {
	fmt.Println("hi", increment)

	cmd := "git tag | grep v | xargs"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err
	}
	tags := strings.Fields(string(out))

	vs := make([]*semver.Version, 0)
	for _, t := range tags {
		v, _ := semver.NewVersion(t)
		vs = append(vs, v)
	}
	sort.Sort(semver.Collection(vs))
	var latestVer *semver.Version
	if len(vs) == 0 {
		latestVer, _ = semver.NewVersion("v0.0")
		increment = "major"
	} else {
		latestVer = vs[len(vs)-1]
	}
	var nextVersion semver.Version
	switch increment {
	case "major":
		nextVersion = latestVer.IncMajor()
	case "minor":
		nextVersion = latestVer.IncMinor()
	case "patch":
		nextVersion = latestVer.IncPatch()
	}

	fmt.Println(nextVersion.Original())

	// create tag
	cmd = "git tag -a " + nextVersion.Original() + ` -m "` + message + `"`
	Log(cmd)
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err
	}
	Log(nextVersion.Original(), "tag created")

	// push tag
	cmd = "git push origin " + nextVersion.Original()
	Log(cmd)
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err
	}
	Log(nextVersion.Original(), "tag pushed")
	return nil
}
