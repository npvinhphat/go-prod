package data

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

const dataPath string = "data"

func LoadData(stacks []string) (map[string][]Item, error) {
	result := make(map[string][]Item)

	for _, stack := range stacks {
		filePath := filepath.Join(dataPath, stack+".yaml")
		content, err := loadYamlFile(filePath)
		if err != nil {
			return make(map[string][]Item), err
		}
		appendStackToItems(content, stack)

		mergo.Merge(&result, content, mergo.WithAppendSlice)
	}

	return result, nil
}

func FilterData(input map[string][]Item, level string) {
	for section, items := range input {
		var filteredItems []Item
		for _, item := range items {
			if contains(item.Levels, level) {
				filteredItems = append(filteredItems, item)
			}
		}
		input[section] = filteredItems
	}
}

func ApplyData(input map[string][]Item, tmpl string) (string, error) {
	funcMap := template.FuncMap{
		"contains": contains,
	}
	t, err := template.New("template").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return "", err
	}
	var output strings.Builder
	err = t.Execute(&output, input)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func contains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}

func appendStackToItems(content map[string][]Item, stack string) {
	// "default" is a generic stack, so we don't need to specify the name
	// It's also to avoid the stack name in the markdown file
	if stack == "default" {
		return
	}
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
