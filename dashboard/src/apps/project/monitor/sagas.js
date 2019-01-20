import { delay } from 'redux-saga'
import { put, call, takeEvery, all } from 'redux-saga/effects'
import actions from '../../../actions'
import modules from '../../../modules'

function* syncState(action) {
  console.log("monitor.sagas.syncState", action.payload)
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
      IndexMap: payload.IndexMap,
    }
    console.log(monitor)
    yield put(actions.monitor.monitorSyncStateSuccess(monitor))
    console.log("yield puted actions.monitor.monitorSyncStateSuccess")
  }
}

function* syncIndexState(action) {
  console.log("monitor.sagas.syncIndexState", action.payload)
  const {payload, error} = yield call(modules.monitor.syncIndexState, action.payload)

  console.log(payload)
  console.log(error)

  if (error) {
    yield put(actions.monitor.monitorSyncIndexStateFailure(""))
  } else if (payload.error && payload.error != "") {
    yield put(actions.monitor.monitorSyncIndexStateFailure(""))
  } else {
    console.log("sagas.syncIndexState Success")
    const indexState = {
      HostMap: payload.IndexMap,
    }
    console.log(indexState)
    yield put(actions.monitor.monitorSyncIndexStateSuccess(indexState))
    console.log("yield puted actions.monitor.monitorSyncIndexStateSuccess")
  }
}

function* watchSyncState() {
  yield takeEvery(actions.monitor.monitorSyncState, syncState)
}

function* watchSyncIndexState() {
  yield takeEvery(actions.monitor.monitorSyncIndexState, syncIndexState)
}

export default {
  watchSyncState,
  watchSyncIndexState,
}
