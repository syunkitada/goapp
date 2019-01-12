import { delay } from 'redux-saga'
import { put, call, takeEvery, all } from 'redux-saga/effects'
import actions from '../../../actions'
import modules from '../../../modules'

function* syncState(action) {
  console.log("monitor: syncState", action.payload)
  const {payload, error} = yield call(modules.monitor.syncState, action.payload)

  console.log(payload)
  console.log(error)

  if (error) {
    yield put(actions.monitor.monitorSyncStateFailure(""))
  } else if (payload.error && payload.error != "") {
    yield put(actions.monitor.monitorSyncStateFailure(""))
  } else {
    const monitor = {
      payload: payload,
    }
    yield put(actions.monitor.monitorSyncStateSuccess(monitor))
  }
}

function* watchSyncState() {
  yield takeEvery(actions.monitor.monitorSyncState, syncState)
}

export default {
  watchSyncState,
}
