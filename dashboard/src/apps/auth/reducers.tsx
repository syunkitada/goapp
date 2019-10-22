import {reducerWithInitialState} from 'typescript-fsa-reducers';

import actions from '../../actions';

const defaultState = {
  error: null,
  isFetching: false,
  isSyncState: false,
  redirectToReferrer: false,
  user: null,
};

export default reducerWithInitialState(defaultState)
  .case(actions.auth.authLoginWithToken, state =>
    Object.assign({}, state, {
      isFetching: true,
      isSyncState: true,
    }),
  )
  .case(actions.auth.authLogin, state =>
    Object.assign({}, state, {
      isFetching: true,
    }),
  )
  .case(actions.auth.authLoginSuccess, (state, payload) =>
    Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
      user: payload,
    }),
  )
  .case(actions.auth.authLoginFailure, (state, payload) =>
    Object.assign({}, defaultState, {
      error: payload.error,
      isFetching: false,
    }),
  )
  .case(actions.auth.authLogout, state =>
    Object.assign({}, state, {
      isFetching: true,
    }),
  )
  .case(actions.auth.authLogoutSuccess, state =>
    Object.assign({}, defaultState, {
      isFetching: false,
    }),
  )
  .case(actions.auth.authLogoutFailure, (state, payload) =>
    Object.assign({}, state, {
      error: payload.error,
      isFetching: false,
    }),
  );
