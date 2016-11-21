package daakia

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Renderer struct {
	BasePath string
}

func (r *Renderer) Render(dst *os.File, languages []string, services ...*Service) error {
	cleaned := make(map[string]bool)
	for _, service := range services {

		for _, lang := range languages {

			path := dst.Name() + "/" + lang + "/" + strings.ToLower(service.Namespace)
			if !cleaned[path] {
				err := Clean(path)
				if err != nil {
					return err
				}
				cleaned[path] = true
			}

			err := r.GenCode(path, "server", lang, service)
			if err != nil {
				return err
			}
			err = r.GenCode(path, "client", lang, service)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Renderer) GenCode(root, code_type, extension string, service *Service) error {
	file, err := MkFile(root + "/" + code_type, service.Name + "." + extension)
	if err != nil {
		return err
	}
	defer file.Close()
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"lower": strings.ToLower,
	}
	tmpl, err := template.New(extension + "." + code_type + ".tmpl").Funcs(funcMap).ParseFiles(r.BasePath + "/" + extension + "." + code_type + ".tmpl")
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, service)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	return nil
}
