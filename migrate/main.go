package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	"io"
	"os"
	"strings"

	"HRSystem/src/entity"
)

func main() {
	sb := &strings.Builder{}
	loadModels(sb)

	io.WriteString(os.Stdout, sb.String())
}

func loadModels(sb *strings.Builder) {
	models := []interface{}{
		&entity.Employee{},
		&entity.Request{},
		&entity.User{},
		&entity.Department{},
		&entity.Organization{},
	}

	stmts, err := gormschema.New("mysql").Load(models...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	sb.WriteString(stmts)
}
