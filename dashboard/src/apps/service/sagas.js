import { put, call, takeEvery, cancel, cancelled, fork, take, delay } from 'redux-saga/effects'
import actions from '../../actions'
import modules from '../../modules'

function* post(action) {
  console.log("sagas.post")
  const {payload, error} = yield call(modules.service.post, action.payload)

  if (error) {
    yield put(actions.service.servicePostFailure(action, error, null))
  } else if (payload.error && payload.error !== "") {
    yield put(actions.service.servicePostFailure(action, null, payload.error))
  } else {
    yield put(actions.service.servicePostSuccess(action, payload))
  }
}

function* sync(action) {
  try {
    while (true) {
      console.log("DEBUG sync", action)
      yield call(post, action)
      // const result = yield call(someApi)
      // yield put(actions.requestSuccess(result))
      yield delay(10000)
    }
  } finally {
    if (yield cancelled()) {
      console.log("DEBUG sync cancelled")
      // yield put(actions.requestFailure('Sync cancelled!'))
    }
  }
}

function* bgSync(action) {
   // starts the task in the background
  const bgSyncTask = yield fork(sync, action)

  // wait for the user stop action
  yield take(actions.service.serviceStopBackgroundSync)
  // user clicked stop. cancel the background task
  // this will cause the forked bgSync task to jump into its finally block
  yield cancel(bgSyncTask)
}

function* watchGetIndex() {
  yield takeEvery(actions.service.serviceGetIndex, post)
}

function* watchStartBackgroundSync() {
  yield takeEvery(actions.service.serviceStartBackgroundSync, bgSync)
}

function* watchGetQueries() {
  yield takeEvery(actions.service.serviceGetQueries, post)
}

function* watchSubmitQueries() {
  yield takeEvery(actions.service.serviceSubmitQueries, post)
}

export default {
  watchGetIndex,
  watchGetQueries,
  watchSubmitQueries,
  watchStartBackgroundSync,
}
