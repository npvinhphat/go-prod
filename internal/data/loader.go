package data

import (
	"os"
	"path/filepath"
	"text/template"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

var tmpl = `
{{- define "table" -}}
| Name | Description | Basic | Intermediate | Advanced |
|------|-------------|-------| ------------ | -------- |
{{ range . }}| {{ .Name }} | {{ .Description }} | {{ if contains .Levels "basic" }}✅{{ end }} | {{ if contains .Levels "intermediate" }}✅{{ end }} | {{ if contains .Levels "advanced"}}✅{{ end }} |
{{ end }}
{{- end -}}
# Production Readiness Checklist

{{- if .overview }}
## Overview

{{ template "table" .overview }}
{{- end }}
{{- if .observability }}
## Observability

{{ template "table" .observability }}
{{- end }}`

var dataPath = filepath.Join("data")

func LoadData(stacks []string) (map[string][]Item, error) {
	sections := make(map[string][]Item)

	// Load all the stacks content first
	for _, stack := range stacks {
		filePath := filepath.Join(dataPath, stack+".yaml")
		content, err := loadYamlFile(filePath)
		if err != nil {
			return make(map[string][]Item), err
		}
		appendStackToItems(content, stack)

		mergo.Merge(&sections, content, mergo.WithAppendSlice)
	}
	// Templating
	funcMap := template.FuncMap{
		"contains": func(list []string, target string) bool {
			for _, s := range list {
				if s == target {
					return true
				}
			}
			return false
		},
	}
	t, err := template.New("tableTemplate").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return make(map[string][]Item), err
	}
	err = t.Execute(os.Stdout, sections)
	if err != nil {
		return make(map[string][]Item), err
	}

	return sections, nil
}
func appendStackToItems(content map[string][]Item, stack string) {
	for _, items := range content {
		for i := range items {
			items[i].Stack = stack
		}
	}
}

func loadYamlFile(filePath string) (map[string][]Item, error) {
	sectionsMap := make(map[string][]Item)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return sectionsMap, err
	}

	err = yaml.Unmarshal(content, &sectionsMap)
	if err != nil {
		return sectionsMap, err
	}

	return sectionsMap, nil
}
