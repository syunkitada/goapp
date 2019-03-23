import { handleActions } from 'redux-actions';

import actions from '../../actions'

const defaultState = {
  isFetching: false,
  error: null,
  payloadError: null,
  index: {Index: null, Data: null},
  datacenterIndex: null,
};

export default handleActions({
  [actions.service.serviceGetIndex]: (state) => {
    console.log("action.serviceGetIndex")
    return Object.assign({}, state, {
      isFetching: true,
    })
  },
  [actions.service.serviceGetQueries]: (state) => {
    console.log("action.serviceGetQueries")
    return Object.assign({}, state, {
      isFetching: true,
    })
  },

  [actions.service.servicePostSuccess]: (state, action) => {
    console.log("DEBUG: servicePostSuccess: ", action.payload.action.type)
    console.log(action)
    let newState = Object.assign({}, state, {
      isFetching: false,
      redirectToReferrer: true,
    })
    let stateKey = action.payload.action.payload.stateKey
    let data = action.payload.data
    newState[stateKey].Index = data.Index
    let newData = Object.assign({}, newState[stateKey].Data, data.Data)
    newState[stateKey].Data = newData
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
