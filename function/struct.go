package function

import "encoding/json"

// StructToMap struct转map
func StructToMap(s interface{}) (m map[string]string, err error) {
	var b []byte
	b, err = json.Marshal(s)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &m)
	return
}
