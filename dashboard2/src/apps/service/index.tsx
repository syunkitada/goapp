import Dashboard from "../../components/core/Dashboard";
import auth from "../../apps/auth";
import client from "../../client";
import data_utils from "../../lib/data_utils";
import Index from "../../components/Index";

function init(input: any) {
    let { serviceName, projectName } = input;
    if (!serviceName) {
        // Called on init
        const params = data_utils.getServiceParams();
        serviceName = params.serviceName;
        projectName = params.projectName;
    } else {
        // Called on click LeftSidebar link
        data_utils.setServiceParams(input);
        data_utils.setLocationData({});
    }

    const location = { Path: ["Root"] };
    const locationData = data_utils.getLocationData();
    let initLocation = false;
    if (locationData.Path) {
        initLocation = true;
    }
    client.get_service_index({
        serviceName,
        projectName,
        initLocation,
        location,
        onSuccess: function (input: any) {
            Index.Render({
                id: "root-content",
                View: input.Index.View
            });
        },
        onError: function (input: any) {
            console.log("onError", input);
        }
    });
    console.log("DEBUG service", serviceName, projectName);

    Dashboard.Render({ id: "root", logout: auth.logout });
}

const index = {
    init
};
export default index;
