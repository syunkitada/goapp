import { createActions } from 'redux-actions';

export default createActions({
  MONITOR_SYNC_STATE: (projectName) => ({projectName: projectName}),
  MONITOR_SYNC_STATE_SUCCESS: (monitor) => ({monitor: monitor}),
  MONITOR_SYNC_STATE_FAILURE: (error) => ({error: error}),

  MONITOR_SYNC_INDEX_STATE: (projectName, indexName) => ({projectName: projectName, indexName: indexName}),
  MONITOR_SYNC_INDEX_STATE_SUCCESS: (indexState) => ({indexState: indexState}),
  MONITOR_SYNC_INDEX_STATE_FAILURE: (error) => ({error: error}),
})
