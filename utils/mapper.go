package utils

import "encoding/json"

func Mapper(objOne interface{}, objTwo interface{}) error {
	data, err := json.Marshal(objOne)
	if err != nil {
		return err
	}
	json.Unmarshal(data, objTwo)

	return nil
}
