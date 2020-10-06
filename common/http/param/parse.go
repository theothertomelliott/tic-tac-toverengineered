package param

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Parse will extract the first value of the specified from an http request.
// The parameter will be parsed from json format and applied to the value pointer.
// String values may be accepted without surrounding quotes.
func Parse(req *http.Request, key string, value interface{}, options ParseOptions) error {
	params, ok := req.URL.Query()[key]
	var valueJSON []byte
	if !ok || len(params) == 0 {
		if options.Required {
			return fmt.Errorf("%q is required", key)
		}
		var err error
		valueJSON, err = json.Marshal(options.Default)
		if err != nil {
			return fmt.Errorf("default value: %w", err)
		}
	} else {
		valueJSON = []byte(params[0])
	}
	err := json.Unmarshal(valueJSON, value)
	if err != nil {
		// Allow for string-type inputs that aren't quote delimited
		if err2 := json.Unmarshal([]byte(fmt.Sprintf("\"%v\"", string(valueJSON))), value); err2 != nil {
			return fmt.Errorf("parsing %q: %w", key, err)
		}
	}
	return nil
}

// ParseOptions defines configuration options for parsing a parameter
type ParseOptions struct {
	Required bool
	Default  interface{}
}
