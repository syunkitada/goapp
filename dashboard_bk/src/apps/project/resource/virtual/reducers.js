import { handleActions } from 'redux-actions';

import actions from '../../../../actions'

const defaultState = {
  isSyncState: false,
  isFetching: false,
  error: null,
  monitor: null,
  indexState: null,
};

export default handleActions({
  [actions.monitor.monitorSyncState]: (state) => {
    console.log("monitor.reducers.syncState")
    return Object.assign({}, state, {
      isFetching: true,
    })
  },
  [actions.monitor.monitorSyncStateSuccess]: (state, action) => {
    console.log("monitor.reducers.syncStateSuccess", state, action)
    return Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
      monitor: action.payload.monitor,
    })
  },
  [actions.monitor.monitorSyncStateFailure]: (state, action) => {
    console.log("monitor.reducers.syncStateFailure")
    return Object.assign({}, defaultState, {
      isFetching: false,
      error: action.payload.error,
    })
  },
  [actions.monitor.monitorSyncIndexState]: (state) => {
    console.log("monitor.reducers.syncIndexState")
    return Object.assign({}, state, {
      isFetching: true,
    })
  },
  [actions.monitor.monitorSyncIndexStateSuccess]: (state, action) => {
    console.log("monitor.reducers.syncIndexStateSuccess")
    return Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
      indexState: action.payload.indexState,
    })
  },
  [actions.monitor.monitorSyncIndexStateFailure]: (state, action) => Object.assign({}, state, {
    isFetching: false,
    redirectToReferrer: true,
    indexState: action.payload.indexState,
  }),
}, defaultState);
