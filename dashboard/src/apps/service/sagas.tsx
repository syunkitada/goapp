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

function* post(action) {
    const dataQueries: any[] = [];
    const serviceState = Object.assign({}, store.getState().service);
    const { projectName, serviceName, index, location } = serviceState;
    const { isSync, queryKind, dataKind, items, fieldMap } = action.payload;

    console.log("DEBUG TODO post init", index, action.type, location);
    if (!location) {
        logger.warning("Invalid location", location);
        return;
    }

    let payload: any = null;

    switch (action.type) {
        case "SERVICE_GET_INDEX":
            if (serviceState.projectName) {
                payload = {
                    projectName: serviceState.projectName,
                    queries: [
                        {
                            Data: JSON.stringify({
                                Name: serviceState.serviceName
                            }),
                            Name: "GetProjectServiceDashboardIndex"
                        }
                    ],
                    serviceName: serviceState.serviceName,
                    stateKey: "index"
                };
            } else {
                payload = {
                    projectName: serviceState.projectName,
                    queries: [
                        {
                            Data: JSON.stringify({
                                Name: serviceState.serviceName
                            }),
                            Name: "GetServiceDashboardIndex"
                        }
                    ],
                    serviceName: serviceState.serviceName,
                    stateKey: "index"
                };
            }
            if (serviceState.initLocation && location.DataQueries) {
                const queries: any[] = [];
                const tmpData = Object.assign(
                    {},
                    location.Params,
                    location.SearchQueries
                );
                const data = JSON.stringify(tmpData);
                for (
                    let i = 0, len = location.DataQueries.length;
                    i < len;
                    i++
                ) {
                    queries.push({
                        Data: data,
                        Name: location.DataQueries[i]
                    });
                }
                payload.queries = payload.queries.concat(queries);
            }
            break;

        case "SERVICE_GET_QUERIES":
            const syncQueryMap: any[] = [];
            const queryData = Object.assign(
                {},
                location.Params,
                location.SearchQueries
            );
            const data = JSON.stringify(queryData);
            if (location.DataQueries) {
                for (
                    let i = 0, len = location.DataQueries.length;
                    i < len;
                    i++
                ) {
                    dataQueries.push({
                        Data: data,
                        Name: location.DataQueries[i]
                    });
                }
                payload = {
                    isSync,
                    projectName: projectName,
                    serviceName: serviceName,
                    queries: dataQueries,
                    stateKey: "index",
                    syncQueryMap
                };
            }

            break;

        case "SERVICE_SUBMIT_QUERIES":
            const specs: any[] = [];
            const spec = Object.assign({}, location.params);
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

            if (serviceState.syncQueryMap) {
                for (const key of Object.keys(serviceState.syncQueryMap)) {
                    dataQueries.push(serviceState.syncQueryMap[key]);
                }
            }

            payload = {
                projectName: projectName,
                serviceName: serviceName,
                queries: dataQueries,
                stateKey: "index"
            };

            break;
    }

    let isError = false;
    if (payload) {
        console.log("DEBUG TODO post", payload);
        const { result, error } = yield call(modules.service.post, payload);
        console.log("DEBUG TODO post", payload, result, error);

        if (error) {
            isError = true;
            yield put(
                actions.service.servicePostFailure({ action, payload, error })
            );
        } else {
            yield put(
                actions.service.servicePostSuccess({
                    action,
                    payload,
                    result,
                    websocket: null
                })
            );
        }
    }

    if (!isError && location.WebSocketQuery) {
        console.log("DEBUG TODO WebSocketQuery");
        const syncQueryMap: any[] = [];
        const queryData = Object.assign(
            {},
            location.Params,
            location.SearchQueries
        );
        const data = JSON.stringify(queryData);
        const dataQueries = [
            {
                Data: data,
                Name: location.WebSocketQuery
            }
        ];
        const payload = {
            isSync,
            projectName: projectName,
            serviceName: serviceName,
            queries: dataQueries,
            stateKey: "index",
            syncQueryMap
        };

        yield put(actions.service.serviceStartWebSocket({ action, payload }));
        return;
    }
    return {};
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
                const route = {
                    params: {
                        project: serviceState.projectName,
                        service: serviceState.serviceName
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
                logger.info(
                    "saga",
                    "sync",
                    "syncAction",
                    action,
                    postAction,
                    payload
                );

                // TODO
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

function createWebSocketConnection() {
    const url: any = process.env.REACT_APP_AUTHPROXY_URL;
    const wsUrl: any = url.replace("http", "ws");
    const socket = new WebSocket(wsUrl + "/ws");
    return new Promise(resolve => {
        socket.onopen = event => {
            resolve(socket);
        };
    });
}

function createSocketChannel(socket) {
    // eventChannelは、WebSocketなどからの外部イベントを受けて、ActionとしてChannelに積むためのもの
    return eventChannel(emit => {
        // subscriptionを設定
        // socketから受け取ったメッセージをactionデータとしてchannelに入れる(emit)
        const messageHandler = event => {
            emit(actions.service.serviceWebSocketOnMessage({ event }));
        };
        const errorHandler = event => {
            emit(actions.service.serviceWebSocketOnError({ event }));
        };
        socket.onmessage = messageHandler;
        socket.onerror = errorHandler;

        // 最後にunsubscribeを返す
        // unsubscribeは、sagaが `channel.close`メソッドを実行したときに呼び出されるので、後処理を行う
        const unsubscribe = () => {
            socket.close();
        };
        return unsubscribe;
    });
}

function* startWebSocket(action) {
    logger.info("startWebSocket", action);
    const socket = yield call(createWebSocketConnection);
    const socketChannel = yield call(createSocketChannel, socket);
    const { projectName, serviceName, queries } = action.payload.payload;
    logger.info("createdSocketChannel", projectName, serviceName, queries);

    const body = JSON.stringify({
        Project: projectName,
        Queries: queries,
        Service: serviceName
    });

    // コネクション確立後の初回メッセージにより認証が行われる
    socket.send(body);
    let isInit = true;

    while (true) {
        try {
            const payload = yield take(socketChannel);
            switch (payload.type) {
                case "SERVICE_WEB_SOCKET_ON_MESSAGE":
                    console.log("TODO taked on message", payload);
                    const data = JSON.parse(payload.payload.event.data);
                    if (isInit) {
                        const tmpAction = Object.assign(
                            {},
                            action.payload.action,
                            {
                                type: "SERVICE_GET_QUERIES"
                            }
                        );
                        const result = data;
                        yield put(
                            actions.service.servicePostSuccess({
                                action: tmpAction,
                                payload: action.payload.payload,
                                result,
                                websocket: socket
                            })
                        );
                        isInit = false;
                        break;
                    }

                    yield put(
                        actions.service.serviceWebSocketEmitMessage({
                            action: action.payload.action,
                            payload: action.payload.payload,
                            result: data
                        })
                    );

                    break;
                case "SERVICE_WEB_SOCKET_ON_ERROR":
                    console.log("TODO taked on error", payload);
                    break;
                default:
                    logger.error(
                        "take(socketChannel): Invalid action.type",
                        payload.type
                    );
                    break;
            }
            // yield put({ type: INCOMING_PONG_PAYLOAD, payload });
            // yield fork(pong, socket);
        } catch (err) {
            logger.error("TODO handle socket err", err);
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

function* watchStartWebSocket() {
    yield takeEvery(actions.service.serviceStartWebSocket, startWebSocket);
}

export default {
    watchGetIndex,
    watchGetQueries,
    watchStartBackgroundSync,
    watchStartWebSocket,
    watchSubmitQueries
};
