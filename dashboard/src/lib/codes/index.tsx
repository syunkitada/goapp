// This code is auto generated.
// Don't modify this code.
const Unknown = 0;
const Ok = 20;
const OkCreated = 21;
const OkUpdated = 22;
const OkDeleted = 23;
const OkAccepted = 30;
const OkCreateAccepted = 31;
const OkUpdateAccepted = 32;
const OkDeleteAccepted = 33;
const ClientBadRequest = 100;
const ClientInvalidRequest = 101;
const ClientNotFound = 102;
const ClientInvalidAuth = 103;
const ClientConflict = 110;
const ClientAlreadyExists = 111;
const ServerInternalError = 150;
const RemoteTimeout = 161;
const RemoteError = 162;

export function toStringFromStatusCode(code) {
  switch(code) {
    case Unknown:
      return "Unknown";
    case Ok:
      return "Ok";
    case OkCreated:
      return "OkCreated";
    case OkUpdated:
      return "OkUpdated";
    case OkDeleted:
      return "OkDeleted";
    case OkAccepted:
      return "OkAccepted";
    case OkCreateAccepted:
      return "OkCreateAccepted";
    case OkUpdateAccepted:
      return "OkUpdateAccepted";
    case OkDeleteAccepted:
      return "OkDeleteAccepted";
    case ClientBadRequest:
      return "ClientBadRequest";
    case ClientInvalidRequest:
      return "ClientInvalidRequest";
    case ClientNotFound:
      return "ClientNotFound";
    case ClientInvalidAuth:
      return "ClientInvalidAuth";
    case ClientConflict:
      return "ClientConflict";
    case ClientAlreadyExists:
      return "ClientAlreadyExists";
    case ServerInternalError:
      return "ServerInternalError";
    case RemoteTimeout:
      return "RemoteTimeout";
    case RemoteError:
      return "RemoteError";
  }
  return "Unknown";
}
