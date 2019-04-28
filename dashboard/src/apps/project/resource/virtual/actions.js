import { createActions } from 'redux-actions';

export default createActions({
  RESORUCE_VIRTUAL_GET_INDEX: (projectName) => ({projectName: projectName}),
  RESORUCE_VIRTUAL_GET_INDEX_SUCCESS: (index) => ({index: index}),
  RESORUCE_VIRTUAL_GET_INDEX_FAILURE: (error) => ({error: error}),
})
