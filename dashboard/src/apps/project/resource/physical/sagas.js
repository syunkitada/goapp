import { delay } from 'redux-saga'
import { put, call, takeEvery, all } from 'redux-saga/effects'
import actions from '../../../../actions'
import modules from '../../../../modules'

function* getIndex(action) {
  console.log("sagas.resource.physical.getIndex", action.payload)
  const {payload, error} = yield call(modules.resourcePhysical.getIndex, action.payload)

  console.log(payload)
  console.log(error)

  if (error) {
    yield put(actions.resourcePhysical.resourcePhysicalGetIndexFailure(""))
  } else if (payload.error && payload.error != "") {
    yield put(actions.resourcePhysical.resourcePhysicalGetIndexFailure(""))
  } else {
    console.log("sagas.resource.physical.getIndex Success")
    const index = {
      Datacenters: payload.Datacenters,
    }
    yield put(actions.resourcePhysical.resourcePhysicalGetIndexSuccess(index))
    console.log("yield puted actions.resourcePhysical.resourcePhysicalGetIndexSuccess")
  }
}

function* watchGetIndex() {
  yield takeEvery(actions.resourcePhysical.resourcePhysicalGetIndex, getIndex)
}

export default {
  watchGetIndex,
}
