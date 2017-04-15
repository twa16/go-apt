package dpkg


import (
	"fmt"
	"os/exec"
	"bytes"
	"regexp"
)


const DPKG_CMD_NAME = "dpkg"

type Package struct {
	Name string
	Version string
	Architecture string
	Description string
}

func GetInstalledPackages() ([]Package, error) {
	//Get regex ready
	installedPackageRegexString := `(?m)(^[iuprh]\S)\s+(\S+)\s+(\S+)\s+(\S+)\s+(.+)`
	regex, _ := regexp.Compile(installedPackageRegexString)

	// dpkg command
	cmdName := "dpkg"
	cmdArgs := []string{"--list"}

	cmd := exec.Command(cmdName, cmdArgs...)

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput

	// Execute command
	err := cmd.Run() // will wait for command to return
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	// Only output the commands stdout
	var packages []Package
	res := regex.FindAllStringSubmatch(string(cmdOutput.Bytes()), -1)
	for i, line := range res {
		iPackage := Package{
			Name: line[1],
			Version: line[2],
			Architecture: line[3],
			Description: line[4],
		}
		packages = append(packages, iPackage)
		return nil, err
	}
	return packages, nil

}