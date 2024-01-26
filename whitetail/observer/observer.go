package observer

import "whitetail/stream"

type Observer struct {
	Streams map[string]stream.Stream `yaml:"streams"`
}
