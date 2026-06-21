package helpers

import (
	"bytes"
	"encoding/json"
)

// JSONPretty returns formatted JSON string for debugging
func JSONDebuger[T any](v T) string {
	jsonBytes, _ := json.MarshalIndent(v, "", "  ")
	return string(jsonBytes)
}

// JSONDebuggerClean returns formatted JSON string with empty/nil values removed
// Generic type T ensures type awareness for IDE autocomplete
// Preserves original key order from JSON
func JSONDebuggerClean[T any](v T) string {
	b, _ := json.Marshal(v)

	// Parse and filter while preserving key order
	filtered, _ := filterJSONPreserveOrder(b)

	// Format with indentation
	var out bytes.Buffer
	json.Indent(&out, filtered, "", "  ")
	return out.String()
}

// filterJSONPreserveOrder iterates through JSON using Decoder to preserve key order
// Completely rebuilds JSON without unmarshaling to maps (which causes sorting)
func filterJSONPreserveOrder(jsonBytes []byte) ([]byte, error) {
	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))

	// Read first token to know if it's object or array
	t, err := decoder.Token()
	if err != nil {
		return jsonBytes, err
	}

	var buf bytes.Buffer

	// Handle objects
	if t == json.Delim('{') {
		buf.WriteString("{")
		first := true

		for decoder.More() {
			// Read key token
			keyToken, _ := decoder.Token()
			key := keyToken.(string)

			// Read value as RawMessage to preserve nested structure
			var rawVal json.RawMessage
			decoder.Decode(&rawVal)

			// Check if value is empty by decoding to any
			var tmpVal any
			json.Unmarshal(rawVal, &tmpVal)
			if isEmptyJSONValue(tmpVal) {
				continue
			}

			// Add comma if not first
			if !first {
				buf.WriteString(",")
			}
			first = false

			// Write key
			keyJSON, _ := json.Marshal(key)
			buf.Write(keyJSON)
			buf.WriteString(":")

			// Recursively filter nested value (this preserves order since rawVal is original JSON)
			filtered, _ := filterJSONPreserveOrder(rawVal)
			buf.Write(filtered)
		}

		buf.WriteString("}")
		return buf.Bytes(), nil
	}

	// Handle arrays
	if t == json.Delim('[') {
		buf.WriteString("[")
		first := true

		for decoder.More() {
			// Read value as RawMessage to preserve structure
			var rawVal json.RawMessage
			decoder.Decode(&rawVal)

			// Check if value is empty
			var tmpVal any
			json.Unmarshal(rawVal, &tmpVal)
			if isEmptyJSONValue(tmpVal) {
				continue
			}

			if !first {
				buf.WriteString(",")
			}
			first = false

			// Recursively filter nested value
			filtered, _ := filterJSONPreserveOrder(rawVal)
			buf.Write(filtered)
		}

		buf.WriteString("]")
		return buf.Bytes(), nil
	}

	// Scalar value - return as is
	return jsonBytes, nil
}

func isEmptyJSONValue(v any) bool {
	switch val := v.(type) {
	case nil:
		return true
	case string:
		return val == ""
	case map[string]any:
		return len(val) == 0
	case []any:
		return len(val) == 0
	default:
		return false
	}
}
