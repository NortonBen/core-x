package errors

import "encoding/json"

type Error struct {
	Message string
	Code int32
}

func New(massage string, code int32) Error {
	return Error{Message:massage, Code: code}
}

func (e Error) Error() string {
	data,_  := json.Marshal(e)
	return string(data)
}

func ToError(err error) (Error, error)  {
	var data Error
	err = json.Unmarshal([]byte(err.Error()), &data)
	if err != nil {
		return Error{}, nil
	}
	return data, nil
}
