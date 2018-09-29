import { delay } from 'redux-saga'
import { put, call, takeEvery, all } from 'redux-saga/effects'
import actions from '../actions'
import api from '../modules/api'

function* syncState(action) {
  console.log("syncState", action.payload)
  const {payload, error} = yield call(api.syncState)

  if (error) {
    yield put(actions.auth.loginFailure(""))
  } else if (payload.error && payload.error != "") {
    yield put(actions.auth.loginFailure(""))
  } else {
    const user = {
      Name: payload.Name,
      Authority: payload.Authority,
    }
    yield put(actions.auth.loginSuccess(user))
  }
}

function* watchSyncState() {
  yield takeEvery(actions.auth.syncState, syncState)
}

function* login(action) {
  const {payload, error} = yield call(api.login, action.payload)

  if (error) {
    yield put(actions.auth.loginFailure(""))
  } else if (payload.error && payload.error != "") {
    yield put(actions.auth.loginFailure(""))
  } else {
    const user = {
      Name: payload.Username,
      Authority: payload.Authority,
    }
    yield put(actions.auth.loginSuccess(user))
  }
}

function* watchLogin() {
  yield takeEvery(actions.auth.login, login)
}

function* logout(action) {
  console.log("logout", action.payload)
  const {user, error} = yield call(api.logout, action.payload)

  if (error) {
    yield put(actions.auth.logoutFailure(error))
  } else {
    yield put(actions.auth.logoutSuccess(user))
  }
}

function* watchLogout() {
  yield takeEvery(actions.auth.logout, logout)
}

export default {
  watchSyncState,
  watchLogin,
  watchLogout,
}
