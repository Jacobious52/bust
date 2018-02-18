package applyCommand

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jacobious52/bust/pkg/bustfile"
	"github.com/Jacobious52/bust/pkg/templatefile"
	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type applyCommand struct {
	bustFilePath    string
	targetFilePaths []string
	outputFilePaths []string
}

func Configure(app *kingpin.Application) {
	r := &applyCommand{}
	c := app.Command("apply", "applys a target file from a bust file to a output").
		PreAction(r.outputs).
		Action(r.run)

	c.Arg("targets", "location of target file to apply").
		Required().
		ExistingFilesVar(&r.targetFilePaths)

	c.Flag("bust", "location of the bust var file").
		Default("bust.yaml").
		StringVar(&r.bustFilePath)
}

func (r *applyCommand) outputs(c *kingpin.ParseContext) error {
	for _, filename := range r.targetFilePaths {
		if !strings.HasSuffix(filename, ".bust") {
			return fmt.Errorf("filename %v must in end .bust", filename)
		}
		r.outputFilePaths = append(r.outputFilePaths, filename[:len(filename)-len(".bust")])
	}
	return nil
}

func (r *applyCommand) run(c *kingpin.ParseContext) error {
	// load and read the bust file
	bustFileReader, err := os.Open(r.bustFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	bustFile, err := bustfile.NewBustFile(bustFileReader)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Debugf("%#v\n", bustFile.Busts)

	if len(r.targetFilePaths) != len(r.outputFilePaths) {
		log.Fatalf("targetfiles(%v) != outputfiles(%v)",
			len(r.targetFilePaths),
			len(r.outputFilePaths))
		return fmt.Errorf("Failed to parse target files")
	}

	for i := range r.targetFilePaths {
		renderTemplate(bustFile, r.targetFilePaths[i], r.outputFilePaths[i])
	}

	return nil
}

func renderTemplate(bustFile *bustfile.BustFile, targetFilePath, outputFilePath string) error {

	// create a new template and render
	targetFile := templatefile.NewTemplateFile(targetFilePath, targetFilePath, bustFile)

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = targetFile.Apply(outputFile)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
