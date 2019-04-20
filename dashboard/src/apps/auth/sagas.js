import { put, call, takeEvery } from 'redux-saga/effects'
import actions from '../../actions'
import modules from '../../modules'

function* syncState(action) {
  const {payload, error} = yield call(modules.auth.syncState)

  if (error) {
    console.log("DEBUG syncState", error)
    console.dir(error)
    yield put(actions.auth.authLoginFailure(error.message))
  } else if (payload.Err && payload.Err !== "") {
    console.log("DEBUG payload.Err")
    yield put(actions.auth.authLoginFailure(payload.Err))
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
    yield put(actions.auth.authLoginFailure(error.message))
  } else if (payload.error && payload.error !== "") {
    yield put(actions.auth.authLoginFailure(payload.error))
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
