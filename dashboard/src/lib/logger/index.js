var IS_DEVELOP = true;

function debug(data) {
  if (!IS_DEVELOP) {
    return
  }

  let now = new Date()
  console.error(now.toISOString() + " DEBUG", arguments)
}

function info() {
  if (!IS_DEVELOP) {
    return
  }

  let now = new Date()
  console.info(now.toISOString() + " INFO", arguments)
}

function warn() {
  if (!IS_DEVELOP) {
    return
  }

  let now = new Date()
  console.info(now.toISOString() + " WARN", arguments)
}

function error(data) {
  if (!IS_DEVELOP) {
    return
  }

  let now = new Date()
  console.error(now.toISOString() + " ERROR", arguments)
}

export default {
  info,
  error,
}
