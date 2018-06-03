// Package validate checks if the given name is an acceptable npm package name.
package validate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidatePackageName(t *testing.T) {
	tests := []struct {
		name   string
		result Result
	}{
		{"some-package", Result{ValidForNewPackages: true, ValidForOldPackages: true}},
		{"example.com", Result{ValidForNewPackages: true, ValidForOldPackages: true}},
		{"under_score", Result{ValidForNewPackages: true, ValidForOldPackages: true}},
		{"period.js", Result{ValidForNewPackages: true, ValidForOldPackages: true}},
		{"123numeric", Result{ValidForNewPackages: true, ValidForOldPackages: true}},
		{"carzy!", Result{ValidForNewPackages: false, ValidForOldPackages: true, Warnings: []string{"name can no longer contain special characters (\"~'!()*\")"}}},
		{"@npm/thingy", Result{ValidForNewPackages: true, ValidForOldPackages: true}},
		{"@npm-zors/money!time.js", Result{ValidForNewPackages: false, ValidForOldPackages: true, Warnings: []string{"name can no longer contain special characters (\"~'!()*\")"}}},
		{"", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"name cannot be empty"}}},
		{".start-with-period", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"name cannot start with a period"}}},
		{"_start-with-underscore", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"name cannot start with an underscore"}}},
		{"contains:colons", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"name can only contain URL-friendly characters"}}},
		{" leading-space", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"name cannot contain leading or trailing spaces", "name can only contain URL-friendly characters"}}},
		{"trailing-space ", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"name cannot contain leading or trailing spaces", "name can only contain URL-friendly characters"}}},
		{"s/l/a/s/h/e/s", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"name can only contain URL-friendly characters"}}},
		{"node_modules", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"node_modules is a blacklisted name"}}},
		{"favicon.ico", Result{ValidForNewPackages: false, ValidForOldPackages: false, Errors: []string{"favicon.ico is a blacklisted name"}}},
		{"http", Result{ValidForNewPackages: false, ValidForOldPackages: true, Warnings: []string{"http is a core module name"}}},
		{"ifyouwanttogetthesumoftwonumberswherethosetwonumbersarechosenbyfindingthelargestoftwooutofthreenumbersandsquaringthemwhichismultiplyingthembyitselfthenyoushouldinputthreenumbersintothisfunctionanditwilldothatforyou-", Result{ValidForNewPackages: false, ValidForOldPackages: true, Warnings: []string{"name can no longer contain more than 214 characters"}}},
		{"ifyouwanttogetthesumoftwonumberswherethosetwonumbersarechosenbyfindingthelargestoftwooutofthreenumbersandsquaringthemwhichismultiplyingthembyitselfthenyoushouldinputthreenumbersintothisfunctionanditwilldothatforyou", Result{ValidForNewPackages: true, ValidForOldPackages: true}},
		{"CAPITAL-LETTERS", Result{ValidForNewPackages: false, ValidForOldPackages: true, Warnings: []string{"name can no longer contain capital letters"}}},
	}
	for _, tc := range tests {
		result := Validate(tc.name)
		if !cmp.Equal(result, tc.result) {
			t.Errorf("returned result doesn't match the required result")
		}
	}
}
