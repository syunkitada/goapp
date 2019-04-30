import { delay } from 'redux-saga'
import { put, call, takeEvery, all } from 'redux-saga/effects'
import actions from '../../../actions'
import modules from '../../../modules'

function* syncState(action) {
  console.log("resource: syncState", action.payload)
  const {payload, error} = yield call(modules.resource.syncState, action.payload)

  console.log(payload)
  console.log(error)

  if (error) {
    yield put(actions.resource.resourceSyncStateFailure(""))
  } else if (payload.error && payload.error != "") {
    yield put(actions.resource.resourceSyncStateFailure(""))
  } else {
    const resource = {
      payload: payload,
    }
    yield put(actions.resource.resourceSyncStateSuccess(resource))
  }
}

function* watchSyncState() {
  yield takeEvery(actions.resource.resourceSyncState, syncState)
}

export default {
  watchSyncState,
}
