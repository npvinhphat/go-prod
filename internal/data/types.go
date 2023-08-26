package data

type Item struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Levels      []string `yaml:"levels"`
	Stack       string
}
