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
  }).then(function(resp) {
    if (!resp.ok) {
      return resp.json().then(function(payload) {
        return {
          error: {
            errCode: resp.status,
            err: payload.Err,
          },
        };
      });
    }

    return resp.json().then(function(payload) {
      return {
        payload: payload,
        error: null,
      };
    });
  }).catch(function(error) {
    return {
      payload: null,
      error: {
        err: error,
      }
    };
  });
}

export default {
  post,
}
