package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/YoungAgency/simplefix-go/generator"
	"github.com/YoungAgency/simplefix-go/utils"
)

func main() {
	var err error

	outputDir := flag.String("o", "./fix44/", "output directory")
	typesMappingPath := flag.String("t", "./source/types.xml", "path to XML file with types mapping")
	sourceXMLPath := flag.String("s", "./source/fix44.xml", "path to main XML file")

	flag.Parse()

	doc := &generator.Doc{}
	if err = utils.ParseXML(*sourceXMLPath, doc); err != nil {
		panic(fmt.Errorf("could not make Doc XML: %s", err))
	}

	config := &generator.Config{}
	if err = utils.ParseXML(*typesMappingPath, config); err != nil {
		panic(fmt.Errorf("could not make Doc XML: %s", err))
	}

	g := generator.NewGenerator(doc, config, filepath.Base(*outputDir))

	err = os.MkdirAll(*outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = g.Execute(*outputDir)
	if err != nil {
		panic(err)
	}
}
