package main

import (
	"fmt"
	"os"

	"github.com/onflow/cadence/runtime/cmd"
	"github.com/onflow/cadence/runtime/common"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	code, err := os.ReadFile("example/Test.cdc")
	if err != nil {
		return err
	}

	memberAccountAccess := map[common.Location]map[common.Location]struct{}{}
	codes := map[common.Location][]byte{}
	loc := common.NewScriptLocation(nil, []byte("example/Test.cdc"))

	program, must := cmd.PrepareProgram(code, loc, codes)
	fmt.Printf("%+v\n", program)

	c, _ := cmd.PrepareChecker(program, loc, codes, memberAccountAccess, must)
	err = c.Check()
	if err != nil {
		return err
	}

	return nil
}
