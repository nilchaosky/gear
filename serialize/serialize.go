package serialize

import (
	"github.com/nilchaosky/gear/serialize/json"
	"github.com/nilchaosky/gear/serialize/protobuf"
)

var (
	Json     = json.NewSerializer()
	Protobuf = protobuf.NewSerializer()
)

type (

	// Marshaler represents a marshal interface
	Marshaler interface {
		Marshal(interface{}) ([]byte, error)
	}

	// Unmarshaler represents a Unmarshal interface
	Unmarshaler interface {
		Unmarshal([]byte, interface{}) error
	}

	// Serializer is the interface that groups the basic Marshal and Unmarshal methods.
	Serializer interface {
		Marshaler
		Unmarshaler
	}
)
