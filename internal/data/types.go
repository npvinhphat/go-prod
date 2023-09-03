package data

type Item struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Levels      []string `yaml:"levels"`
	Link        string   `yaml:"link"`
	Stack       string
}

type Data struct {
	Command string
	Items   map[string][]Item
}
