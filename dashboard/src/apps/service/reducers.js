import { handleActions } from 'redux-actions';

import actions from '../../actions'
import logger from '../../lib/logger'

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
  index: {Index: null, Data: null},
  datacenterIndex: null,
  serviceName: null,
  projectName: null,
  syncAction: null,
  syncDelay: 10000,
  serviceMap: {},
  projectServiceMap: {},
};

export default handleActions({
  [actions.service.serviceGetIndex]: (state, action) => {
    logger.info("reducers", "serviceGetIndex")
    let service = action.payload.serviceName
    let project = action.payload.projectName
    let newState = Object.assign({}, state, {
      serviceName: service,
      projectName: project,
      getIndexTctx: {
        fetching: true
      },
      syncAction: null,
      syncDelay: 10000,
    })

    if (project) {
      if (!newState.projectServiceMap[project]) {
        newState.projectServiceMap[project] = {}
      }
      if (!newState.projectServiceMap[project][service]) {
        newState.projectServiceMap[project][service] = {
          isFetching: true,
        }
      } else {
        newState.projectServiceMap[project][service].isFetching = true
      }
    } else {
      if (!newState.serviceMap[service]) {
        newState.serviceMap[service] = {
          isFetching: true,
        }
      } else {
        newState.serviceMap[service].isFetching = true
      }
    }

    logger.info("reducers", "serviceGetIndex", newState)

    return newState;
  },

  [actions.service.serviceStartBackgroundSync]: (state) => {
    logger.info("reducers", "serviceStartBackgroundSync")
    return Object.assign({}, state, {})
  },

  [actions.service.serviceStopBackgroundSync]: (state) => {
    logger.info("reducers", "serviceStopBackgroundSync")
    return Object.assign({}, state, {})
  },

  [actions.service.serviceGetQueries]: (state) => {
    logger.info("reducers", "serviceGetQueries")
    return Object.assign({}, state, {
      getQueriesTctx: {
        fetching: true
      },
      openGetQueriesTctx: false,
    })
  },

  [actions.service.serviceSubmitQueries]: (state) => {
    logger.info("reducers", "serviceSubmitQueries")
    return Object.assign({}, state, {
      submitQueriesTctx: {
        fetching: true
      },
      openSubmitQueriesTctx: false,
      isFetching: true,
      isSubmitting: true,
    })
  },

  [actions.service.serviceCloseGetQueriesTctx]: (state) => {
    logger.info("reducers", "serviceCloseGetQueriesTctx")
    return Object.assign({}, state, {
      openGetQueriesTctx: false,
    })
  },

  [actions.service.serviceCloseSubmitQueriesTctx]: (state) => {
    logger.info("reducers", "serviceCloseSubmitQueriesTctx")
    return Object.assign({}, state, {
      openSubmitQueriesTctx: false,
    })
  },

  [actions.service.servicePostSuccess]: (state, action) => {
    logger.info("reducers", "servicePostSuccess", action.payload.action.type)
    let newState = Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
    })

    if (action.payload.action.type === 'SERVICE_SUBMIT_QUERIES') {
      Object.assign(newState, {
        isSubmitting: false,
      })
    }

    const actionType = action.payload.action.type
    const tctx = action.payload.data.Data.Tctx

    switch(actionType) {
      case 'SERVICE_GET_INDEX':
        newState.getIndexTctx = tctx
        break
      case 'SERVICE_GET_QUERIES':
        newState.getQueriesTctx = tctx
        newState.openGetQueriesTctx = true
        break
      case 'SERVICE_SUBMIT_QUERIES':
        newState.submitQueriesTctx = tctx
        newState.openSubmitQueriesTctx = true
        break
      default:
        break
    }
    if (tctx.StatusCode >= 300) {
      logger.error("reducers", "servicePostSuccess: newState", newState)
      // TODO handling tctx.Err, tctx.StatusCode
      return newState
    }

    if (action.payload.action.payload.isSync) {
      newState.syncAction = action.payload.action
    } else {
      newState.syncAction = null
    }

    let index = action.payload.data.Index
    if (index) {
      if (index.SyncDelay && index.SyncDelay > 1000) {
        newState.syncDelay = index.SyncDelay
      }
    }

    let service = action.payload.action.payload.serviceName
    let project = action.payload.action.payload.projectName
    if (project) {
      newState.projectServiceMap[project][service].isFetching = false
      newState.projectServiceMap[project][service].Index = action.payload.data.Index
      if (newState.projectServiceMap[project][service].Data) {
        for (let key in action.payload.data.Data) {
          newState.projectServiceMap[project][service].Data[key] = action.payload.data.Data[key]
        }
      } else {
        newState.projectServiceMap[project][service].Data = action.payload.data.Data
      }
    } else {
      newState.serviceMap[service].isFetching = false
      newState.serviceMap[service].Index = action.payload.data.Index
      if (newState.serviceMap[service].Data) {
        for (let key in action.payload.data.Data) {
          newState.serviceMap[service].Data[key] = action.payload.data.Data[key]
        }
      } else {
        newState.serviceMap[service].Data = action.payload.data.Data
      }
    }

    logger.info("reducers", "servicePostSuccess: newState", newState)
    return newState
  },

  [actions.service.servicePostFailure]: (state, action) => {
    let newState = Object.assign({}, defaultState, {
      isFetching: false,
      error: action.payload.error,
      payloadError: action.payload.payloadError,
    })

    if (action.payload.action.type === 'SERVICE_SUBMIT_QUERIES') {
      Object.assign(newState, {
        isSubmitting: false,
      })
    }
    return newState
  },
}, defaultState);
