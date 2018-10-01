import { createActions } from 'redux-actions';

export default createActions({
  AUTH_SYNC_STATE: () => ({}),

  AUTH_LOGIN: (name, password) => ({name: name, password: password}),
  AUTH_LOGIN_SUCCESS: (user) => ({user: user}),
  AUTH_LOGIN_FAILURE: (error) => ({error: error}),

  AUTH_LOGOUT: (user) => ({user: user}),
  AUTH_LOGOUT_SUCCESS: () => ({}),
  AUTH_LOGOUT_FAILURE: () => ({}),
})
