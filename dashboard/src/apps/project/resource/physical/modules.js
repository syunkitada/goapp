import fetch from 'cross-fetch'

function getIndex({projectName}) {
  const body = JSON.stringify({
    Action: {
      ProjectName: projectName,
      ServiceName: 'Resource',
      Name: 'GetPhysicalIndex',
      Data: '{}',
    },
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/Resource', {
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
  getIndex,
}
