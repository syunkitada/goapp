import { put, call, takeEvery } from 'redux-saga/effects'
import actions from '../../actions'
import modules from '../../modules'

function* syncState(action) {
  const {payload, error} = yield call(modules.auth.syncState)

  if (error) {
    yield put(actions.auth.authLoginFailure(""))
  } else if (payload.Err && payload.Err !== "") {
    yield put(actions.auth.authLoginFailure(""))
  } else {
    const user = {
      Name: payload.Name,
      Authority: payload.Authority,
    }
    yield put(actions.auth.authLoginSuccess(user))
  }
}

function* watchSyncState() {
  yield takeEvery(actions.auth.authSyncState, syncState)
}

function* login(action) {
  const {payload, error} = yield call(modules.auth.login, action.payload)

  if (error) {
    yield put(actions.auth.authLoginFailure(""))
  } else if (payload.error && payload.error !== "") {
    yield put(actions.auth.authLoginFailure(""))
  } else {
    const user = {
      Name: payload.Name,
      Authority: payload.Authority,
    }
    yield put(actions.auth.authLoginSuccess(user))
  }
}

function* watchLogin() {
  yield takeEvery(actions.auth.authLogin, login)
}

function* logout(action) {
  console.log("logout", action.payload)
  const {payload, error} = yield call(modules.auth.logout)
  console.log(payload)

  if (error) {
    yield put(actions.auth.authLogoutFailure(error))
  } else {
    yield put(actions.auth.authLogoutSuccess())
  }
}

function* watchLogout() {
  yield takeEvery(actions.auth.authLogout, logout)
}

export default {
  watchSyncState,
  watchLogin,
  watchLogout,
}
