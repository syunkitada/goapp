import fetch from 'cross-fetch'

function syncState({projectName}) {
  const body = JSON.stringify({
    Action: {
      ProjectName: projectName,
      ServiceName: 'Resource',
      Name: 'GetState',
    },
  });
  console.log("DEBUGlalalala")
  console.log(body)

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/resource', {
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
