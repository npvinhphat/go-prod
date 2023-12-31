package data

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"dario.cat/mergo"
	"github.com/npvinhphat/go-prod/assets"
	"gopkg.in/yaml.v3"
)

const stacksPath string = "stacks"

func LoadData(command string, stacks []string) (Data, error) {
	result := make(map[string][]Item)

	for _, stack := range stacks {
		filePath := filepath.Join(stacksPath, stack+".yaml")
		content, err := loadYamlFile(filePath)
		if err != nil {
			return Data{}, err
		}
		appendStackToItems(content, stack)

		mergo.Merge(&result, content, mergo.WithAppendSlice)
	}

	return Data{
		Command: command,
		Items:   result,
	}, nil
}

func FilterData(data Data, level string) {
	for section, items := range data.Items {
		var filteredItems []Item
		for _, item := range items {
			if contains(item.Levels, level) {
				filteredItems = append(filteredItems, item)
			}
		}
		data.Items[section] = filteredItems
	}
}

func ApplyData(data Data, tmpl string) (string, error) {
	funcMap := template.FuncMap{
		"contains":   contains,
		"formatName": formatName,
	}
	t, err := template.New("template").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return "", err
	}
	var output strings.Builder
	err = t.Execute(&output, data)
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

func formatName(item Item) string {
	result := item.Name
	if item.Stack != "" {
		result = fmt.Sprintf("(%s) %s", item.Stack, result)
	}
	if item.Link != "" {
		result = fmt.Sprintf("[%s](%s)", result, item.Link)
	}
	return result
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

	content, err := assets.Assets.ReadFile(filePath)
	if err != nil {
		return sectionsMap, err
	}

	err = yaml.Unmarshal(content, &sectionsMap)
	if err != nil {
		return sectionsMap, err
	}

	return sectionsMap, nil
}
