import Toast from "../../components/toast/Toast";
import Dashboard from "../../components/core/Dashboard";
import auth from "../../apps/auth";
import provider from "../../provider";
import data from "../../data";
import locationData from "../../data/locationData";
import Index from "../../components/Index";
import logger from "../../lib/logger";

function getServiceIndex(input: any) {
    let { location, serviceName, projectName } = input;

    function getServiceIndexOnSuccess(input: any) {
        const { index } = input;
        if (!location) {
            let path = ["Root"];
            if (index.DefaultRoute.Path) {
                path = index.DefaultRoute.Path;
            }
            location = {
                Path: path,
                DataQueries: [],
                SearchQueries: {},
                WebSocketQuery: {},
                Params: {},
                ViewParams: {}
            };
        }
        locationData.setLocationData(location);

        data.service = {
            rootView: index.View,
            data: input.data,
            websocketMap: {}
        };

        const nextView = getViewFromPath(index.View, location.Path);
        if (nextView.WebSocketKey && input.websocket) {
            data.service.websocketMap[location.Path.join(".")] =
                input.websocket;
        }

        Dashboard.NavPath.Render({
            location
        });

        Index.Render({
            id: "root-content",
            View: index.View
        });

        Dashboard.RootContentProgress.StopProgress();
    }

    locationData.setServiceParams(input);

    Dashboard.RootContentProgress.StartProgress();
    provider.getServiceIndex({
        serviceName,
        projectName,
        location,
        onSuccess: getServiceIndexOnSuccess,
        onError: function (input: any) {
            logger.error("service.init.onClickService.onError", input);
            Toast.Error(input);
        }
    });
}

function init() {
    const { serviceName, projectName } = locationData.getServiceParams();

    let location = null;
    const tmpLocationData = locationData.getLocationData();
    if (tmpLocationData.Path) {
        location = tmpLocationData;
    }

    Dashboard.Render({
        id: "root",
        logout: auth.logout,
        onClickService: function (input: any) {
            const { serviceName, projectName } = input;
            getServiceIndex({ serviceName, projectName });
        }
    });

    getServiceIndex({
        serviceName,
        projectName,
        location
    });
}

function gotoService(input: any) {
    const { serviceName } = input;
    const services = $(`.${Dashboard.serviceLinkClass}`);
    for (let i = 0, len = services.length; i < len; i++) {
        const service = $(services[i]);
        if (service.text() == serviceName) {
            service.click();
            break;
        }
    }
}

function getViewFromPath(View: any, path: any): any {
    if (View.Children) {
        for (let i = 0, len = View.Children.length; i < len; i++) {
            const child = View.Children[i];
            if (child.Name !== path[0]) {
                continue;
            }
            return getViewFromPath(child, path.slice(1));
        }
    }

    return View;
}

function getQueries(input: any) {
    const { location, view } = input;
    const { serviceName, projectName } = locationData.getServiceParams();
    const nextView = getViewFromPath(data.service.rootView, location.Path);

    if (nextView.WebSocketQuery) {
        location.WebSocketQuery = nextView.WebSocketQuery;
    } else {
        location.WebSocketQuery = null;
    }

    if (!nextView.DataQueries) {
        if (view) {
            Index.Render(view);
        } else {
            Index.Render({
                id: "root-content",
                View: data.service.rootView
            });
        }
        console.log("DEBUG location", location);
        locationData.setLocationData(location);
        Dashboard.NavPath.Render({
            location
        });
        return;
    }

    location.DataQueries = nextView.DataQueries;

    if (view && view.View.ViewParams) {
        location.ViewParams = view.View.ViewParams;
    } else {
        location.ViewParams = {};
    }

    logger.info("service.getQueries", location, view, nextView);

    locationData.setLocationData(location);
    Dashboard.RootContentProgress.StartProgress();

    provider.getQueries({
        serviceName,
        projectName,
        location,
        onSuccess: function (_input: any) {
            Dashboard.RootContentProgress.StopProgress();

            data.service.data = Object.assign(data.service.data, _input.data);

            const pathKey = location.Path.join(".");
            if (nextView.WebSocketKey && _input.websocket) {
                data.service.websocketMap[pathKey] = _input.websocket;
            }
            for (const [key, value] of Object.entries(
                data.service.websocketMap
            )) {
                if (key !== pathKey) {
                    data.service.websocketMap[key].close();
                    delete data.service.websocketMap[key];
                }
            }

            Dashboard.NavPath.Render({
                location
            });

            if (view) {
                Index.Render(view);
            } else {
                Index.Render({
                    id: "root-content",
                    View: data.service.rootView
                });
            }
            logger.info("getQueries.onSuccess", _input);
        },
        onError: function (_input: any) {
            logger.error("getQueries.onError", _input);
            Toast.Error(_input);
        }
    });
}

function submitQueries(input: any) {
    const { queries, location, params, onSuccess } = input;
    const { serviceName, projectName } = locationData.getServiceParams();

    Dashboard.RootContentProgress.StartProgress();

    provider.submitQueries({
        serviceName,
        projectName,
        queries,
        location,
        params,
        onSuccess: function (_input: any) {
            logger.info("submitQueries.onInfo", input, _input);

            Dashboard.RootContentProgress.StopProgress();

            data.service.data = Object.assign(data.service.data, _input.data);
            onSuccess();
        },
        onError: function (_input: any) {
            logger.error("submitQueries.onError", input, _input);
            Toast.Error(input);
        }
    });
}

const index = {
    init,
    getViewFromPath,
    getQueries,
    gotoService,
    submitQueries
};
export default index;
