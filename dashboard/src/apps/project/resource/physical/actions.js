import { createActions } from 'redux-actions';

export default createActions({
  RESOURCE_PHYSICAL_GET_INDEX: (projectName) => ({
    stateKey: 'index',
    serviceName: 'Resource',
    actionName: 'UserQuery',
    projectName: projectName,
    data: {
      queries: [
        {kind: "GetDatacenters"},
      ],
    },
  }),
  RESOURCE_PHYSICAL_GET_DATACENTER_INDEX: (projectName, datacenterName) => ({
    stateKey: 'datacenterIndex',
    serviceName: 'Resource',
    actionName: 'UserQuery',
    projectName: projectName,
    data: {
      queries: [
        {kind: "GetPhysicalResources", datacenterName: datacenterName},
        {kind: "GetFloors", datacenterName: datacenterName},
        {kind: "GetRacks", datacenterName: datacenterName},
      ],
    },
  }),

  RESOURCE_PHYSICAL_POST_SUCCESS: (action, data) => ({
    action: action,
    data: data
  }),
  RESOURCE_PHYSICAL_POST_FAILURE: (action, error, payloadError) => ({
    action: action,
    error: error,
    payloadError: payloadError
  }),
})
