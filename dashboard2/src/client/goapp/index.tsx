import logger from "../../lib/logger";
import data from "../../data";
import { IClient } from "../../client/IClient";

function query(input: any) {
    const { projectName, serviceName, queries, onSuccess, onError } = input;

    const body = JSON.stringify({
        Project: projectName,
        Queries: queries,
        Service: serviceName
    });
    logger.info("goapp.query", body);

    return fetch(process.env.REACT_APP_AUTHPROXY_URL + "/q", {
        body,
        credentials: "include",
        method: "POST",
        mode: "cors"
    })
        .then(res => res.json())
        .then(payload => {
            console.log("DEBUG payload", payload);
            onSuccess(payload.ResultMap);
        })
        .catch(error => {
            console.log("DEBUG error", error);
            onError({ error });
        });
}

class Client implements IClient {
    loginWithToken(input: any): void {
        const serviceName = "Auth";
        const queries = [
            {
                Data: JSON.stringify({}),
                Name: "LoginWithToken"
            }
        ];

        query({
            queries,
            serviceName,
            onSuccess: function (_input: any) {
                const result = _input.LoginWithToken;
                if (result.Error && result.Error !== "") {
                    input.onError(result.Error);
                } else {
                    data.auth = result.Data;
                    input.onSuccess();
                }
            },
            onError: function (_input: any) {
                input.onError(_input);
            }
        });
    }

    login(input: any): void {
        const { userName, password } = input;

        const serviceName = "Auth";
        const queries = [
            {
                Data: JSON.stringify({ User: userName, Password: password }),
                Name: "Login"
            }
        ];

        query({
            queries,
            serviceName,
            onSuccess: function (_input: any) {
                const result = _input.Login;
                if (result.Error && result.Error !== "") {
                    input.onError(result.Error);
                } else {
                    data.auth = result.Data;
                    input.onSuccess();
                }
            },
            onError: function (_input: any) {
                input.onError(_input);
            }
        });
    }

    logout(input: any): void {
        const serviceName = "Auth";
        const queries = [
            {
                Data: JSON.stringify({}),
                Name: "Logout"
            }
        ];

        query({
            queries,
            serviceName,
            onSuccess: function (_input: any) {
                const result = _input.Logout;
                if (result.Error && result.Error !== "") {
                    input.onError(result.Error);
                } else {
                    input.onSuccess();
                }
            },
            onError: function (_input: any) {
                input.onError(_input);
            }
        });
    }

    get_service_index(input: any): void {
        const { serviceName, projectName, location, initLocation } = input;
        let queries: any = null;
        if (projectName) {
            queries = [
                {
                    Data: JSON.stringify({
                        Name: serviceName
                    }),
                    Name: "GetProjectServiceDashboardIndex"
                }
            ];
        } else {
            queries = [
                {
                    Data: JSON.stringify({
                        Name: serviceName
                    }),
                    Name: "GetServiceDashboardIndex"
                }
            ];
        }

        if (initLocation && location && location.DataQueries) {
            const tmpQueries: any[] = [];
            const tmpData = Object.assign(
                {},
                location.Params,
                location.SearchQueries
            );
            const data = JSON.stringify(tmpData);
            for (let i = 0, len = location.DataQueries.length; i < len; i++) {
                tmpQueries.push({
                    Data: data,
                    Name: location.DataQueries[i]
                });
            }
            queries = queries.concat(tmpQueries);
        }

        query({
            queries,
            serviceName,
            onSuccess: function (_input: any) {
                let result: any = null;
                if (projectName) {
                    result = _input.GetProjectServiceDashboardIndex;
                } else {
                    result = _input.GetServiceDashboardIndex;
                }
                if (result.Error && result.Error !== "") {
                    input.onError(result.Error);
                } else {
                    input.onSuccess(result.Data);
                }
            },
            onError: function (_input: any) {
                console.log("error", _input);
            }
        });
    }
}

const index = {
    query,
    Client
};
export default index;
