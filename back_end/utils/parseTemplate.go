package utils

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 解析指定目录和其子文件夹下的html模板
func ParseTemplates(path string) *template.Template {
	tmpl := template.New("")
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = tmpl.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return tmpl
}
