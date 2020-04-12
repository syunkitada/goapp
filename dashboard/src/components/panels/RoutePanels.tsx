import * as React from "react";
import { connect } from "react-redux";
import { Route } from "react-router-dom";

import ExpansionPanels from "./ExpansionPanels";

import logger from "../../lib/logger";

interface IRoutePanels {
    render;
    routes;
    index;
}

class RoutePanels extends React.Component<IRoutePanels> {
    public render() {
        const { render, routes, index } = this.props;

        const beforeRoute = routes.slice(-1)[0];
        logger.info("RoutePanels", "render()", beforeRoute);

        return (
            <div>
                {index.Panels.map(v => (
                    <Route
                        exact={true}
                        path={beforeRoute.match.path + v.Route}
                        key={v.Name}
                        render={props => {
                            const newRoutes = routes.slice(0);
                            newRoutes.push(props);
                            return (
                                <ExpansionPanels
                                    render={render}
                                    routes={newRoutes}
                                    index={index}
                                    root={v}
                                />
                            );
                        }}
                    />
                ))}
            </div>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {};
}

function mapDispatchToProps(dispatch, ownProps) {
    return {};
}

export default connect(mapStateToProps, mapDispatchToProps)(RoutePanels);
