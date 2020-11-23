import Dashboard from "../../components/core/Dashboard";
import auth from "../../apps/auth";
import client from "../../client";
import data from "../../data";
import locationData from "../../data/locationData";
import Index from "../../components/Index";

function init() {
    const { serviceName, projectName } = locationData.getServiceParams();

    const location = { Path: ["Root"] };
    const tmpLocationData = locationData.getLocationData();
    let initLocation = false;
    if (tmpLocationData.Path) {
        initLocation = true;
    }
    client.get_service_index({
        serviceName,
        projectName,
        initLocation,
        location,
        onSuccess: function (input: any) {
            if (input.Index.DefaultRoute.Path) {
                location.Path = input.Index.DefaultRoute.Path;
            }
            data.service = {
                location,
                data: input.Data,
                rootView: input.Index.View
            };

            Index.Render({
                id: "root-content",
                View: input.Index.View
            });
        },
        onError: function (input: any) {
            console.log("onError", input);
        }
    });

    Dashboard.Render({
        id: "root",
        logout: auth.logout,
        onClickService: function (input: any) {
            const { serviceName, projectName } = input;
            const initLocation = false;
            const location = { Path: ["Root"] };

            locationData.setServiceParams(input);

            client.get_service_index({
                serviceName,
                projectName,
                initLocation,
                location,
                onSuccess: function (input: any) {
                    if (input.Index.DefaultRoute.Path) {
                        location.Path = input.Index.DefaultRoute.Path;
                    }
                    data.service = {
                        location,
                        data: input.Data,
                        rootView: input.Index.View
                    };

                    Index.Render({
                        id: "root-content",
                        View: input.Index.View
                    });
                },
                onError: function (input: any) {
                    console.log("onError", input);
                }
            });
        }
    });
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
    const { location, View } = input;
    const { serviceName, projectName } = locationData.getServiceParams();
    const nextView = getViewFromPath(data.service.rootView, location.Path);
    // const subPathMap = location.SubPathMap
    // location.DataQueries = index.DataQueries
    // location.WebSocketQuery = index.WebSocketQuery
    // const params = Object.assign(
    // )
    location.DataQueries = nextView.DataQueries;
    console.log("DEBUG getQueries", input);

    locationData.setLocationData(location);
    $("#root-content-progress").html('<div class="indeterminate"></div>');

    client.get_queries({
        serviceName,
        projectName,
        location,
        onSuccess: function (input: any) {
            $("#root-content-progress").html(
                '<div class="determinate" style="width: 0%"></div>'
            );

            data.service.data = Object.assign(data.service.data, input.data);
            Index.Render({
                id: "root-content",
                View: data.service.rootView
            });
            console.log("DEBUG getQueries onSuccess", data.service.data);
        },
        onError: function (input: any) {
            console.log("DEBUG getQueries onError", input);
        }
    });
}

const index = {
    init,
    getQueries
};
export default index;
