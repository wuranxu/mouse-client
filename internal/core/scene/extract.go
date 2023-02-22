// extract params from response
// include jsonpath and regex, and so on

package scene

import (
	json "github.com/json-iterator/go"
	"github.com/oliveagle/jsonpath"
	"github.com/wuranxu/mouse-client/internal/entity"
)

type (
	FromType    string
	ExtractType string
)

const (
	JSONPath ExtractType = "JSONPath"
	Regex    ExtractType = "Regex"
)

const (
	Response       FromType = "Response"
	RequestHeader  FromType = "RequestHeader"
	ResponseHeader FromType = "ResponseHeader"
	StatusCode     FromType = "StatusCode"
)

type Extractor interface {
	Extract(source *entity.HTTPResponse) ([]byte, error)
}

type JSONPathExtractor struct {
	expression string
	from       FromType
}

func NewExtractor(out *Out) Extractor {
	switch out.ExtractType {
	case JSONPath:
		return NewJSONPathExtractor(out)
	default:
		return nil
	}
}

func From(source *entity.HTTPResponse, from FromType) any {
	switch from {
	case Response:
		var resp any
		err := json.Unmarshal(ToBytes(source.Data), &resp)
		if err != nil {
			return source.Data
		}
		return resp
	case RequestHeader:
		return source.Request.Headers
	case StatusCode:
		return source.StatusCode
	case ResponseHeader:
		return source.Headers
	default:
		return nil
	}
}

func (j *JSONPathExtractor) Extract(source *entity.HTTPResponse) (s []byte, err error) {
	src := From(source, j.from)
	if src == nil {
		return
	}
	lookup, err := jsonpath.JsonPathLookup(src, j.expression)
	if err != nil {
		return
	}
	return json.Marshal(lookup)
}

func NewJSONPathExtractor(out *Out) Extractor {
	return &JSONPathExtractor{
		expression: out.Expression,
		from:       out.From,
	}
}

type RegexExtractor struct {
	expression string
	from       FromType
}
