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
  const dataQueries: any[] = [];
  const {
    params,
    queries,
    isSync,
    queryKind,
    dataKind,
    items,
    fieldMap,
  } = action.payload;

  let payload: any = null;

  switch (action.type) {
    case 'SERVICE_GET_INDEX':
      payload = {
        projectName: params.project,
        queries: [
          {
            Data: JSON.stringify({Name: params.service}),
            Name: 'GetServiceDashboardIndex',
          },
        ],
        serviceName: params.service,
        stateKey: 'index',
      };
      break;

    case 'SERVICE_GET_QUERIES':
      const syncQueryMap: any[] = [];
      for (let i = 0, len = queries.length; i < len; i++) {
        dataQueries.push({Name: queries[i], Data: JSON.stringify(params)});
        if (isSync) {
          syncQueryMap[queries[i]] = {
            Data: JSON.stringify(params),
            Name: queries[i],
          };
        }
      }
      payload = {
        isSync,
        projectName: params.project,
        queries: dataQueries,
        serviceName: params.service,
        stateKey: 'index',
        syncQueryMap,
      };

      break;

    case 'SERVICE_SUBMIT_QUERIES':
      const specs: any[] = [];
      const spec = Object.assign({}, params);
      for (const key of Object.keys(fieldMap)) {
        const field = fieldMap[key];
        spec[key] = field.value;
      }

      for (let i = 0, len = items.length; i < len; i++) {
        specs.push({
          Kind: dataKind,
          Spec: Object.assign({}, spec, items[i]),
        });
      }

      dataQueries.push({
        Data: JSON.stringify({Spec: JSON.stringify(specs)}),
        Name: queryKind,
      });

      const serviceState = Object.assign({}, store.getState().service);
      if (serviceState.syncQueryMap) {
        for (const key of Object.keys(serviceState.syncQueryMap)) {
          dataQueries.push(serviceState.syncQueryMap[key]);
        }
      }

      payload = {
        projectName: params.project,
        queries: dataQueries,
        serviceName: params.service,
        stateKey: 'index',
      };

      break;

    default:
      return {};
  }

  const {result, error} = yield call(modules.service.post, payload);

  if (error) {
    yield put(actions.service.servicePostFailure({action, payload, error}));
  } else {
    yield put(actions.service.servicePostSuccess({action, payload, result}));
  }
}

function* sync(action) {
  try {
    while (true) {
      const serviceState = Object.assign({}, store.getState().service);
      if (serviceState.syncQueryMap) {
        const queries: any[] = [];
        for (const key of Object.keys(serviceState.syncQueryMap)) {
          queries.push(serviceState.syncQueryMap[key]);
        }
        const postAction = {
          payload: {
            params: {
              project: serviceState.projectName,
              service: serviceState.serviceName,
            },
          },
          type: 'SERVICE_GET_QUERIES',
        };
        const payload = {
          actionName: 'SERVICE_GET_QUERIES',
          projectName: serviceState.projectName,
          queries,
          serviceName: serviceState.serviceName,
        };
        logger.info('saga', 'sync', 'syncAction', action, postAction, payload);

        const {result, error} = yield call(modules.service.post, payload);
        if (error) {
          yield put(
            actions.service.servicePostFailure({
              action: postAction,
              error,
              payload,
            }),
          );
        } else {
          yield put(
            actions.service.servicePostSuccess({
              action: postAction,
              payload,
              result,
            }),
          );
        }
      } else {
        logger.info('saga', 'sync', 'syncAction is null');
      }
      yield delay(serviceState.syncDelay);
    }
  } finally {
    if (yield cancelled()) {
      logger.info('saga', 'sync', 'finally');
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
