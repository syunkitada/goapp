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
    console.log("sagas.syncState Success")
    const monitor = {
      HostMap: payload.HostMap,
    }
    console.log(monitor)
    yield put(actions.monitor.monitorSyncStateSuccess(monitor))
    console.log("yield puted actions.monitor.monitorSyncStateSuccess")
  }
}

function* watchSyncState() {
  yield takeEvery(actions.monitor.monitorSyncState, syncState)
}

export default {
  watchSyncState,
}
