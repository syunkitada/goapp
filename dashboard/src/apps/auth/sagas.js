import { delay } from 'redux-saga'
import { put, call, takeEvery, all } from 'redux-saga/effects'
import actions from '../../actions'
import modules from '../../modules'

function* syncState(action) {
  console.log("auth: syncState", action.payload)
  const {payload, error} = yield call(modules.auth.syncState)

  if (error) {
    yield put(actions.auth.authLoginFailure(""))
  } else if (payload.error && payload.error != "") {
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
  } else if (payload.error && payload.error != "") {
    yield put(actions.auth.authLoginFailure(""))
  } else {
    const user = {
      Name: payload.Username,
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
  const {user, error} = yield call(modules.auth.authLogout, action.payload)

  if (error) {
    yield put(actions.auth.authLogoutFailure(error))
  } else {
    yield put(actions.auth.authLogoutSuccess(user))
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
