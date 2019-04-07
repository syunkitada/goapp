import { handleActions } from 'redux-actions';

import actions from '../../actions'

const defaultState = {
  isFetching: false,
  isSubmitting: false,
  isSubmitSuccess: false,
  isSubmitFailed: false,
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
    console.log("DEBUG serviceGetIndex")
    let service = action.payload.serviceName
    let project = action.payload.projectName
    let newState = Object.assign({}, state, {
      serviceName: service,
      projectName: project,
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

    console.log(newState)

    return newState;
  },

  [actions.service.serviceStartBackgroundSync]: (state) => {
    console.log("action.serviceStartBackgroundSync")
    return Object.assign({}, state, {
      isFetching: true,
    })
  },

  [actions.service.serviceStopBackgroundSync]: (state) => {
    console.log("action.serviceStopBackgroundSync")
    return Object.assign({}, state, {
      isFetching: false,
    })
  },

  [actions.service.serviceGetQueries]: (state) => {
    console.log("action.serviceGetQueries")
    return Object.assign({}, state, {
      isFetching: true,
    })
  },

  [actions.service.serviceSubmitQueries]: (state) => {
    console.log("action.serviceSubmitQueries")
    return Object.assign({}, state, {
      isFetching: true,
      isSubmitting: true,
      isSubmitSuccess: false,
      isSubmitFailed: false,
    })
  },

  [actions.service.servicePostSuccess]: (state, action) => {
    console.log("DEBUG: servicePostSuccess: ", action.payload.action.type)
    console.log(action.payload.action)
    let newState = Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
    })

    if (action.payload.action.type === 'SERVICE_SUBMIT_QUERIES') {
      Object.assign(newState, {
        isSubmitting: false,
        isSubmitSuccess: true,
        isSubmitFailed: false,
      })
    }

    let tctx = action.payload.data.Data.Tctx
    console.log(tctx)
    // TODO handling tctx.Err, tctx.StatusCode

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

    console.log(newState)
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
        isSubmitSuccess: false,
        isSubmitFailed: true,
      })
    }
    return newState
  },
}, defaultState);
