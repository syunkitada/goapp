import { handleActions } from 'redux-actions';

import actions from '../../actions'

const defaultState = {
  isSyncState: false,
  isFetching: false,
  redirectToReferrer: false,
  error: null,
  user: null,
};

export default handleActions({
  [actions.auth.authSyncState]: (state) => Object.assign({}, state, {
    isSyncState: true,
    isFetching: true,
  }),
  [actions.auth.authLogin]: (state) => Object.assign({}, state, {
    isFetching: true,
  }),
  [actions.auth.authLoginSuccess]: (state, action) => Object.assign({}, state, {
    isFetching: false,
    redirectToReferrer: true,
    user: action.payload.user,
  }),
  [actions.auth.authLoginFailure]: (state, action) => Object.assign({}, defaultState, {
    isFetching: false,
    error: action.payload.error,
  }),
  [actions.auth.authLogout]: (state) => Object.assign({}, state, {
    isFetching: true,
  }),
  [actions.auth.authLogoutSuccess]: (state, action) => Object.assign({}, defaultState, {
    isFetching: false,
  }),
  [actions.auth.authLogoutFailure]: (state, action) => Object.assign({}, state, {
    isFetching: false,
    error: action.payload.error,
  }),
}, defaultState);
