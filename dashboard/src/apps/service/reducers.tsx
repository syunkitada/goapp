import {reducerWithInitialState} from 'typescript-fsa-reducers';

import actions from '../../actions';
import logger from '../../lib/logger';

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
  serviceMap: {},
};

export default reducerWithInitialState(defaultState)
  .case(actions.service.serviceGetIndex, (state, payload) => {
    logger.info('reducers', 'serviceGetIndex', payload.params);
    const service = payload.params.service;
    const project = payload.params.project;
    const newState = Object.assign({}, state, {
      getIndexTctx: {
        fetching: true,
      },
      projectName: project,
      serviceName: service,
      syncDelay: 10000,
      syncQueryMap: {},
    });

    if (project) {
      if (!newState.projectServiceMap[project]) {
        newState.projectServiceMap[project] = {};
      }
      if (!newState.projectServiceMap[project][service]) {
        newState.projectServiceMap[project][service] = {
          isFetching: true,
        };
      } else {
        newState.projectServiceMap[project][service].isFetching = true;
      }
    } else {
      if (!newState.serviceMap[service]) {
        newState.serviceMap[service] = {
          isFetching: true,
        };
      } else {
        newState.serviceMap[service].isFetching = true;
      }
    }

    logger.info('reducers', 'serviceGetIndex: newState', newState);

    return newState;
  })
  .case(actions.service.serviceStartBackgroundSync, state => {
    logger.info('reducers', 'serviceStartBackgroundSync');
    return Object.assign({}, state, {});
  })
  .case(actions.service.serviceStopBackgroundSync, state => {
    logger.info('reducers', 'serviceStopBackgroundSync');
    return Object.assign({}, state, {});
  })
  .case(actions.service.serviceGetQueries, state => {
    logger.info('reducers', 'serviceGetQueries');
    return Object.assign({}, state, {
      getQueriesTctx: {
        fetching: true,
      },
      openGetQueriesTctx: false,
    });
  })
  .case(actions.service.serviceSubmitQueries, state => {
    logger.info('reducers', 'serviceSubmitQueries');
    return Object.assign({}, state, {
      isFetching: true,
      isSubmitting: true,
      openSubmitQueriesTctx: false,
      submitQueriesTctx: {
        fetching: true,
      },
    });
  })
  .case(actions.service.serviceCloseErr, state => {
    logger.info('reducers', 'serviceCloseErr');
    return Object.assign({}, state, {
      error: null,
    });
  })
  .case(actions.service.serviceCloseGetQueriesTctx, state => {
    logger.info('reducers', 'serviceCloseGetQueriesTctx');
    return Object.assign({}, state, {
      openGetQueriesTctx: false,
    });
  })
  .case(actions.service.serviceCloseSubmitQueriesTctx, state => {
    logger.info('reducers', 'serviceCloseSubmitQueriesTctx');
    return Object.assign({}, state, {
      openSubmitQueriesTctx: false,
    });
  })
  .case(actions.service.servicePostSuccess, (state, payload) => {
    logger.info('reducers', 'servicePostSuccess', payload.action.type, payload);
    const newState = Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
    });

    if (payload.action.type === 'SERVICE_SUBMIT_QUERIES') {
      Object.assign(newState, {
        isSubmitting: false,
      });
    }

    const actionType = payload.action.type;
    const tctx = payload.result.Tctx;

    let isGetIndex = false;
    switch (actionType) {
      case 'SERVICE_GET_INDEX':
        newState.getIndexTctx = tctx;
        isGetIndex = true;
        break;
      case 'SERVICE_GET_QUERIES':
        newState.getQueriesTctx = tctx;
        newState.openGetQueriesTctx = true;
        break;
      case 'SERVICE_SUBMIT_QUERIES':
        newState.submitQueriesTctx = tctx;
        newState.openSubmitQueriesTctx = true;
        break;
      default:
        console.log('DEBUG unknownaction', actionType);
        break;
    }
    if (!tctx) {
      logger.error('reducers', 'servicePostError: not found tctx', newState);
      return newState;
    }

    if (tctx.StatusCode >= 300) {
      logger.error('reducers', 'servicePostError: newState', newState);
      // TODO handling tctx.Err, tctx.StatusCode
      return newState;
    }

    if (payload.action.payload.isSync) {
      console.log('DEBUG reducer query', payload.action.payload.syncQueryMap);
      newState.syncQueryMap = Object.assign(
        {},
        state.syncQueryMap,
        payload.payload.syncQueryMap,
      );
    }

    const index = payload.result.Index;
    if (index) {
      if (index.SyncDelay && index.SyncDelay > 1000) {
        newState.syncDelay = index.SyncDelay;
      }
    }

    const service = payload.action.payload.params.service;
    const project = payload.action.payload.params.project;
    if (project) {
      newState.projectServiceMap[project][service].isFetching = false;
      if (isGetIndex) {
        newState.projectServiceMap[project][service].Index =
          payload.result.Index;
      }
      if (newState.projectServiceMap[project][service].Data) {
        for (const key of Object.keys(payload.result.Data)) {
          newState.projectServiceMap[project][service].Data[key] =
            payload.result.Data[key];
        }
      } else {
        newState.projectServiceMap[project][service].Data = payload.result.Data;
      }
    } else {
      newState.serviceMap[service].isFetching = false;
      if (isGetIndex) {
        newState.serviceMap[service].Index = payload.result.Index;
      }
      if (newState.serviceMap[service].Data) {
        for (const key of Object.keys(payload.result.Data)) {
          newState.serviceMap[service].Data[key] = payload.result.Data[key];
        }
      } else {
        newState.serviceMap[service].Data = payload.result.Data;
      }
    }

    logger.info('reducers', 'servicePostSuccess: newState', newState);
    return newState;
  })
  .case(actions.service.servicePostFailure, (state, payload) => {
    const newState = Object.assign({}, state, {
      error: payload.error,
      isFetching: false,
    });

    if (payload.action.type === 'SERVICE_SUBMIT_QUERIES') {
      Object.assign(newState, {
        isSubmitting: false,
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
      'reducers',
      'servicePostFailure: newState',
      payload.action.type,
      newState,
    );
    return newState;
  });
