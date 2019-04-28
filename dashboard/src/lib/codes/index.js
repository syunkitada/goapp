const OkCreated        = 201
const OkUpdated        = 202
const OkDeleted        = 203
const OkAccepted       = 210
const OkCreateAccepted = 211
const OkUpdateAccepted = 212
const OkDeleteAccepted = 213


export function toStringFromStatusCode(code) {
  switch(code) {
    case OkCreated:
      return "Created"
    case OkUpdated:
      return "Updated"
    case OkDeleted:
      return "Deleted"
    case OkAccepted:
      return "Accepted Request"
    case OkCreateAccepted:
      return "Accepted Create Request"
    case OkUpdateAccepted:
      return "Accepted Update Request"
    case OkDeleteAccepted:
      return "Accepted Delete Request"
  }
  return "Unknown"
}
