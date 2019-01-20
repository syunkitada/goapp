import fetch from 'cross-fetch'

function syncState({projectName}) {
  const body = JSON.stringify({
    Action: {
      ProjectName: projectName,
      ServiceName: 'Monitor',
      Name: 'GetUserState',
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

function syncIndexState({projectName, indexName}) {
  const body = JSON.stringify({
    Action: {
      ProjectName: projectName,
      ServiceName: 'Monitor',
      Name: 'GetIndexState',
      Data: '',
    },
  });
  console.log("monitor.modules.syncIndexState")
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
  syncIndexState,
}
