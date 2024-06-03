package Streaming

import (
	"encoding/json"
)

func Valid(data []byte) bool {
	if json.Valid(data) {

		return true
	} else {
		return false
	}

}
