package main

import (
	"os"

	"github.com/bitwormhole/nasyun/modules/nasyun"
	"github.com/starter-go/starter"
)

func main() {
	m := nasyun.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
