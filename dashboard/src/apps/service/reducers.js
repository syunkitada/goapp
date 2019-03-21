import { handleActions } from 'redux-actions';

import actions from '../../actions'

const defaultState = {
  isFetching: false,
  error: null,
  payloadError: null,
  index: null,
  datacenterIndex: null,
};

export default handleActions({
  [actions.service.serviceGetIndex]: (state) => {
    console.log("action.serviceGetIndex")
    return Object.assign({}, state, {
      isFetching: true,
    })
  },

  [actions.service.servicePostSuccess]: (state, action) => {
    let newState = Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
    })
    newState[action.payload.action.payload.stateKey] = action.payload.data
    console.log("Debug: servicePostSuccess")
    console.log(newState)
    return newState
  },
  [actions.service.servicePostFailure]: (state, action) => {
    return Object.assign({}, defaultState, {
      isFetching: false,
      error: action.payload.error,
      payloadError: action.payload.payloadError,
    })
  },
}, defaultState);
