package apt

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

type PackageListing struct {
	Name string
	Description string
}

func GetAPTPackageNameList() ([]PackageListing, error) {
	//Get regex ready
	//apt list regex
	packageListingRegexString := `(^\S+) - (.+)`
	regex, _ := regexp.Compile(packageListingRegexString)

	// dpkg command
	cmdName := "apt-cache"
	cmdArgs := []string{"search", "."}

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
	var packages []PackageListing
	res := regex.FindAllStringSubmatch(string(cmdOutput.Bytes()), -1)
	for _, line := range res {
		lPackage := PackageListing{
			Name: line[1],
			Description: line[2],
		}
		packages = append(packages, lPackage)
		//fmt.Printf("[%d]: Got Install Package: %s, Version:%s\n", i, iPackage.Name, iPackage.Version)
	}
	return packages, nil
}
