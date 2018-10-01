import { handleActions } from 'redux-actions';

import actions from '../../actions'

const defaultState = {
  isSyncState: false,
  isFetching: false,
  data: null,
  error: null,
};

export default handleActions({
  [actions.home.homeSyncState]: (state) => Object.assign({}, state, {
    isSyncState: true,
    isFetching: true,
  }),
  [actions.home.homeSuccessSyncState]: (state, action) => Object.assign({}, state, {
    isFetching: false,
    data: action.payload.data,
  }),
  [actions.home.homeFailedSyncState]: (state, action) => Object.assign({}, state, {
    isFetching: false,
    error: action.payload.error,
  }),
}, defaultState);
