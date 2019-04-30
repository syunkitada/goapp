import fetch from 'cross-fetch';

function syncState() {
  const body = JSON.stringify({});

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/state', {
    method: 'POST',
    credentials: 'include',
    mode: 'cors',
    body: body,
  })
    .then(res => res.json())
    .then(function(payload) {
      return {
        payload: payload,
      };
    })
    .catch(function(error) {
      return {
        error: error,
      };
    });
}

function login({name, password}) {
  const body = JSON.stringify({
    username: name,
    password: password,
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/login', {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
    body: body,
  })
    .then(res => res.json())
    .then(function(payload) {
      return {
        payload: payload,
      };
    })
    .catch(function(error) {
      return {
        error: error,
      };
    });
}

function logout() {
  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/logout', {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
  })
    .then(res => res.json())
    .then(function(payload) {
      return {
        payload: payload,
      };
    })
    .catch(function(error) {
      return {
        error: error,
      };
    });
}

export default {
  syncState,
  login,
  logout,
};
