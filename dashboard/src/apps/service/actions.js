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

  SERVICE_START_BACKGROUND_SYNC: (queries, params) => {
    let dataQueries = [];
    for (let i = 0, len = queries.length; i < len; i ++) {
      dataQueries.push({Kind: queries[i], StrParams: params})
    }
    return {
      stateKey: 'index',
      serviceName: params.service,
      actionName: 'UserQuery',
      projectName: params.project,
      data: {
        queries: dataQueries,
      },
    }
  },

  SERVICE_STOP_BACKGROUND_SYNC: () => {},

  SERVICE_GET_QUERIES: (queries, params) => {
    let dataQueries = [];
    for (let i = 0, len = queries.length; i < len; i ++) {
      dataQueries.push({Kind: queries[i], StrParams: params})
    }
    return {
      stateKey: 'index',
      serviceName: params.service,
      actionName: 'UserQuery',
      projectName: params.project,
      data: {
        queries: dataQueries,
      },
    }
  },

  SERVICE_SUBMIT_QUERIES: (action, fieldMap, targets, params) => {
    let kind = action.Name + action.DataKind
    let dataQueries = [];
    let strParams = Object.assign({}, params)
    let numParams = {}
    for (let key in fieldMap) {
      let field = fieldMap[key]
      switch (field.Type) {
        case "text":
          strParams[key] = field.value
          break
        default:
          break
      }
    }

    for (let i = 0, len = targets.length; i < len; i ++) {
      dataQueries.push({Kind: kind, StrParams: strParams, NumParams: numParams})
    }
    return {
      stateKey: 'index',
      serviceName: params.service,
      actionName: 'UserQuery',
      projectName: params.project,
      data: {
        queries: dataQueries,
      },
    }
  },

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
