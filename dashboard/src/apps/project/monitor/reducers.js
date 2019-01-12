import { handleActions } from 'redux-actions';

import actions from '../../../actions'

const defaultState = {
  isSyncState: false,
  isFetching: false,
  error: null,
  monitor: null,
};

export default handleActions({
  [actions.monitor.syncState]: (state) => Object.assign({}, state, {
    isFetching: true,
  }),
  [actions.monitor.syncStateSuccess]: (state, action) => Object.assign({}, state, {
    isFetching: false,
    redirectToReferrer: true,
    monitor: action.payload.monitor,
  }),
  [actions.monitor.syncStateFailure]: (state, action) => Object.assign({}, defaultState, {
    isFetching: false,
    error: action.payload.error,
  }),
}, defaultState);
