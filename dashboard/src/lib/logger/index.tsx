const IS_DEVELOP = true;

function debug(...args: any[]) {
  if (!IS_DEVELOP) {
    return;
  }

  const now = new Date();
  console.error(now.toISOString() + ' DEBUG', args);
}

function info(...args: any[]) {
  if (!IS_DEVELOP) {
    return;
  }

  const now = new Date();
  console.info(now.toISOString() + ' INFO', args);
}

function warn(...args: any[]) {
  if (!IS_DEVELOP) {
    return;
  }

  const now = new Date();
  console.info(now.toISOString() + ' WARN', args);
}

function error(...args: any[]) {
  if (!IS_DEVELOP) {
    return;
  }

  const now = new Date();
  console.error(now.toISOString() + ' ERROR', args);
}

export default {
  debug,
  error,
  info,
  warn,
};
