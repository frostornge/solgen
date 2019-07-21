package proto

import (
	"bytes"
	"html/template"
	"io"
	"os"
)

const templatePath = "./proto/templates/*"

func render(writer io.Writer, c Contract) error {
	tmpl := template.Must(template.New("Bind").ParseGlob(templatePath))
	if err := tmpl.Execute(writer, c); err != nil {
		return err
	}
	return nil
}

func Render(c Contract) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := render(buffer, c); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func RenderFile(path string, c Contract) error {
	var out *os.File
	if _, err := os.Stat(path); os.IsNotExist(err) {
		out, err = os.Create(path)
		if err != nil {
			return err
		}
	} else {
		out, err = os.OpenFile(path, os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return render(out, c)
}
