package codes

const (
	// 2xx Success
	Ok int64 = 200

	// 4xx Client Errors
	ClientBadRequest     int64 = 400
	ClientNotFound       int64 = 404
	ClientInvalidRequest int64 = 406
	ClientConflict       int64 = 409

	// 50x Server Internal Errors
	ServerInternalError int64 = 500

	// 52x Server Remote Errors
	RemoteTimeout int64 = 520
	RemoteDbError int64 = 521
)
