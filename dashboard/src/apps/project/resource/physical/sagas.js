import { delay } from 'redux-saga'
import { put, call, takeEvery, all } from 'redux-saga/effects'
import actions from '../../../../actions'
import modules from '../../../../modules'


function* post(action) {
  const {payload, error} = yield call(modules.resourcePhysical.post, action.payload)
  console.log("HOGEllwlwlwlwllwlwsaga")
  console.log(action)

  if (error) {
    yield put(actions.resourcePhysical.resourcePhysicalPostFailure(action, error, null))
  } else if (payload.error && payload.error != "") {
    yield put(actions.resourcePhysical.resourcePhysicalPostFailure(action, null, payload.error))
  } else {
    yield put(actions.resourcePhysical.resourcePhysicalPostSuccess(action, payload))
  }
}

function* watchGetIndex() {
  yield takeEvery(actions.resourcePhysical.resourcePhysicalGetIndex, post)
}

function* watchGetDatacenterIndex() {
  yield takeEvery(actions.resourcePhysical.resourcePhysicalGetDatacenterIndex, post)
}

export default {
  watchGetIndex,
  watchGetDatacenterIndex,
}
