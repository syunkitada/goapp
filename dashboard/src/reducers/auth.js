import { handleActions } from 'redux-actions';

import actions from '../actions'

const defaultState = {
  isSyncState: false,
  isFetching: false,
  redirectToReferrer: false,
  user: null,
  error: null,
};

export default handleActions({
  [actions.auth.syncState]: (state) => Object.assign({}, state, {
    isSyncState: true,
    isFetching: true,
  }),
  [actions.auth.login]: (state) => Object.assign({}, state, {
    isFetching: true,
  }),
  [actions.auth.loginSuccess]: (state, action) => Object.assign({}, state, {
    isFetching: false,
    redirectToReferrer: true,
    user: action.payload.user,
  }),
  [actions.auth.loginFailure]: (state, action) => Object.assign({}, defaultState, {
    isFetching: false,
    error: action.payload.error,
  }),
  [actions.auth.logout]: () => Object.assign({}, defaultState, {
    isFetching: true,
  }),
  [actions.auth.logoutSuccess]: () => Object.assign({}, defaultState, {
    isFetching: false,
  }),
  [actions.auth.logoutFailure]: () => Object.assign({}, defaultState, {
    isFetching: false,
  }),
}, defaultState);
