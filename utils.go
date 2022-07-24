package neisgo

import "github.com/valyala/fastjson"

func getError(value *fastjson.Value) error {
	if !value.Exists("RESULT") {
		return nil
	}

	errResponse := Error{
		Code:    string(value.GetStringBytes("RESULT", "CODE")),
		Message: string(value.GetStringBytes("RESULT", "MESSAGE")),
	}

	switch errResponse.Code {
	case "ERROR-300":
		return ErrMissingValues
	case "ERROR-290":
		return ErrInvalidKey
	case "ERROR-310":
		return ErrServiceNotFound
	case "ERROR-333":
		return ErrInvalidLVType
	case "ERROR-336":
		return ErrTooManyData
	case "ERROR-337":
		return ErrTooManyRequest
	case "ERROR-500":
		return ErrServerError
	case "ERROR-600":
		return ErrDBConnection
	case "ERROR-601":
		return ErrSQLQuery
	case "INFO-300":
		return ErrKeyRestricted
	case "INFO-200":
		return ErrMissingData
	}
	return nil
}
