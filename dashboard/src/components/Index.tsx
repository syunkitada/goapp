import * as React from "react";

import IndexForm from "./forms/IndexForm";
import SearchForm from "./forms/SearchForm";
import RoutePanels from "./panels/RoutePanels";
import Panes from "./panes/Panes";
import IndexTable from "./tables/IndexTable";
import RouteTabs from "./tabs/RouteTabs";
import IndexView from "./views/IndexView";

import logger from "../lib/logger";

function Index(props) {
    logger.info("Index.render", props.Kind, props);
    switch (props.Kind) {
        case "Msg":
            return <div>Msg</div>;
        case "RoutePanels":
            return <RoutePanels index={props} />;
        case "RouteTabs":
            return <RouteTabs index={props} />;
        case "RoutePanes":
            return <Panes index={props} />;
        case "Table":
            return <IndexTable index={props} />;
        case "View":
            return <IndexView index={props} />;
        case "SearchForm":
            return <SearchForm index={props} />;
        case "Form":
            return <IndexForm index={props} />;
        default:
            return <div>Unsupported Kind: {props.Kind}</div>;
    }
}

export default Index;
