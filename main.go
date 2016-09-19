package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/fhs/go-netrc/netrc"
)

var (
	file     string
	tempFlag = flag.String("t", "{{.Login}}:{{.Password}}", "output template")
	pass     = flag.Bool("p", false, "output password")
)

func init() {
	file = path.Join(os.Getenv("HOME"), ".netrc")
}

func main() {
	flag.Parse()

	m, err := netrc.FindMachine(file, flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	if *pass {
		fmt.Println(m.Password)
		return
	}

	if t := flag.Arg(1); t != "" {
		*tempFlag = t
	}
	temp, err := template.New("temp").Parse(*tempFlag)
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(os.Stdout, m)
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [name] [template]\n", os.Args[0])
	flag.PrintDefaults()
}
