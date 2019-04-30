import { createActions } from 'redux-actions';

export default createActions({
  HOME_SYNC_STATE: () => ({}),
  HOME_SUCCESS_SYNC_STATE: (data) => ({data: data}),
  HOME_FAILURE_SYNC_STATE: (error) => ({error: error}),
})
