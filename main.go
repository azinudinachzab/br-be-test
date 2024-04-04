package main

import (
	"github.com/azinudinachzab/br-be-test/app"
	_ "github.com/lib/pq"
)

func main() {
	app.New().Run()
}
