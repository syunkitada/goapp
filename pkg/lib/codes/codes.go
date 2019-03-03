package codes

import "fmt"

const (
	// 2xx Success
	Ok int64 = 200

	// 4xx Client Errors
	ClientBadRequest     int64 = 400
	ClientInvalidRequest int64 = 401
	ClientNotFound       int64 = 404

	// 42x Client Request Data Errors
	ClientConflict      int64 = 420
	ClientAlreadyExists int64 = 422

	// 50x Server Internal Errors
	ServerInternalError int64 = 500

	// 52x Server Remote Errors
	RemoteTimeout      int64 = 520
	RemoteDbError      int64 = 521
	RemoteClusterError int64 = 525
)

func GetMsg(statusCode int64, data interface{}) string {
	switch statusCode {
	case ClientAlreadyExists:
		return fmt.Sprintf("AlreadyExists: %v", data)
	case ClientBadRequest:
		return fmt.Sprintf("BadRequest: %v", data)
	case RemoteDbError:
		return fmt.Sprintf("DbError: %v", data)
	}

	return fmt.Sprintf("Unknown: %v", data)
}
