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
  .case(actions.auth.authSyncState, state =>
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
      isFetching: false,
      // error: payload.action.payload.error,
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
      isFetching: false,
      // error: payload.action.payload.error,
    }),
  );

// export default handleActions(
//   {
//     [actions.auth.authSyncState]: state =>
//       Object.assign({}, state, {
//         isSyncState: true,
//         isFetching: true,
//       }),
//     [actions.auth.authLogin]: state =>
//       Object.assign({}, state, {
//         isFetching: true,
//       }),
//     [actions.auth.authLoginSuccess]: (state, action) =>
//       Object.assign({}, state, {
//         isFetching: false,
//         redirectToReferrer: true,
//         user: action.payload.user,
//       }),
//     [actions.auth.authLoginFailure]: (state, action) =>
//       Object.assign({}, defaultState, {
//         isFetching: false,
//         error: action.payload.error,
//       }),
//     [actions.auth.authLogout]: state =>
//       Object.assign({}, state, {
//         isFetching: true,
//       }),
//     [actions.auth.authLogoutSuccess]: (state, action) =>
//       Object.assign({}, defaultState, {
//         isFetching: false,
//       }),
//     [actions.auth.authLogoutFailure]: (state, action) =>
//       Object.assign({}, state, {
//         isFetching: false,
//         error: action.payload.error,
//       }),
//   },
//   defaultState,
// );
