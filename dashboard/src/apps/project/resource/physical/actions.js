import { createActions } from 'redux-actions';

export default createActions({
  RESOURCE_PHYSICAL_GET_INDEX: (projectName) => ({projectName: projectName}),
  RESOURCE_PHYSICAL_GET_INDEX_SUCCESS: (index) => ({index: index}),
  RESOURCE_PHYSICAL_GET_INDEX_FAILURE: (error) => ({error: error}),

  // RESROUCE_SYNC_INDEX_STATE: (projectName, indexName) => ({projectName: projectName, indexName: indexName}),
  // RESROUCE_SYNC_INDEX_STATE_SUCCESS: (indexState) => ({indexState: indexState}),
  // RESROUCE_SYNC_INDEX_STATE_FAILURE: (error) => ({error: error}),
})
