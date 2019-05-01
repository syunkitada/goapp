import fetch from 'cross-fetch';

interface IResponse {
  payload: any;
  error: any;
}

function post({
  serviceName,
  actionName,
  projectName,
  queries,
}): Promise<IResponse> {
  const body = JSON.stringify({
    Action: {
      ServiceName: serviceName,
      Name: actionName,
      ProjectName: projectName,
      Queries: queries,
    },
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/' + serviceName, {
    method: 'POST',
    credentials: 'include',
    mode: 'cors',
    body: body,
  })
    .then(function(resp) {
      if (!resp.ok) {
        return resp.json().then(function(payload) {
          const result: IResponse = {
            payload: null,
            error: {
              errCode: resp.status,
              err: payload.Err,
            },
          };
          return result;
        });
      }

      return resp.json().then(function(payload) {
        const result: IResponse = {
          payload: payload,
          error: null,
        };
        return result;
      });
    })
    .catch(function(error) {
      const result: IResponse = {
        payload: null,
        error: {
          err: error,
        },
      };
      return result;
    });
}

export default {
  post,
};
