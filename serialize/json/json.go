package json

import (
	jsoniter "github.com/json-iterator/go"
)

// Serializer implements the serialize.Serializer interface
type Serializer struct{}

// NewSerializer returns a new Serializer.
func NewSerializer() *Serializer {
	return &Serializer{}
}

// Marshal returns the JSON encoding of v.
func (s *Serializer) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v)
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func (s *Serializer) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, v)
}
