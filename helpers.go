package gonvs

import "encoding/json"

func wrapJSONParam(value string) json.RawMessage {
	return json.RawMessage(`"` + value + `"`)
}
