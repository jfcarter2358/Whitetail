package stream

type Stream struct {
	Schema    map[string]string      `yaml:"schema"`
	Probe     string                 `yaml:"probe"`
	Analyst   string                 `yaml:"analyst"`
	Arguments map[string]interface{} `yaml:"arguments"`
	Status    string                 `yaml:"status"`
}
