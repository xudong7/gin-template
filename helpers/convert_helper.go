package helpers

import (
	"encoding/json"
	"strconv"
)

func ConvertToInt(idValue interface{}) int {
	switch v := idValue.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	case string:
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		return intVal
	default:
		return 0
	}
}

func ConvertStringIdsToInt(data map[string]interface{}, idFields ...string) {
	// if len(idFields) == 0 {
	// 	idFields = []string{"id", "ID", "buyer_id", "seller_id", "item_id", "address_id"}
	// }

	for _, field := range idFields {
		if val, exists := data[field]; exists {
			data[field] = ConvertToInt(val)
		}
	}
}

func MapToStruct(data map[string]interface{}, target interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, target)
}
