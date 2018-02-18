package bustfile

import (
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type BustFile struct {
	Busts map[string]interface{} `yaml:"busts,omitempty"`
}

func NewBustFile(reader io.Reader) (*BustFile, error) {
	var bustFile *BustFile

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return bustFile, err
	}

	err = yaml.Unmarshal(b, &bustFile)
	return bustFile, err
}
