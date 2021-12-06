package unmarshal

import (
	"io"

	json "github.com/json-iterator/go"

	"github.com/metrico/loki-apache/pkg/logproto"
)

// DecodePushRequest directly decodes json to a logproto.PushRequest
func DecodePushRequest(b io.Reader, r *logproto.PushRequest) error {
	return json.NewDecoder(b).Decode(r)
}
