package main

import (
	"encoding/csv"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/pflag"
)

const (
	tpl = `// Code generated DO NOT EDIT!

package {{.Package}}

var {{.Var}} = [][]{{.Type}} {
{{- range .Content }}
	{ {{- range .}}{{.}},{{end -}} },
{{- end}}
}
`
)

type layer struct {
	Package string
	Var     string
	Type    string
	Content [][]string
}

var (
	in = pflag.StringP("in", "i", "", "input directory")
)

func init() {
	pflag.Parse()
}

func main() {
	if in == nil || *in == "" {
		panic("in directory is missing: specify via `-i`")
	}

	dir, err := ioutil.ReadDir(*in)
	if err != nil {
		panic(err)
	}

	for _, info := range dir {
		if path.Ext(info.Name()) != ".csv" {
			continue
		}
		log.Println("generate level data:", info.Name())
		err = createLevel(path.Join(*in, info.Name()))
		if err != nil {
			panic(err)
		}
	}

}

func createLevel(file string) error {
	fileReader, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("file %q: %w", file, err)
	}
	defer func() {
		_ = fileReader.Close()
	}()

	data, err := parseCSV(fileReader)
	if err != nil {
		return err
	}

	out, err := writeGoFile(file, data)
	if err != nil {
		return err
	}

	err = gofmt(err, out)
	if err != nil {
		return err
	}

	return nil
}

func parseCSV(fileReader *os.File) ([][]string, error) {
	r := csv.NewReader(fileReader)

	var data [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read: %w", err)
		}

		data = append(data, record)
	}
	return data, nil
}

func writeGoFile(in string, data [][]string) (string, error) {
	fileName := path.Base(in)
	varName := fileName[strings.LastIndex(fileName, "_")+1 : len(fileName)-len(path.Ext(fileName))]
	varName = strings.ToLower(varName)
	packageName := fileName[:strings.LastIndex(fileName, "_")]
	out := path.Join(packageName, varName+".go")

	t := template.Must(template.New("layer").Parse(tpl))
	writer, err := os.Create(out)
	if err != nil {
		return "", fmt.Errorf("file %q: %w", out, err)
	}
	defer func() {
		_ = writer.Close()
	}()

	err = t.Execute(writer, layer{
		Package: packageName,
		Var:     strings.Title(varName),
		Type:    "int",
		Content: data,
	})
	if err != nil {
		return "", fmt.Errorf("execute template: %s", err)
	}

	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("write file: %s", err)
	}
	return out, err
}

func gofmt(err error, out string) error {
	code, err := ioutil.ReadFile(out)
	if err != nil {
		return fmt.Errorf("file %q: %w", out, err)
	}

	code, err = format.Source(code)
	if err != nil {
		return fmt.Errorf("go fmt %q: %s", out, err)
	}

	err = ioutil.WriteFile(out, code, 0666)
	if err != nil {
		return fmt.Errorf("file %q: %s", out, err)
	}

	return nil
}
