package main

import (
	"fmt"
	"os"
	"path/filepath"
	"flag"
	"github.com/daakia/daakia"
	"errors"
	"log"
)

type File struct {
	file *os.File
}

func (i *File) String() string {
	return fmt.Sprintf("In File: %v", i.file)
}

func (i *File) Set(path string) error {
	if path == "" {
		return errors.New(fmt.Sprintf("Input path %s can not be empty", path))
	}
	real_path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	// Try to open the file
	i.file, err = os.Open(real_path)
	if err != nil {
		return err
	}
	stat, err := i.file.Stat()
	if err != nil {
		return err
	}
	if !stat.Mode().IsRegular() {
		return errors.New(fmt.Sprintf("%s should be a file",i.file.Name()))
	}
	return nil
}

func (i *File) File() *os.File {
	return i.file
}



func main() {
	var in File
	var golang, js bool
	var out string
	flag.Var(&in, "i", "Path to the input file")
	flag.StringVar(&out, "o", "./out", "Path where generated code should be written to")
	flag.BoolVar(&golang, "go", false, "Output in go?")
	flag.BoolVar(&js, "js", false, "Output in js?")
	flag.Parse()
	var InFile = in.File()
	if InFile == nil {
		flag.Usage()
		log.Fatal("No input file specified")
	}
	out_langs := make([]string,0,10)
	if golang {
		out_langs = append(out_langs,"go")
	}
	if js {
		out_langs = append(out_langs,"js")
	}

	if len(out_langs) == 0 {
		flag.Usage()
		log.Fatal("No Output languages specified\n")
	}
	out_path,err := filepath.Abs(out)
	if err != nil {
		log.Fatal(err)
	}
	out_dir, err := os.Open(out_path)
	if os.IsNotExist(err) {
		err = daakia.Mkdir(out_path)
		out_dir, err = os.Open(out_path)
	}
	if err !=nil {
		log.Fatal(err)
	}
	var daakia_template_dir string
	daakia_template_dir = os.Getenv("DAAKIA_TEMPLATE_DIR")
	if daakia_template_dir == "" {
		daakia_template_dir = os.Getenv("GOPATH")+"/src/github.com/daakia/daakia/templates"
		fmt.Println("environment variable $DAAKIA_TEMPLTE_DIR not set, using default", daakia_template_dir)
	}
	base_path, err := filepath.Abs(daakia_template_dir)
	if err != nil {
		log.Fatal(err)
	}
	services, err := daakia.ParseToml(InFile,11)
	if err != nil {
		log.Fatal(err)
	}
	renderer := &daakia.Renderer{
		BasePath: base_path,
	}
	err = renderer.Render(out_dir, out_langs, services...)
	if err != nil {
		log.Fatal(err)
	}
}
