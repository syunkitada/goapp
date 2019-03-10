import { handleActions } from 'redux-actions';

import actions from '../../../../actions'

const defaultState = {
  isSyncState: false,
  isFetching: false,
  error: null,
  index: null,
};

export default handleActions({
  [actions.resourcePhysical.resourcePhysicalGetIndex]: (state) => {
    console.log("resourcePhysical.reducers.syncState")
    return Object.assign({}, state, {
      isFetching: true,
    })
  },
  [actions.resourcePhysical.resourcePhysicalGetIndexSuccess]: (state, action) => {
    console.log("resourcePhysical.reducers.syncStateSuccess", state, action)
    return Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
      index: action.payload.index,
    })
  },
  [actions.resourcePhysical.resourcePhysicalGetIndexFailure]: (state, action) => {
    console.log("resourcePhysical.reducers.syncStateFailure")
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
}, defaultState);
