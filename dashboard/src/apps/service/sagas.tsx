import {
  call,
  cancel,
  cancelled,
  delay,
  fork,
  put,
  take,
  takeEvery,
} from 'redux-saga/effects';

import actions from '../../actions';
import logger from '../../lib/logger';
import modules from '../../modules';
import store from '../../store';

function* post(action) {
  console.log('DEBUG post', action);
  const dataQueries: any[] = [];
  const {
    queries,
    params,
    isSync,
    queryKind,
    fieldMap,
    targets,
  } = action.payload;

  let payload: any = null;

  switch (action.type) {
    case 'SERVICE_GET_INDEX':
      payload = {
        projectName: params.project,
        queries: [{Kind: 'GetIndex', StrParams: action.payload}],
        serviceName: params.service,
        stateKey: 'index',
      };
      break;
    case 'SERVICE_GET_QUERIES':
      for (let i = 0, len = queries.length; i < len; i++) {
        dataQueries.push({Kind: queries[i], StrParams: params});
      }
      payload = {
        actionName: 'UserQuery',
        isSync,
        projectName: params.project,
        queries: dataQueries,
        serviceName: params.service,
        stateKey: 'index',
      };

      break;

    case 'SERVICE_SUBMIT_QUERIES':
      const strParams = Object.assign({}, params);
      const numParams = {};

      const spec = Object.assign({}, params);
      for (const key of Object.keys(fieldMap)) {
        const field = fieldMap[key];
        spec[key] = field.value;
      }
      const specsStr = JSON.stringify([spec]);
      strParams.Specs = specsStr;

      if (targets) {
        for (let i = 0, len = targets.length; i < len; i++) {
          const target = targets[i];
          strParams.Target = target;
          dataQueries.push({
            Kind: queryKind,
            NumParams: numParams,
            StrParams: strParams,
          });
        }
      } else {
        dataQueries.push({
          Kind: queryKind,
          NumParams: numParams,
          StrParams: strParams,
        });
      }

      payload = {
        actionName: 'UserQuery',
        projectName: params.project,
        queries: dataQueries,
        serviceName: params.service,
        stateKey: 'index',
      };
      break;

    default:
      return {};
  }

  // SERVICE_GET_QUERIES: (queries, isSync, params) => {
  //   let dataQueries = [];
  //   for (let i = 0, len = queries.length; i < len; i ++) {
  //     dataQueries.push({Kind: queries[i], StrParams: params})
  //   }
  //   return {
  //     stateKey: 'index',
  //     serviceName: params.service,
  //     actionName: 'UserQuery',
  //     projectName: params.project,
  //     queries: dataQueries,
  //     isSync: isSync,
  //   }
  // },

  // SERVICE_SUBMIT_QUERIES: (queryKind, action, fieldMap, targets, params) => {
  //   let dataQueries = [];
  //   let strParams = Object.assign({}, params)
  //   let numParams = {}

  //   let spec = Object.assign({}, params)
  //   for (let key in fieldMap) {
  //     let field = fieldMap[key]
  //     spec[key] = field.value
  //   }
  //   let specsStr = JSON.stringify([spec])
  //   strParams['Specs'] = specsStr

  //   if (targets) {
  //     for (let i = 0, len = targets.length; i < len; i ++) {
  //       let target = targets[i]
  //       strParams.Target = target
  //       dataQueries.push({Kind: queryKind, StrParams: strParams, NumParams: numParams})
  //     }
  //   } else {
  //     dataQueries.push({Kind: queryKind, StrParams: strParams, NumParams: numParams})
  //   }

  //   return {
  //     stateKey: 'index',
  //     serviceName: params.service,
  //     actionName: 'UserQuery',
  //     projectName: params.project,
  //     queries: dataQueries,
  //   }
  // },

  console.log('DEBUG post2', payload);

  const {result, error} = yield call(modules.service.post, payload);

  if (error) {
    yield put(actions.service.servicePostFailure(action, error));
  } else {
    yield put(actions.service.servicePostSuccess(action, result));
  }
}

function* sync(action) {
  try {
    while (true) {
      const serviceState = Object.assign({}, store.getState().service);
      if (serviceState.syncAction) {
        logger.info(['saga', 'sync', 'syncAction']);
        yield call(post, serviceState.syncAction);
      } else {
        logger.info(['saga', 'sync', 'syncAction is null']);
      }
      yield delay(serviceState.syncDelay);
    }
  } finally {
    if (yield cancelled()) {
      logger.info(['saga', 'sync', 'finally']);
      // yield put(actions.requestFailure('Sync cancelled!'))
    }
  }
}

function* bgSync(action) {
  // starts the task in the background
  const bgSyncTask = yield fork(sync, action);

  // wait for the user stop action
  yield take(actions.service.serviceStopBackgroundSync);
  // user clicked stop. cancel the background task
  // this will cause the forked bgSync task to jump into its finally block
  yield cancel(bgSyncTask);
}

function* watchGetIndex() {
  yield takeEvery(actions.service.serviceGetIndex, post);
}

function* watchStartBackgroundSync() {
  yield takeEvery(actions.service.serviceStartBackgroundSync, bgSync);
}

function* watchGetQueries() {
  yield takeEvery(actions.service.serviceGetQueries, post);
}

function* watchSubmitQueries() {
  yield takeEvery(actions.service.serviceSubmitQueries, post);
}

export default {
  watchGetIndex,
  watchGetQueries,
  watchStartBackgroundSync,
  watchSubmitQueries,
};
