import fetch from 'cross-fetch'

function syncState() {
  const body = JSON.stringify({});
  console.log("auth.modules.syncState")

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/state', {
    method: 'POST',
    credentials: 'include',
    mode: 'cors',
    body: body,
  }).then(res => res.json()).then(function(payload) {
    console.log("auth.modules.syncState payload")
    console.log(payload)
    return {
      payload: payload,
    };
  }).catch(function(error) {
    console.log("auth.modules.syncState error")
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
  console.log("auth.modules.login")

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

function logout() {
  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/logout', {
    method: "POST",
    mode: 'cors',
    credentials: 'include',
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
  login,
  logout,
}
