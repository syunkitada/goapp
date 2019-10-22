import fetch from 'cross-fetch';

interface IResponse {
  result: any;
  error: any;
}

function post({
  serviceName,
  actionName,
  projectName,
  queries,
}): Promise<IResponse> {
  const body = JSON.stringify({
    Project: projectName,
    Queries: queries,
    Service: serviceName,
  });

  return fetch(process.env.REACT_APP_AUTHPROXY_URL + '/q', {
    body,
    credentials: 'include',
    method: 'POST',
    mode: 'cors',
  })
    .then(resp => {
      if (!resp.ok) {
        return resp.json().then(payload => {
          const result: IResponse = {
            error: {
              err: payload.Error,
              errCode: resp.status,
            },
            result: null,
          };
          return result;
        });
      }

      return resp.json().then(payload => {
        const result: IResponse = {
          error: null,
          result: payload,
        };
        return result;
      });
    })
    .catch(error => {
      const result: IResponse = {
        error: {
          err: error,
        },
        result: null,
      };
      return result;
    });
}

export default {
  post,
};
