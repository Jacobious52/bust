package templatefile

import (
	"io"

	"text/template"

	"github.com/Jacobious52/bust/pkg/bustfile"
)

type TemplateFile struct {
	name       string
	targetFile string
	bustFile   *bustfile.BustFile
}

func NewTemplateFile(name, targetFile string, bustFile *bustfile.BustFile) *TemplateFile {
	return &TemplateFile{name, targetFile, bustFile}
}

func (t *TemplateFile) Apply(output io.Writer) error {
	templ := template.New(t.name)
	templ, err := templ.ParseFiles(t.targetFile)
	if err != nil {
		return err
	}
	err = templ.Execute(output, t.bustFile.Busts)
	return err
}
