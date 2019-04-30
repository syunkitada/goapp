import { delay } from 'redux-saga'
import { put, call, takeEvery, all } from 'redux-saga/effects'
import actions from '../../actions'
import modules from '../../modules'

function* syncState(action) {
  console.log("home: syncState", action.payload)
}

function* watchSyncState() {
  yield takeEvery(actions.home.homeSyncState, syncState)
}

export default {
  watchSyncState
}
