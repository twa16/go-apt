package main

import (
	"fmt"
	"github.com/twa16/go-apt"
)

func main() {
	pkgs, err := apt.GetInstalledPackages()
	if err != nil {
		fmt.Println("Error: "+err.Error())
	}
	for i, pkg := range pkgs {
		fmt.Printf("[%d]: Got Installed Package: %s, Version:%s\n", i, pkg.Name, pkg.Version)
	}

	pkgListing, err := apt.GetAPTPackageNameList()
	for i, pkg := range pkgListing {
		fmt.Printf("[%d]: Got Available Package: %s, Description:%s\n", i, pkg.Name, pkg.Description)
	}
}
