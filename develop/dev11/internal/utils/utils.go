package utils

import (
	"encoding/json"
)

func SerializingToJSON(v interface{}) ([]byte, error) {
	js, err := json.Marshal(v)
	if err != nil {
		return js, err
	}
	return js, nil
}

func ParseEvent() {

}

func ValidateEvent() {

}
