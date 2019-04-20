var IS_DEVELOP = true;

function info(data) {
  if (!IS_DEVELOP) {
    return
  }

  let now = new Date()
  console.info(now.toISOString() + " INFO", data)
}

function error(data) {
  if (!IS_DEVELOP) {
    return
  }

  let now = new Date()
  console.error(now.toISOString() + " ERROR", data)
}

export default {
  info,
  error,
}
