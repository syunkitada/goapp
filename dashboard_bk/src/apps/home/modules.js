import fetch from 'cross-fetch'

function syncState() {
  return fetch('https://192.168.10.103:8000/dashboard/state', {
    method: 'GET',
    credentials: 'include',
    mode: 'cors',
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
