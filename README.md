# go-validate-npm-package-name

[![Go Report Card](https://goreportcard.com/badge/github.com/Armour/go-validate-npm-package-name)](https://goreportcard.com/report/github.com/Armour/go-validate-npm-package-name)
[![Go Project Layout](https://img.shields.io/badge/go-layout-blue.svg)](https://github.com/golang-standards/project-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/Armour/go-validate-npm-package-name/cmd/validate)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Template from jarvis](https://img.shields.io/badge/Hi-Jarvis-ff69b4.svg)](https://github.com/Armour/Jarvis)

## Install

```bash
go get github.com/Armour/go-validate-npm-package-name/cmd/validate
```

## Example

```go
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
```

## Contributing

1. Fork it!
1. Create your feature branch: `git checkout -b my-new-feature`
1. Commit your changes: `git commit -am 'Add some feature'`
1. Push to the branch: `git push origin my-new-feature`
1. Submit a pull request :D

## License

[MIT License](https://github.com/Armour/go-validate-npm-package-name/blob/master/LICENSE)
