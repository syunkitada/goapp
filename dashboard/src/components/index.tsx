import * as React from "react";

import IndexForm from "./forms/IndexForm";
import SearchForm from "./forms/SearchForm";
import RoutePanels from "./panels/RoutePanels";
import Panes from "./panes/Panes";
import IndexTable from "./tables/IndexTable";
import RouteTabs from "./tabs/RouteTabs";
import IndexView from "./views/IndexView";

import logger from "../lib/logger";

function renderIndex(index) {
    if (!index) {
        return <div>Not Found</div>;
    }
    logger.info("Index", "renderIndex:", index.Kind);
    switch (index.Kind) {
        case "Msg":
            return <div>{index.Name}</div>;
        case "RoutePanels":
            return <RoutePanels render={renderIndex} index={index} />;
        case "RouteTabs":
            return <RouteTabs render={renderIndex} index={index} />;
        case "RoutePanes":
            return <Panes render={renderIndex} index={index} />;
        case "Table":
            return <IndexTable render={renderIndex} index={index} />;
        case "View":
            return <IndexView render={renderIndex} index={index} />;
        case "SearchForm":
            return <SearchForm index={index} />;
        case "Form":
            return <IndexForm index={index} />;
        default:
            return <div>Unsupported Kind: {index.Kind}</div>;
    }
}

export default {
    renderIndex
};
