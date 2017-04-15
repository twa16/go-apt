package main

import (
	"github.com/twa16/go-dpkg"
	"fmt"
)

func main() {
	pkgs, _ := dpkg.GetInstalledPackages()
	for i, pkg := range pkgs {
		fmt.Printf("[%d]: Got Install Package: %s, Version:%s\n", i, pkg.Name, pkg.Version)
	}
}
