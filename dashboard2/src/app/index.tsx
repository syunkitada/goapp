import logger from "../appui/src/lib/logger";
import data from "../appui/src/data";
import { IProvider } from "../appui/src/provider/IProvider";

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
        .then(res => {
            if (res.status !== 200) {
                logger.error("goapp.query.status !== 200", res.status, res);
                onError({ error: `UnexpectedStatus: ${res.status}` });
            }
            return res.json();
        })
        .then(payload => {
            logger.info("goapp.query.onSuccess", payload);
            onSuccess(payload.ResultMap);
        })
        .catch(error => {
            logger.error("goapp.query.onError", error);
            onError({ error });
        });
}

class Provider implements IProvider {
    getInitData(input: any): any {
        return {
            DefaultServiceName: "Home",
            DefaultProjectServiceName: "HomeProject",
            Logo: {
                Kind: "Text",
                Name: "Home"
            },
            LoginView: {
                Name: "Please Log In",
                Fields: [
                    {
                        Name: "User Name",
                        Key: "User",
                        Kind: "Text",
                        Required: true
                    },
                    {
                        Name: "Password",
                        Key: "Password",
                        Kind: "Password",
                        Required: true
                    },
                    {
                        Name: "Remember me",
                        Key: "rememberMe",
                        Kind: "Checkbox"
                    }
                ]
            }
        };
    }

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
                    input.onError({ error: result.Error });
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
        const { params, onError, onSuccess } = input;

        const serviceName = "Auth";
        const queries = [
            {
                Data: JSON.stringify(params),
                Name: "Login"
            }
        ];

        query({
            queries,
            serviceName,
            onSuccess: function (_input: any) {
                const result = _input.Login;
                if (result.Error && result.Error !== "") {
                    onError(result.Error);
                } else {
                    data.auth = result.Data;
                    onSuccess();
                }
            },
            onError: function (_input: any) {
                onError(_input);
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

    getServiceIndex(input: any): void {
        const { serviceName, projectName, location, onSuccess } = input;
        const that = this;
        let queries: any = [];

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

        let data = {};
        if (location && (location.DataQueries || location.WebSocketQuery)) {
            const tmpQueries: any[] = [];
            const tmpData = Object.assign(
                {},
                location.Params,
                location.ViewParams,
                location.SearchQueries
            );
            data = JSON.stringify(tmpData);
            if (location.DataQueries) {
                for (
                    let i = 0, len = location.DataQueries.length;
                    i < len;
                    i++
                ) {
                    tmpQueries.push({
                        Data: data,
                        Name: location.DataQueries[i]
                    });
                }
                queries = queries.concat(tmpQueries);
            }
        }

        query({
            queries,
            serviceName,
            projectName,
            onSuccess: function (_input: any) {
                let index: any;
                const resultData: any = {};
                const errors: any = [];
                for (let i = 0, len = queries.length; i < len; i++) {
                    const query = queries[i];
                    const result = _input[query.Name];
                    if (result.Error && result.Error !== "") {
                        errors.push({ Error: result.Error });
                        continue;
                    }
                    if (result.Code >= 100) {
                        errors.push({
                            Error: `UnexpectedCode: ${result.Code}`
                        });
                        continue;
                    }
                    if (
                        query.Name === "GetProjectServiceDashboardIndex" ||
                        query.Name === "GetServiceDashboardIndex"
                    ) {
                        index = result.Data.Index;
                        Object.assign(resultData, result.Data.Data);
                    } else {
                        Object.assign(resultData, result.Data);
                    }
                }
                if (errors.length > 0) {
                    input.onError({
                        errors
                    });
                    return;
                }

                if (location && location.WebSocketQuery) {
                    that.startWebSocket({
                        serviceName,
                        projectName,
                        location,
                        data,
                        onSuccess,
                        resultData,
                        index
                    });
                } else {
                    onSuccess({
                        data: resultData,
                        index: index
                    });
                }
            },
            onError: function (_input: any) {
                input.onError(_input);
            }
        });
    }

    getQueries(input: any): void {
        const { serviceName, projectName, location, onSuccess } = input;
        const that = this;

        const queryData = Object.assign(
            {},
            location.Params,
            location.ViewParams,
            location.SearchQueries
        );
        const data = JSON.stringify(queryData);

        const queries: any = [];
        if (!location.DataQueries) {
            if (location.WebSocketQuery) {
                that.startWebSocket({
                    serviceName,
                    projectName,
                    location,
                    data,
                    onSuccess,
                    resultData: {}
                });
            } else {
                onSuccess({});
            }
            return;
        }
        for (let i = 0, len = location.DataQueries.length; i < len; i++) {
            queries.push({
                Data: data,
                Name: location.DataQueries[i]
            });
        }

        query({
            queries,
            serviceName,
            projectName,
            onSuccess: function (_input: any) {
                const resultData: any = {};
                const errors: any = [];
                for (let i = 0, len = queries.length; i < len; i++) {
                    const query = queries[i];
                    const result = _input[query.Name];
                    if (result.Error && result.Error !== "") {
                        errors.push({ Error: result.Error });
                        continue;
                    }
                    if (result.Code >= 100) {
                        errors.push({
                            Error: `UnexpectedCode: ${result.Code}`
                        });
                        continue;
                    }
                    Object.assign(resultData, result.Data);
                }
                if (errors.length > 0) {
                    input.onError({
                        errors
                    });
                    return;
                }

                if (location.WebSocketQuery) {
                    that.startWebSocket({
                        serviceName,
                        projectName,
                        location,
                        data,
                        onSuccess,
                        resultData
                    });
                } else {
                    onSuccess({
                        data: resultData
                    });
                }
            },
            onError: function (_input: any) {
                input.onError(_input);
            }
        });
    }

    startWebSocket(input: any) {
        const {
            projectName,
            serviceName,
            location,
            data,
            onSuccess,
            resultData,
            index
        } = input;
        const queries = [
            {
                Data: data,
                Name: location.WebSocketQuery
            }
        ];

        const url: any = process.env.REACT_APP_AUTHPROXY_URL;
        const wsUrl: any = url.replace("http", "ws");
        const socket = new WebSocket(wsUrl + "/ws");

        let isInit = true;
        socket.onclose = event => {
            logger.info("index.startWebSocket.onclose", event);
        };
        socket.onerror = event => {
            logger.error("index.startWebSocket.onerror", event);
        };
        socket.onmessage = event => {
            const eventData = JSON.parse(event.data);
            if (isInit) {
                const newResultData = Object.assign(resultData, eventData);

                if (index) {
                    // getServiceindex
                    onSuccess({
                        index,
                        data: newResultData,
                        websocket: socket
                    });
                } else {
                    onSuccess({
                        data: newResultData,
                        websocket: socket
                    });
                }
                isInit = false;
                return;
            }
        };
        socket.onopen = event => {
            logger.info("index.startWebSocket.onopen", event);
            const initMessage = JSON.stringify({
                Project: projectName,
                Service: serviceName,
                Queries: queries
            });
            socket.send(initMessage);
        };
    }

    submitQueries(input: any): void {
        console.log("TODO submit_queries");
    }
}

const index = {
    Provider
};
export default index;
