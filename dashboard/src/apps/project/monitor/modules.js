import fetch from 'cross-fetch'

function syncState({projectName}) {
  const body = JSON.stringify({
    Action: {
      ProjectName: projectName,
      ServiceName: 'Monitor',
      Name: 'GetState',
    },
  });
  console.log("monitor.modules.syncState")
  console.log(body)

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/monitor', {
    method: 'POST',
    credentials: 'include',
    mode: 'cors',
    body: body,
  }).then(res => res.json()).then(function(payload) {
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
  syncState,
}
