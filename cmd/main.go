package main

import (
	"fmt"

	"github.com/porrporporrpor/hydra-trust-issuer/cmd/cmds"
)

func main() {
	fmt.Printf(`service %s, built with Go %s`, cmds.Version, cmds.GoVersion)
	cmds.Execute()
}
