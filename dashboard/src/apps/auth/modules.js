import fetch from 'cross-fetch'

function syncState() {
  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/state', {
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

function login({name, password}) {
  const body = JSON.stringify({
    username: name,
    password: password
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/login', {
    method: "POST",
    mode: 'cors',
    credentials: 'include',
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

function logout({name}) {
  const user = {
    name: name,
  }

  console.log("DEBUG: api logout")
  return {user: user, error: null}
}

export default {
  syncState,
  login,
  logout,
}
