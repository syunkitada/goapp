package base_const

const (
	// 1x Unknown
	CodeUnknown uint8 = 10

	// 2x Success
	CodeOkRead           uint8 = 20
	CodeOkCreated        uint8 = 21
	CodeOkUpdated        uint8 = 22
	CodeOkDeleted        uint8 = 23
	CodeOkAccepted       uint8 = 30
	CodeOkCreateAccepted uint8 = 31
	CodeOkUpdateAccepted uint8 = 32
	CodeOkDeleteAccepted uint8 = 33

	// 10x Client Errors
	CodeClientBadRequest     uint8 = 100
	CodeClientInvalidRequest uint8 = 101
	CodeClientNotFound       uint8 = 102
	CodeClientInvalidAuth    uint8 = 103

	// 11x Client Request Data Errors
	CodeClientConflict      uint8 = 110
	CodeClientAlreadyExists uint8 = 111

	// 15x Server Internal Errors
	CodeServerInternalError uint8 = 150

	// 16x Server Remote Errors
	CodeRemoteTimeout uint8 = 161
	CodeRemoteError   uint8 = 162
)
