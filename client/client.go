package client

import (
	"encoding/json"
	"io"
	"os"
)

func dump(v interface{}) {
	encodeJSON(os.Stdout, v)
}

func encodeJSON(w io.Writer, v interface{}) {
	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	enc.Encode(v)
}
