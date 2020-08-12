package utils

import (
	"encoding/json"
	"smfbackend/models"
)

// ExtractProjectIds : extract project ids from project array
func ExtractProjectIds(projects []models.Project) []uint {
	ids := make([]uint, len(projects))
	for index, project := range projects {
		ids[index] = project.ID
	}
	return ids
}

func ExtractKeys(values []byte, keys []string) ([]byte, error) {
	var tempMap []map[string]interface{}
	json.Unmarshal(values, &tempMap)

	var respArr []map[string]interface{}
	for _, mapValue := range tempMap {
		respMap := make(map[string]interface{})
		for _, key := range keys {
			respMap[key] = mapValue[key]
		}
		respArr = append(respArr, respMap)
	}

	return json.Marshal(respArr)
}
