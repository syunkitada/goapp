package base_spec

type codeMap map[uint8]string

var CodeMap = codeMap{
	0: "Unknown",

	// 2x Success
	20: "Ok",
	21: "OkNotFound",
	22: "OkCreated",
	23: "OkUpdated",
	24: "OkDeleted",
	30: "OkAccepted",
	31: "OkCreateAccepted",
	32: "OkUpdateAccepted",
	33: "OkDeleteAccepted",

	// 10x Client Errors
	100: "ClientBadRequest",
	101: "ClientInvalidRequest",
	102: "ClientNotFound",
	103: "ClientInvalidAuth",

	// 11x Client Request Data Errors
	110: "ClientConflict",
	111: "ClientAlreadyExists",

	// 15x Server Internal Errors
	150: "ServerInternalError",

	// 16x Server Remote Errors
	161: "RemoteTimeout",
	162: "RemoteError",
}
