import * as React from "react";

import IndexForm from "./forms/IndexForm";
import SearchForm from "./forms/SearchForm";
import RoutePanels from "./panels/RoutePanels";
import Panes from "./panes/Panes";
import IndexTable from "./tables/IndexTable";
import Tabs from "./tabs/Tabs";
import IndexView from "./views/IndexView";

import logger from "../lib/logger";

function renderIndex(routes, index) {
    if (!index) {
        return <div>Not Found</div>;
    }
    logger.info("Index", "renderIndex:", index.Kind, routes);
    switch (index.Kind) {
        case "Msg":
            return <div>{index.Name}</div>;
        case "RoutePanels":
            return (
                <RoutePanels
                    render={renderIndex}
                    routes={routes}
                    index={index}
                />
            );
        case "RouteTabs":
            return <Tabs render={renderIndex} routes={routes} index={index} />;
        case "RoutePanes":
            return <Panes render={renderIndex} routes={routes} index={index} />;
        case "Table":
            return (
                <IndexTable
                    render={renderIndex}
                    routes={routes}
                    index={index}
                />
            );
        case "View":
            return (
                <IndexView render={renderIndex} routes={routes} index={index} />
            );
        case "SearchForm":
            return <SearchForm routes={routes} index={index} />;
        case "Form":
            return <IndexForm routes={routes} index={index} />;
        default:
            return <div>Unsupported Kind: {index.Kind}</div>;
    }
}

export default {
    renderIndex
};
