import { eventChannel } from "redux-saga";
import {
  // apply,
  call,
  cancel,
  cancelled,
  delay,
  fork,
  put,
  take,
  takeEvery
} from "redux-saga/effects";

import actions from "../../actions";
import logger from "../../lib/logger";
import modules from "../../modules";
import store from "../../store";

function createWebSocketConnection() {
  const url: any = process.env.REACT_APP_AUTHPROXY_URL;
  const wsUrl: any = url.replace("http", "ws");
  const socket = new WebSocket(wsUrl + "/ws");
  return new Promise(resolve => {
    socket.onopen = event => {
      logger.info("websocket.onopen", event);
      resolve(socket);
    };
  });
}

function* post(action) {
  const dataQueries: any[] = [];
  const {
    index,
    state,
    route,
    searchQueries,
    isSync,
    queryKind,
    dataKind,
    items,
    fieldMap
  } = action.payload;

  console.log("DEBUG TODO post", index, state, route);
  if (!route.match) {
    logger.error("Invalid route", index, state, route);
    return;
  }
  const params = route.match.params;

  let payload: any = null;

  switch (action.type) {
    case "SERVICE_GET_INDEX":
      console.log("DEBUG TODO getindex");
      payload = {
        projectName: params.project,
        queries: [
          {
            Data: JSON.stringify({ Name: params.service }),
            Name: "GetServiceDashboardIndex"
          }
        ],
        serviceName: params.service,
        stateKey: "index"
      };
      break;

    case "SERVICE_GET_QUERIES":
      const syncQueryMap: any[] = [];
      const queryData = Object.assign({}, params, searchQueries);
      console.log("DEBUG TODO getqueries", index, state, route);

      for (let i = 0, len = index.DataQueries.length; i < len; i++) {
        dataQueries.push({
          Data: JSON.stringify(queryData),
          Name: index.DataQueries[i]
        });
        if (isSync) {
          syncQueryMap[index.DataQueries[i]] = {
            Data: JSON.stringify(params),
            Name: index.DataQueries[i]
          };
        }
      }
      payload = {
        isSync,
        projectName: params.project,
        queries: dataQueries,
        serviceName: params.service,
        stateKey: "index",
        syncQueryMap
      };

      console.log("DEBUG TODO getqueries", index, state, route, payload);
      break;

    case "SERVICE_SUBMIT_QUERIES":
      const specs: any[] = [];
      const spec = Object.assign({}, params);
      for (const key of Object.keys(fieldMap)) {
        const field = fieldMap[key];
        spec[key] = field.value;
      }

      for (let i = 0, len = items.length; i < len; i++) {
        specs.push({
          Kind: dataKind,
          Spec: Object.assign({}, spec, items[i])
        });
      }

      dataQueries.push({
        Data: JSON.stringify({ Spec: JSON.stringify(specs) }),
        Name: queryKind
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
        stateKey: "index"
      };

      break;

    default:
      return {};
  }

  console.log("DEBUG TODO post", payload);
  const { result, error } = yield call(modules.service.post, payload);
  console.log("DEBUG TODO post", payload, result, error);

  if (error) {
    yield put(actions.service.servicePostFailure({ action, payload, error }));
  } else {
    yield put(actions.service.servicePostSuccess({ action, payload, result }));
  }
  return {};
}

function* sync(action) {
  return;
  try {
    while (true) {
      const serviceState = Object.assign({}, store.getState().service);
      if (serviceState.syncQueryMap) {
        const queries: any[] = [];
        for (const key of Object.keys(serviceState.syncQueryMap)) {
          queries.push(serviceState.syncQueryMap[key]);
        }
        const route = {
          match: {
            params: {
              project: serviceState.projectName,
              service: serviceState.serviceName
            }
          }
        };
        const postAction = {
          payload: {
            route
          },
          type: "SERVICE_GET_QUERIES"
        };
        const payload = {
          actionName: "SERVICE_GET_QUERIES",
          projectName: serviceState.projectName,
          queries,
          route,
          serviceName: serviceState.serviceName
        };
        logger.info("saga", "sync", "syncAction", action, postAction, payload);

        // const { result, error } = yield call(modules.service.post, payload);
        // if (error) {
        //   yield put(
        //     actions.service.servicePostFailure({
        //       action: postAction,
        //       error,
        //       payload
        //     })
        //   );
        // } else {
        //   yield put(
        //     actions.service.servicePostSuccess({
        //       action: postAction,
        //       payload,
        //       result
        //     })
        //   );
        // }
      } else {
        logger.info("saga", "sync", "syncAction is null");
      }
      yield delay(serviceState.syncDelay);
    }
  } finally {
    if (yield cancelled()) {
      logger.info("saga", "sync", "finally");
      // yield put(actions.requestFailure('Sync cancelled!'))
    }
  }
}

// this function creates an event channel from a given socket
// Setup subscription to incoming `ping` events
function createSocketChannel(socket) {
  // `eventChannel` takes a subscriber function
  // the subscriber function takes an `emit` argument to put messages onto the channel
  return eventChannel(emit => {
    const pingHandler = event => {
      // puts event payload into the channel
      // this allows a Saga to take this payload from the returned channel
      emit(event.payload);
    };

    const errorHandler = errorEvent => {
      // create an Error object and put it into the channel
      emit(new Error(errorEvent.reason));
    };

    // setup the subscription
    socket.onmessage = pingHandler;
    socket.onerror = errorHandler;

    // the subscriber must return an unsubscribe function
    // this will be invoked when the saga calls `channel.close` method
    const unsubscribe = () => {
      socket.off("ping", pingHandler);
    };

    return unsubscribe;
  });
}

// reply with a `pong` message by invoking `socket.emit('pong')`
function* pong(socket) {
  yield delay(5000);
  console.log("pong");
  // yield apply(socket, socket.emit, ["pong"]); // call `emit` as a method with `socket` as context
}

function* watchWebSocket() {
  const socket = yield call(createWebSocketConnection);
  const socketChannel = yield call(createSocketChannel, socket);

  while (true) {
    try {
      // An error from socketChannel will cause the saga jump to the catch block
      const payload = yield take(socketChannel);
      console.log("DEBUG payload", payload);
      // yield put({ type: INCOMING_PONG_PAYLOAD, payload });
      yield fork(pong, socket);
    } catch (err) {
      console.error("socket error:", err);
      // socketChannel is still open in catch block
      // if we want end the socketChannel, we need close it explicitly
      // socketChannel.close()
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
  watchWebSocket
};
