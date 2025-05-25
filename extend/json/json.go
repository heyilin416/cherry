package cherryJson

import (
	"os"

	jsoniter "github.com/json-iterator/go"
)

func ToJson(i interface{}) string {
	if i == nil {
		return ""
	}

	bytes, err := jsoniter.Marshal(i)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func DeepCopy(in, out interface{}) error {
	b, _ := jsoniter.Marshal(in)
	return jsoniter.Unmarshal(b, out)
}

func ReadMaps(path string, maps map[string]interface{}) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = jsoniter.Unmarshal(bytes, &maps)
	if err != nil {
		return err
	}
	return nil
}
