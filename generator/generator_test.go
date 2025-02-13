package generator

import (
	"fmt"
	"os"
	"testing"

	"github.com/YoungAgency/simplefix-go/utils"
)

var generator *Generator

func TestMain(m *testing.M) {
	var err error
	doc := &Doc{}
	if err = utils.ParseXML("./testdata/fix.4.4.xml", doc); err != nil {
		panic(fmt.Errorf("could not make XML document: %s", err))
	}

	config := &Config{}
	if err = utils.ParseXML("./testdata/types.xml", config); err != nil {
		panic(fmt.Errorf("could not make XML document: %s", err))
	}

	generator = NewGenerator(doc, config, "fix")

	m.Run()
	os.Exit(0)
}
