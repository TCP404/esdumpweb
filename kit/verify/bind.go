package verify

import (
	"github.com/gin-gonic/gin/binding"
)

// BindingDefault returns the appropriate BindingBody instance based on the content type.
func BindingDefault(contentType string) binding.BindingBody {
	switch contentType {
	case binding.MIMEJSON:
		return binding.JSON
	case binding.MIMEXML, binding.MIMEXML2:
		return binding.XML
	case binding.MIMEPROTOBUF:
		return binding.ProtoBuf
	case binding.MIMEMSGPACK, binding.MIMEMSGPACK2:
		return binding.MsgPack
	case binding.MIMEYAML:
		return binding.YAML
	case binding.MIMETOML:
		return binding.TOML
	default: // case MIMEPOSTForm:
		return binding.JSON
	}
}
