import { handleActions } from 'redux-actions';

import actions from '../../../../actions'

const defaultState = {
  isFetching: false,
  error: null,
  payloadError: null,
  index: null,
  datacenterIndex: null,
};

export default handleActions({
  [actions.resourcePhysical.resourcePhysicalGetIndex]: (state) => {
    return Object.assign({}, state, {
      isFetching: true,
    })
  },
  [actions.resourcePhysical.resourcePhysicalGetDatacenterIndex]: (state) => {
    return Object.assign({}, state, {
      isFetching: true,
    })
  },

  [actions.resourcePhysical.resourcePhysicalPostSuccess]: (state, action) => {
    let newState = Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
    })
    newState[action.payload.action.payload.stateKey] = action.payload.data
    console.log("Debug: resourcePhysicalPostSuccess")
    console.log(newState)
    return newState
  },
  [actions.resourcePhysical.resourcePhysicalPostFailure]: (state, action) => {
    return Object.assign({}, defaultState, {
      isFetching: false,
      error: action.payload.error,
      payloadError: action.payload.payloadError,
    })
  },
}, defaultState);
