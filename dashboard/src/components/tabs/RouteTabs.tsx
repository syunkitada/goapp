import * as React from "react";
import { connect } from "react-redux";
import { Route } from "react-router-dom";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import Tabs from "./Tabs";

import logger from "../../lib/logger";

const styles = (theme: Theme): StyleRules =>
    createStyles({
        root: {
            width: "100%"
        }
    });

interface IRouteTabs extends WithStyles<typeof styles> {
    render;
    routes;
    index;
}

class RouteTabs extends React.Component<IRouteTabs> {
    public render() {
        const { classes, render, routes, index } = this.props;
        logger.info("RouteTabs", "render", routes);

        const beforeRoute = routes.slice(-1)[0];

        return (
            <div className={classes.root}>
                {index.Tabs.map(v => (
                    <Route
                        exact={v.Route === ""}
                        path={beforeRoute.match.path + v.Route}
                        key={v.Name}
                        render={props => {
                            const newRoutes = routes.slice(0);
                            newRoutes.push(props);
                            return (
                                <Tabs
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

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(withStyles(styles)(RouteTabs));
