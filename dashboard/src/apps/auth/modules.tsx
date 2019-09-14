import fetch from 'cross-fetch';

function loginWithToken() {
  const body = JSON.stringify({
    Queries: [
      {
        Data: JSON.stringify({}),
        Name: 'LoginWithToken',
      },
    ],
    Service: 'Auth',
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/q', {
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

function login({username, password}) {
  const body = JSON.stringify({
    Queries: [
      {
        Data: JSON.stringify({User: username, Password: password}),
        Name: 'Login',
      },
    ],
    Service: 'Auth',
  });

  console.log('debug login', username, password);

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/q', {
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
  const body = JSON.stringify({
    Queries: [
      {
        Data: JSON.stringify({}),
        Name: 'Logout',
      },
    ],
    Service: 'Auth',
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/q', {
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

export default {
  login,
  loginWithToken,
  logout,
};
