import { createActions } from 'redux-actions';

export default createActions({
  SERVICE_GET_INDEX: (params) => ({
    stateKey: 'index',
    serviceName: params.service,
    actionName: 'UserQuery',
    projectName: params.project,
    queries: [
      {Kind: "GetIndex", StrParams: params},
    ],
  }),

  SERVICE_START_BACKGROUND_SYNC: () => {},

  SERVICE_STOP_BACKGROUND_SYNC: () => {},

  SERVICE_GET_QUERIES: (queries, isSync, params) => {
    let dataQueries = [];
    for (let i = 0, len = queries.length; i < len; i ++) {
      dataQueries.push({Kind: queries[i], StrParams: params})
    }
    return {
      stateKey: 'index',
      serviceName: params.service,
      actionName: 'UserQuery',
      projectName: params.project,
      queries: dataQueries,
      isSync: isSync,
    }
  },

  SERVICE_SUBMIT_QUERIES: (action, fieldMap, targets, params) => {
    let kind = action.Name + action.DataKind
    let dataQueries = [];
    let strParams = Object.assign({}, params)
    let numParams = {}

    let spec = Object.assign({}, params)
    for (let key in fieldMap) {
      let field = fieldMap[key]
      spec[key] = field.value
    }
    let specsStr = JSON.stringify([spec])
    strParams['Specs'] = specsStr

    if (targets) {
      for (let i = 0, len = targets.length; i < len; i ++) {
        let target = targets[i]
        strParams.Target = target
        dataQueries.push({Kind: kind, StrParams: strParams, NumParams: numParams})
      }
    } else {
      dataQueries.push({Kind: kind, StrParams: strParams, NumParams: numParams})
    }

    return {
      stateKey: 'index',
      serviceName: params.service,
      actionName: 'UserQuery',
      projectName: params.project,
      queries: dataQueries,
    }
  },

  SERVICE_CLOSE_GET_QUERIES_TCTX: () => {},

  SERVICE_CLOSE_SUBMIT_QUERIES_TCTX: () => {},

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
