import { reducerWithInitialState } from "typescript-fsa-reducers";

import actions from "../../actions";
import logger from "../../lib/logger";
import data_utils from "../../lib/data_utils";

type index = {
    DataQueries?: any;
    EnableWebSocket?: boolean;
    WebSocketKey: string;
    View: any;
};

type serviceState = {
    isFetching: boolean;
    isSubmitting: boolean;

    getIndexTctx: any;
    getQueriesTctx: any;
    submitQueriesTctx: any;

    dashboardIndex: any;
    rootIndex: any;
    index?: index;
    location: any;
    initLocation: boolean;

    openGetQueriesTctx: boolean;
    openSubmitQueriesTctx: boolean;

    error: any;
    payloadError: any;

    projectName: any;
    serviceName: any;

    syncDelay: number;
    syncQueryMap: any;

    projectServiceMap: any;
    serviceMap: any;
};

const defaultState: serviceState = {
    isFetching: false,
    isSubmitting: false,

    getIndexTctx: null,
    getQueriesTctx: null,
    submitQueriesTctx: null,

    dashboardIndex: null,
    rootIndex: null,
    location: null,
    initLocation: false,

    openGetQueriesTctx: false,
    openSubmitQueriesTctx: false,

    error: null,
    payloadError: null,

    projectName: null,
    serviceName: null,

    syncDelay: 10000,
    syncQueryMap: {},

    projectServiceMap: {},
    serviceMap: {}
};

export default reducerWithInitialState(defaultState)
    .case(actions.service.serviceGetIndex, (state, payload) => {
        let { serviceName, projectName } = payload;
        if (!serviceName) {
            const params = data_utils.getServiceParams();
            serviceName = params.serviceName;
            projectName = params.projectName;
        } else {
            data_utils.setServiceParams(payload);
        }
        const newState = Object.assign({}, state, {
            getIndexTctx: {
                fetching: true
            },
            projectName: projectName,
            serviceName: serviceName,
            syncDelay: 10000,
            syncQueryMap: {},
            location: {
                Path: ["Root"]
            }
        });

        if (projectName) {
            if (!newState.projectServiceMap[projectName]) {
                newState.projectServiceMap[projectName] = {};
            }
            if (!newState.projectServiceMap[projectName][serviceName]) {
                newState.projectServiceMap[projectName][serviceName] = {
                    isFetching: true
                };
            } else {
                newState.projectServiceMap[projectName][
                    serviceName
                ].isFetching = true;
            }
        } else {
            if (!newState.serviceMap[serviceName]) {
                newState.serviceMap[serviceName] = {
                    isFetching: true
                };
            } else {
                newState.serviceMap[serviceName].isFetching = true;
            }
        }

        let data = data_utils.getLocationData();
        if (data != null) {
            newState.location = data;
            newState.initLocation = true;
        }
        logger.info("reducers", "serviceGetIndex: newState", newState);
        return newState;
    })
    .case(actions.service.serviceStartWebSocket, state => {
        logger.info("reducers", "serviceStartWebSocket");
        return Object.assign({}, state, {});
    })
    .case(actions.service.serviceStartBackgroundSync, state => {
        logger.info("reducers", "serviceStartBackgroundSync");
        return Object.assign({}, state, {});
    })
    .case(actions.service.serviceStopBackgroundSync, state => {
        logger.info("reducers", "serviceStopBackgroundSync");
        return Object.assign({}, state, {});
    })
    .case(actions.service.serviceGetQueries, (state, payload) => {
        const rootIndex = state.rootIndex;
        const index = data_utils.getIndex(rootIndex, payload.location.Path);
        const location = payload.location;
        location.DataQueries = index.DataQueries;
        location.WebSocketQuery = index.WebSocketQuery;
        const params = Object.assign(
            {},
            state.location.Params,
            location.Params
        );
        location.Params = params;
        data_utils.setLocationData(location);
        console.log("DEBUG TODO serviceGetQueries", location);

        return Object.assign({}, state, {
            isFetching: true,
            location: location,
            index: index,
            getQueriesTctx: {
                fetching: true
            },
            openGetQueriesTctx: false
        });
    })
    .case(actions.service.serviceSubmitQueries, state => {
        logger.info("reducers", "serviceSubmitQueries");
        return Object.assign({}, state, {
            isFetching: true,
            isSubmitting: true,
            openSubmitQueriesTctx: false,
            submitQueriesTctx: {
                fetching: true
            }
        });
    })
    .case(actions.service.serviceCloseErr, state => {
        logger.info("reducers", "serviceCloseErr");
        return Object.assign({}, state, {
            error: null
        });
    })
    .case(actions.service.serviceCloseGetQueriesTctx, state => {
        logger.info("reducers", "serviceCloseGetQueriesTctx");
        return Object.assign({}, state, {
            openGetQueriesTctx: false
        });
    })
    .case(actions.service.serviceCloseSubmitQueriesTctx, state => {
        logger.info("reducers", "serviceCloseSubmitQueriesTctx");
        return Object.assign({}, state, {
            openSubmitQueriesTctx: false
        });
    })
    .case(actions.service.servicePostSuccess, (state, payload) => {
        logger.info(
            "reducers",
            "servicePostSuccess",
            payload.action.type,
            payload
        );
        const newState = Object.assign({}, state, {
            isFetching: false,
            redirectToReferrer: true
        });

        if (payload.action.type === "SERVICE_SUBMIT_QUERIES") {
            Object.assign(newState, {
                isSubmitting: false
            });
        }

        const actionType = payload.action.type;

        // Merge tctx
        const tctx: any = {
            Error: payload.result.Error,
            StatusCode: payload.result.Code,
            TraceId: payload.result.TraceId
        };
        if (payload.result.ResultMap) {
            for (const key of Object.keys(payload.result.ResultMap)) {
                const result: any = payload.result.ResultMap[key];
                if (tctx.StatusCode < result.Code) {
                    tctx.StatusCode = result.Code;
                    tctx.Error = result.Error;
                }
            }
        }

        let isGetIndex = false;
        switch (actionType) {
            case "SERVICE_GET_INDEX":
                newState.getIndexTctx = tctx;
                isGetIndex = true;
                break;
            case "SERVICE_GET_QUERIES":
                newState.getQueriesTctx = tctx;
                newState.openGetQueriesTctx = true;
                break;
            case "SERVICE_SUBMIT_QUERIES":
                newState.submitQueriesTctx = tctx;
                newState.openSubmitQueriesTctx = true;
                break;
            default:
                console.log("DEBUG unknownaction", actionType);
                break;
        }

        if (tctx.StatusCode >= 300) {
            logger.error("reducers", "servicePostError: newState", newState);
            // TODO handling tctx.Err, tctx.StatusCode
            return newState;
        }

        // updateSyncQueryMap
        if (payload.action.payload.isSync) {
            newState.syncQueryMap = Object.assign(
                {},
                state.syncQueryMap,
                payload.payload.syncQueryMap
            );
        }

        // Merge data
        let data = {};
        if (isGetIndex) {
            for (const query of payload.payload.queries) {
                if (payload.result.ResultMap[query.Name]) {
                    if (payload.result.ResultMap[query.Name].Data) {
                        if (payload.result.ResultMap[query.Name].Data.Data) {
                            data = Object.assign(
                                data,
                                payload.result.ResultMap[query.Name].Data.Data
                            );
                        } else {
                            data = Object.assign(
                                data,
                                payload.result.ResultMap[query.Name].Data
                            );
                        }
                    } else {
                        logger.warning(
                            "reducers",
                            "servicePostSuccess: QueryData is not found",
                            query.Name
                        );
                    }
                } else {
                    logger.warning(
                        "reducers",
                        "servicePostSuccess: QueryResult is not found",
                        query.Name
                    );
                }
            }
        } else {
            for (const query of payload.payload.queries) {
                if (query.Name in payload.result.ResultMap) {
                    data = Object.assign(
                        data,
                        payload.result.ResultMap[query.Name].Data
                    );
                } else {
                    logger.warning(
                        "reducers",
                        "servicePostSuccess: QueryResult is not found",
                        query.Name
                    );
                    return newState;
                }
            }
        }

        let dashboardIndex: any = null;
        if (isGetIndex) {
            dashboardIndex =
                payload.result.ResultMap.GetServiceDashboardIndex.Data.Index;
            if (dashboardIndex.SyncDelay && dashboardIndex.SyncDelay > 1000) {
                newState.syncDelay = dashboardIndex.SyncDelay;
            }
            newState.dashboardIndex = dashboardIndex;
            newState.rootIndex = dashboardIndex.View;
            if (!state.initLocation) {
                newState.location = dashboardIndex.DefaultRoute;
            }
        }

        const index = data_utils.getIndex(
            newState.rootIndex,
            newState.location.Path
        );

        const { websocket } = payload;
        const service = state.serviceName;
        const project = state.projectName;
        // set data, and websocket
        if (project) {
            newState.projectServiceMap[project][service].isFetching = false;
            // Set Data
            if (newState.projectServiceMap[project][service].Data) {
                for (const key of Object.keys(data)) {
                    newState.projectServiceMap[project][service].Data[key] =
                        data[key];
                }
            } else {
                newState.projectServiceMap[project][service].Data = data;
            }

            // Set WebSocket
            if (
                actionType === "SERVICE_GET_QUERIES" &&
                index &&
                index.EnableWebSocket
            ) {
                if (
                    !newState.projectServiceMap[project][service].WebSocketMap
                ) {
                    newState.projectServiceMap[project][
                        service
                    ].WebSocketMap = {};
                    newState.projectServiceMap[project][
                        service
                    ].WebSocketDataMap = {};
                }
                // TODO check exists websocket
                newState.projectServiceMap[project][service].WebSocketMap[
                    index.WebSocketKey
                ] = websocket;
            }
        } else {
            // Set Data
            if (newState.serviceMap[service].Data) {
                for (const key of Object.keys(data)) {
                    newState.serviceMap[service].Data[key] = data[key];
                }
            } else {
                newState.serviceMap[service].Data = data;
            }

            // Set WebSocket
            if (
                actionType === "SERVICE_GET_QUERIES" &&
                index &&
                index.EnableWebSocket
            ) {
                if (!newState.serviceMap[service].WebSocketMap) {
                    newState.serviceMap[service].WebSocketMap = {};
                    newState.serviceMap[service].WebSocketDataMap = {};
                }
                // TODO check exists websocket
                newState.serviceMap[service].WebSocketMap[
                    index.WebSocketKey
                ] = websocket;
            }
        }

        logger.info("reducers", "servicePostSuccess: newState", newState);
        return newState;
    })
    .case(actions.service.servicePostFailure, (state, payload) => {
        const newState = Object.assign({}, state, {
            error: payload.error,
            isFetching: false
        });

        if (payload.action.type === "SERVICE_SUBMIT_QUERIES") {
            Object.assign(newState, {
                isSubmitting: false
            });
        }

        const service = state.serviceName;
        const project = state.projectName;
        if (project) {
            newState.projectServiceMap[project][service].isFetching = false;
        } else {
            newState.serviceMap[service].isFetching = false;
        }

        logger.error(
            "reducers",
            "servicePostFailure: newState",
            payload.action.type,
            newState
        );
        return newState;
    })
    .case(actions.service.serviceWebSocketEmitMessage, (state, payload) => {
        logger.info(
            "reducers",
            "serviceWebSocketEmitMessage",
            payload.action.type,
            payload
        );
        const { index } = payload.action.payload;
        const newState = Object.assign({}, state);
        const data = payload.result;

        if (!index || !index.EnableWebSocket) {
            return newState;
        }
        console.log("TODO value:", data.Value, data.Bytes);
        if (!data.Bytes) {
            return newState;
        }

        const service = state.serviceName;
        const project = state.projectName;
        let value = "";
        // get value
        if (project) {
            if (
                index.WebSocketKey in
                newState.projectServiceMap[project][service].WebSocketDataMap
            ) {
                value =
                    newState.projectServiceMap[project][service]
                        .WebSocketDataMap[index.WebSocketKey];
            } else {
                value = "";
            }
        } else {
            if (
                index.WebSocketKey in
                newState.serviceMap[service].WebSocketDataMap
            ) {
                value =
                    newState.serviceMap[service].WebSocketDataMap[
                        index.WebSocketKey
                    ];
            } else {
                value = "";
            }
        }

        console.log("TODO value:", data.Value, data.Bytes);
        let specificPrefix = false;
        const decoder = new TextDecoder();
        for (let i = 0, len = data.Bytes.length; i < len; i++) {
            const bytes = data.Bytes[i];
            const lenbytes = bytes.length;
            console.log("TODO byte:", bytes);
            if (lenbytes === 1) {
                if (specificPrefix) {
                    console.log("TODO byte2:", bytes);
                    if (bytes[0] === 8) {
                        value = value.slice(0, value.length - 1);
                    }
                    // up 75
                    // down 97
                    // left 67
                } else {
                    specificPrefix = false;
                    // left 8
                    if (bytes[0] > 22) {
                        value += decoder.decode(Uint8Array.from(bytes));
                    }
                }
            } else if (lenbytes === 2) {
                if (bytes[0] === 8 && bytes[1] === 27) {
                    specificPrefix = true;
                } else if (specificPrefix && bytes[0] === 91) {
                    console.log("TODO byte2:", bytes[1]);
                    if (bytes[1] === 8) {
                        value = value.slice(0, value.length - 1);
                    }
                } else {
                    specificPrefix = false;
                    value += decoder.decode(Uint8Array.from(bytes));
                }
            } else if (lenbytes === 3) {
                specificPrefix = false;
                if (bytes[0] === 8 && bytes[1] === 27 && bytes[2] === 91) {
                    value = value.slice(0, value.length - 2);
                } else {
                    value += decoder.decode(Uint8Array.from(bytes));
                }
            } else if (lenbytes === 4) {
                if (bytes[0] === 8 && bytes[1] === 27 && bytes[2] === 91) {
                    console.log("TODO byte2:", bytes[3]);
                    if (bytes[3] === 75) {
                        // bs
                        value = value.slice(0, value.length - 1);
                    }
                } else {
                    value += decoder.decode(Uint8Array.from(bytes));
                }
            } else {
                specificPrefix = false;
                value += decoder.decode(Uint8Array.from(bytes));
            }
        }

        // set websocket data
        if (project) {
            newState.projectServiceMap[project][service].WebSocketDataMap[
                index.WebSocketKey
            ] = value;
        } else {
            newState.serviceMap[service].WebSocketDataMap[
                index.WebSocketKey
            ] = value;
        }

        logger.info(
            "reducers",
            "serviceWebSocketEmitSuccess: newState",
            newState
        );
        return newState;
    });
