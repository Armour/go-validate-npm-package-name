// Package validate checks if the given name is an acceptable npm package name.
package validate

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/armour/go-node-builtins/cmd/builtins"
)

// Result is the struct for the result of the validation.
type Result struct {
	ValidForNewPackages bool
	ValidForOldPackages bool
	Warnings            []string
	Errors              []string
}

const (
	scopedPackagePattern = `^(?:@([^/]+?)[/])?([^/]+?)$`
	specialCharacters    = `[\~\'\!\(\)\*]+`
)

var (
	scopedPackageRe     = regexp.MustCompile(scopedPackagePattern)
	specialCharactersRe = regexp.MustCompile(specialCharacters)
	blacklist           = []string{
		"node_modules",
		"favicon.ico",
	}
)

// Validate returns the result of whether the given name is an acceptable npm package name.
func Validate(name string) Result {
	var warnings []string
	var errors []string

	// Generate errors for stuff that not allowed

	if name == "" {
		errors = append(errors, "name cannot be empty")
		return done(warnings, errors)
	}

	if len(name) <= 0 {
		errors = append(errors, "name length must be greater than zero")
	}

	if strings.HasPrefix(name, ".") {
		errors = append(errors, "name cannot start with a period")
	}

	if strings.HasPrefix(name, "_") {
		errors = append(errors, "name cannot start with an underscore")
	}

	if strings.Trim(name, " ") != name {
		errors = append(errors, "name cannot contain leading or trailing spaces")
	}

	for _, b := range blacklist {
		if strings.ToLower(name) == b {
			errors = append(errors, b+" is a blacklisted name")
		}
	}

	// Generate warnings for stuff that used to be allowed

	coreModules, err := builtins.GetVersion("6.13.0")
	if err != nil {
		errors = append(errors, err.Error())
	}
	for _, m := range coreModules {
		if strings.ToLower(name) == m {
			warnings = append(warnings, m+" is a core module name")
		}
	}

	if len(name) > 214 {
		warnings = append(warnings, "name can no longer contain more than 214 characters")
	}

	if strings.ToLower(name) != name {
		warnings = append(warnings, "name can no longer contain capital letters")
	}

	if specialCharactersRe.ReplaceAllString(name, "") != name {
		warnings = append(warnings, "name can no longer contain special characters (\"~'!()*\")")
	}

	if urlSafe(name) != name {
		// Maybe it's a scoped package name, like @user/package
		nameMatch := scopedPackageRe.FindStringSubmatch(name)
		if len(nameMatch) > 2 {
			user := nameMatch[1]
			pkg := nameMatch[2]
			if urlSafe(user) == user && urlSafe(pkg) == pkg {
				return done(warnings, errors)
			}
		}
		errors = append(errors, "name can only contain URL-friendly characters")
	}

	return done(warnings, errors)
}

func done(warnings, errors []string) Result {
	return Result{
		ValidForNewPackages: len(errors) == 0 && len(warnings) == 0,
		ValidForOldPackages: len(errors) == 0,
		Warnings:            warnings,
		Errors:              errors,
	}
}

func urlSafe(str string) string {
	return strings.Replace(url.QueryEscape(str), "%21", "!", -1)
}
