import { handleActions } from 'redux-actions';

import actions from '../../../actions'

const defaultState = {
  isSyncState: false,
  isFetching: false,
  error: null,
  resource: null,
};

export default handleActions({
  [actions.resource.syncState]: (state) => Object.assign({}, state, {
    isFetching: true,
  }),
  [actions.resource.syncStateSuccess]: (state, action) => Object.assign({}, state, {
    isFetching: false,
    redirectToReferrer: true,
    resource: action.payload.resource,
  }),
  [actions.resource.syncStateFailure]: (state, action) => Object.assign({}, defaultState, {
    isFetching: false,
    error: action.payload.error,
  }),
}, defaultState);
