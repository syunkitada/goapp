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
            data.service = {
                location,
                data: input.Data
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
                    data.service = {
                        location,
                        data: input.Data
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

const index = {
    init
};
export default index;
