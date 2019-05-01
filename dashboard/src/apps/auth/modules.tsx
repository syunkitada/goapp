import fetch from 'cross-fetch';

function syncState() {
  const body = JSON.stringify({});

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/state', {
    body,
    credentials: 'include',
    method: 'POST',
    mode: 'cors',
  })
    .then(res => res.json())
    .then(payload => {
      return {
        payload,
      };
    })
    .catch(error => {
      return {error};
    });
}

function login({username, password}) {
  const body = JSON.stringify({
    password,
    username,
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/login', {
    body,
    credentials: 'include',
    method: 'POST',
    mode: 'cors',
  })
    .then(res => res.json())
    .then(payload => {
      return {payload};
    })
    .catch(error => {
      return {error};
    });
}

function logout() {
  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/dashboard/logout', {
    credentials: 'include',
    method: 'POST',
    mode: 'cors',
  })
    .then(res => res.json())
    .then(payload => {
      return {payload};
    })
    .catch(error => {
      return {error};
    });
}

export default {
  login,
  logout,
  syncState,
};
