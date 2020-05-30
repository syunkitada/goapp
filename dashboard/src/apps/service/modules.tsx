import fetch from "cross-fetch";

import logger from "../../lib/logger";

interface IResponse {
    result: any;
    error: any;
}

function post({
    serviceName,
    actionName,
    projectName,
    queries
}): Promise<IResponse> {
    const body = JSON.stringify({
        Project: projectName,
        Queries: queries,
        Service: serviceName
    });
    logger.info("modules.post", body);

    return fetch(process.env.REACT_APP_AUTHPROXY_URL + "/q", {
        body,
        credentials: "include",
        method: "POST",
        mode: "cors"
    })
        .then(resp => {
            logger.info("modules.post", "success", body, resp);
            if (!resp.ok) {
                return resp.json().then(payload => {
                    const result: IResponse = {
                        error: {
                            err: payload.Error,
                            errCode: resp.status
                        },
                        result: null
                    };
                    return result;
                });
            }

            return resp.json().then(payload => {
                const result: IResponse = {
                    error: null,
                    result: payload
                };
                return result;
            });
        })
        .catch(error => {
            logger.error("modules.post", "fail", body, error);
            const result: IResponse = {
                error: {
                    err: error
                },
                result: null
            };
            return result;
        });
}

export default {
    post
};
