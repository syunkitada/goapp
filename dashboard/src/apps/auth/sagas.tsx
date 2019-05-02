import {call, put, takeEvery} from 'redux-saga/effects';

import actions from '../../actions';
import modules from '../../modules';

function* syncState(action) {
  const {payload, error} = yield call(modules.auth.syncState);

  if (error) {
    yield put(actions.auth.authLoginFailure(error.message));
  } else if (payload.Err && payload.Err !== '') {
    yield put(actions.auth.authLoginFailure(payload.Err));
  } else {
    const user = {
      authority: payload.Authority,
      username: payload.Name,
    };
    yield put(actions.auth.authLoginSuccess(user));
  }
}

function* login(action) {
  const {payload, error} = yield call(modules.auth.login, action.payload);

  if (error) {
    yield put(actions.auth.authLoginFailure(error.message));
  } else if (payload.error && payload.error !== '') {
    yield put(actions.auth.authLoginFailure(payload.error));
  } else {
    const user = {
      authority: payload.Authority,
      username: payload.Name,
    };
    yield put(actions.auth.authLoginSuccess(user));
  }
}

function* logout(action) {
  const {error} = yield call(modules.auth.logout);

  if (error) {
    yield put(actions.auth.authLogoutFailure(error));
  } else {
    yield put(actions.auth.authLogoutSuccess());
  }
}

function* watchLogin() {
  yield takeEvery(actions.auth.authLogin, login);
}

function* watchSyncState() {
  yield takeEvery(actions.auth.authSyncState, syncState);
}

function* watchLogout() {
  yield takeEvery(actions.auth.authLogout, logout);
}

export default {
  watchLogin,
  watchLogout,
  watchSyncState,
};
