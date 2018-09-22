import { createActions } from 'redux-actions';

export default createActions({
  SYNC_STATE: () => ({}),

  LOGIN: (name, password) => ({name: name, password: password}),
  LOGIN_SUCCESS: (user) => ({user: user}),
  LOGIN_FAILURE: (error) => ({error: error}),

  LOGOUT: (user) => ({user: user}),
  LOGOUT_SUCCESS: () => ({}),
  LOGOUT_FAILURE: () => ({}),
})
