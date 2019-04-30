import { createActions } from 'redux-actions';

export default createActions({
  RESOURCE_SYNC_STATE: (projectName) => ({projectName: projectName}),
  RESOURCE_SYNC_STATE_SUCCESS: (resource) => ({resource: resource}),
  RESOURCE_SYNC_STATE_FAILURE: (error) => ({error: error}),
})
