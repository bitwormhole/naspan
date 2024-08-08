package main

import (
	"os"

	"github.com/bitwormhole/naspan/modules/naspan"
	"github.com/starter-go/starter"
)

func main() {
	m := naspan.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
