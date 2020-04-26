import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import AppBar from "@material-ui/core/AppBar";
import Tab from "@material-ui/core/Tab";
import CoreTabs from "@material-ui/core/Tabs";
import Typography from "@material-ui/core/Typography";

import actions from "../../actions";
// import logger from "../../lib/logger";

const styles = (theme: Theme): StyleRules =>
    createStyles({
        root: {
            width: "100%",
            backgroundColor: theme.palette.background.paper,
            flexGrow: 1
        }
    });

interface IRouteTabs extends WithStyles<typeof styles> {
    dispatchGetQueries;
    render;
    location;
    index;
    indexPath: any;
}

class RouteTabs extends React.Component<IRouteTabs> {
    public render() {
        const { classes, render, location, index, indexPath } = this.props;
        console.log("DEBUG TODO RouteTabs render", location);

        const tabs: any[] = [];
        let tabContainer: any = null;
        let tabId = 0;
        for (let i = 0, len = index.Children.length; i < len; i++) {
            const tab = index.Children[i];
            if (tab.Name === indexPath) {
                tabId = i;
                tabContainer = (
                    <Typography component="div">{render(tab)}</Typography>
                );
            }
            tabs.push(<Tab key={tab.Name} label={tab.Name} />);
        }

        return (
            <div className={classes.root}>
                <AppBar position="static" color="default">
                    <CoreTabs
                        value={tabId}
                        onChange={this.handleChange}
                        indicatorColor="primary"
                        textColor="primary"
                        variant="scrollable"
                        scrollButtons="auto"
                    >
                        {tabs}
                    </CoreTabs>
                </AppBar>
                {tabContainer}
            </div>
        );
    }

    private handleChange = (event, tabId) => {
        const { location, index } = this.props;
        const child = index.Children[tabId];
        for (let i = 0, len = location.Path.length; i < len; i++) {
            if (index.Name === location.Path[i]) {
                location.Path[i + 1] = child.Name;
                break;
            }
        }
        this.props.dispatchGetQueries(index, location);
    };
}

function mapStateToProps(state, ownProps) {
    const { location } = state.service;

    const { index } = ownProps;
    let indexPath;
    if (index.Name === "Root") {
        indexPath = location.Path[0];
    } else {
        for (let i = 0, len = location.Path.length; i < len; i++) {
            if (index.Name === location.Path[i]) {
                indexPath = location.Path[i + 1];
                break;
            }
        }
    }

    return {
        indexPath,
        location
    };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {
        dispatchGetQueries: (index, location) => {
            dispatch(
                actions.service.serviceGetQueries({
                    index,
                    location,
                    searchQueries: null
                })
            );
        }
    };
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(withStyles(styles)(RouteTabs));
