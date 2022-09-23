package i18n

import (
	"encoding/json"
	"sigs.k8s.io/yaml"
)

// convert yaml bundle to json jsonData mapping
func yamlToJson(buf []byte) (any, error) {
	j, err := yaml.YAMLToJSON(buf)
	if err != nil {
		return nil, err
	}

	m := make(map[string]any)
	if err = json.Unmarshal(j, &m); err != nil {
		return nil, err
	}

	return m, nil
}
