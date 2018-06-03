package main

import (
	"fmt"

	"github.com/Armour/go-validate-npm-package-name/cmd/validate"
)

func main() {
	r := validate.Validate("go-validate-npm-package-name")
	fmt.Println(r.ValidForNewPackages, r.ValidForOldPackages)
	if len(r.Errors) != 0 {
		fmt.Println(r.Errors)
	}
	if len(r.Warnings) != 0 {
		fmt.Println(r.Warnings)
	}
}
