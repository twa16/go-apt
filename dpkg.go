package apt


import (
	"fmt"
	"os/exec"
	"bytes"
	"regexp"
)


const DPKG_CMD_NAME = "dpkg"

type Package struct {
	State string
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
	for _, line := range res {
		iPackage := Package{
			State: line[1],
			Name: line[2],
			Version: line[3],
			Architecture: line[4],
			Description: line[5],
		}
		packages = append(packages, iPackage)
		//fmt.Printf("[%d]: Got Install Package: %s, Version:%s\n", i, iPackage.Name, iPackage.Version)
	}
	return packages, nil

}