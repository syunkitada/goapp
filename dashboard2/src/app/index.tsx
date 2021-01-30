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
                console.log("Failed", res);
            }
            return res.json();
        })
        .then(payload => {
            console.log("DEBUG payload", payload);
            onSuccess(payload.ResultMap);
        })
        .catch(error => {
            console.log("DEBUG error", error);
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

    getLoginView(input: any): any {
        return {
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

        if (
            location &&
            location.DataQueries &&
            location.DataQueries.length > 0
        ) {
            const tmpQueries: any[] = [];
            const tmpData = Object.assign(
                {},
                location.Params,
                location.ViewParams,
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
            projectName,
            onSuccess: function (_input: any) {
                let index: any;
                const data: any = {};
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
                        Object.assign(data, result.Data.Data);
                    } else {
                        Object.assign(data, result.Data);
                    }
                }
                if (errors.length > 0) {
                    input.onError({
                        errors
                    });
                    return;
                }
                input.onSuccess({
                    Data: data,
                    Index: index
                });
            },
            onError: function (_input: any) {
                console.log("error", _input);
            }
        });
    }

    getQueries(input: any): void {
        const { serviceName, projectName, location } = input;

        const queryData = Object.assign(
            {},
            location.Params,
            location.ViewParams,
            location.SearchQueries
        );
        const data = JSON.stringify(queryData);

        const queries: any = [];
        if (!location.DataQueries) {
            input.onSuccess({});
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
                const data: any = {};
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
                    Object.assign(data, result.Data);
                }
                if (errors.length > 0) {
                    input.onError({
                        errors
                    });
                    return;
                }
                input.onSuccess({
                    data
                });
            },
            onError: function (_input: any) {
                input.onError(_input);
            }
        });
    }

    submitQueries(input: any): void {
        console.log("DEBUG submit_queries");
    }
}

const index = {
    Provider
};
export default index;
