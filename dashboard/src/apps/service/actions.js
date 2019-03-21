import { createActions } from 'redux-actions';

export default createActions({
  SERVICE_GET_INDEX: (params) => ({
    stateKey: 'index',
    serviceName: params.service,
    actionName: 'UserQuery',
    projectName: params.project,
    data: {
      queries: [
        {kind: "GetIndex", params: params},
      ],
    },
  }),

  SERVICE_POST_SUCCESS: (action, data) => ({
    action: action,
    data: data
  }),
  SERVICE_POST_FAILURE: (action, error, payloadError) => ({
    action: action,
    error: error,
    payloadError: payloadError
  }),
})
