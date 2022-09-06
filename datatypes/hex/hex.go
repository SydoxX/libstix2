package hex

import (
	"encoding/hex"
	"encoding/json"
)

type Hex []byte

func FromString(hexstring string) (Hex, error) {
	return hex.DecodeString(hexstring)
}

func (h Hex) MarshalJSON() ([]byte, error) {
	if h == nil {
		return []byte("null"), nil
	}
	return json.Marshal(hex.EncodeToString(h))
}

func (h *Hex) UnmarshalJSON(data []byte) error {
	var str *string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str == nil {
		*h = nil
		return nil
	}

	decoded, err := FromString(*str)
	if err != nil {
		return err
	}
	*h = decoded
	return nil
}
