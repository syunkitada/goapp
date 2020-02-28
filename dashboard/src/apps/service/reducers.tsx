import { reducerWithInitialState } from "typescript-fsa-reducers";

import actions from "../../actions";
import logger from "../../lib/logger";

const defaultState = {
  isFetching: false,
  isSubmitting: false,

  getIndexTctx: null,
  getQueriesTctx: null,
  submitQueriesTctx: null,

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
    logger.info("reducers", "serviceGetIndex", payload);
    const params = payload.route.match.params;
    const service = params.service;
    const project = params.project;
    const newState = Object.assign({}, state, {
      getIndexTctx: {
        fetching: true
      },
      projectName: project,
      serviceName: service,
      syncDelay: 10000,
      syncQueryMap: {}
    });

    if (project) {
      if (!newState.projectServiceMap[project]) {
        newState.projectServiceMap[project] = {};
      }
      if (!newState.projectServiceMap[project][service]) {
        newState.projectServiceMap[project][service] = {
          isFetching: true
        };
      } else {
        newState.projectServiceMap[project][service].isFetching = true;
      }
    } else {
      if (!newState.serviceMap[service]) {
        newState.serviceMap[service] = {
          isFetching: true
        };
      } else {
        newState.serviceMap[service].isFetching = true;
      }
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
  .case(actions.service.serviceGetQueries, state => {
    logger.info("reducers", "serviceGetQueries");
    return Object.assign({}, state, {
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
    logger.info("reducers", "servicePostSuccess", payload.action.type, payload);
    const { index, route } = payload.action.payload;
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
            data = Object.assign(
              data,
              payload.result.ResultMap[query.Name].Data.Data
            );
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
          data = Object.assign(data, payload.result.ResultMap[query.Name].Data);
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
    }

    console.log("TODO DEBUG index", index, payload.websocket);
    const { websocket } = payload;
    const service = route.match.params.service;
    const project = route.match.params.project;
    // set data, and websocket
    if (project) {
      newState.projectServiceMap[project][service].isFetching = false;
      if (isGetIndex) {
        newState.projectServiceMap[project][service].Index = dashboardIndex;
      }

      // Set Data
      if (newState.projectServiceMap[project][service].Data) {
        for (const key of Object.keys(data)) {
          newState.projectServiceMap[project][service].Data[key] = data[key];
        }
      } else {
        newState.projectServiceMap[project][service].Data = data;
      }

      // Set WebSocket
      if (index && index.EnableWebSocket) {
        if (!newState.projectServiceMap[project][service].WebSocketMap) {
          newState.projectServiceMap[project][service].WebSocketMap = {};
          newState.projectServiceMap[project][service].WebSocketDataMap = {};
        }
        // TODO check exists websocket
        newState.projectServiceMap[project][service].WebSocketMap[
          index.WebSocketKey
        ] = websocket;
      }
    } else {
      newState.serviceMap[service].isFetching = false;
      if (isGetIndex) {
        newState.serviceMap[service].Index = dashboardIndex;
      }

      // Set Data
      if (newState.serviceMap[service].Data) {
        for (const key of Object.keys(data)) {
          newState.serviceMap[service].Data[key] = data[key];
        }
      } else {
        newState.serviceMap[service].Data = data;
      }

      // Set WebSocket
      if (index && index.EnableWebSocket) {
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

    const service = payload.action.payload.params.service;
    const project = payload.action.payload.params.project;
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
    logger.info("reducers", "servicePostSuccess", payload.action.type, payload);
    const { index, route } = payload.action.payload;
    const newState = Object.assign({}, state);
    const data = payload.result;

    const service = route.match.params.service;
    const project = route.match.params.project;
    // set websocket data
    if (project) {
      // Set WebSocket
      if (index && index.EnableWebSocket) {
        if (
          index.WebSocketKey in
          newState.projectServiceMap[project][service].WebSocketDataMap
        ) {
          let value =
            newState.projectServiceMap[project][service].WebSocketDataMap[
              index.WebSocketKey
            ];
          value += data.Value;
          newState.projectServiceMap[project][service].WebSocketDataMap[
            index.WebSocketKey
          ] = value;
        } else {
          const value = data.Value;
          newState.projectServiceMap[project][service].WebSocketDataMap[
            index.WebSocketKey
          ] = value;
        }
      }
    } else {
      // Set WebSocketData
      if (index && index.EnableWebSocket) {
        if (
          index.WebSocketKey in newState.serviceMap[service].WebSocketDataMap
        ) {
          let value =
            newState.serviceMap[service].WebSocketDataMap[index.WebSocketKey];
          value += data.Value;
          newState.serviceMap[service].WebSocketDataMap[
            index.WebSocketKey
          ] = value;
        } else {
          const value = data.Value;
          newState.serviceMap[service].WebSocketDataMap[
            index.WebSocketKey
          ] = value;
        }
      }
    }

    logger.info("reducers", "serviceWebSocketEmitSuccess: newState", newState);
    return newState;
  });
