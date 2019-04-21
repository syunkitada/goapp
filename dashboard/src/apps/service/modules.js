import fetch from 'cross-fetch'

import logger from '../../lib/logger'

function post({serviceName, actionName, projectName, queries}) {
  const body = JSON.stringify({
    Action: {
      ServiceName: serviceName,
      Name: actionName,
      ProjectName: projectName,
      Queries: queries,
    },
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/' + serviceName, {
    method: 'POST',
    credentials: 'include',
    mode: 'cors',
    body: body,
  }).then(res => res.json()).then(function(payload) {
    logger.info("modules", "post", queries, payload)
    return {
      payload: payload,
    };
  }).catch(function(error) {
    return {
      error: error
    };
  });
}

export default {
  post,
}
