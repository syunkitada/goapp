import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import ExpansionPanel from "@material-ui/core/ExpansionPanel";
import ExpansionPanelDetails from "@material-ui/core/ExpansionPanelDetails";
import ExpansionPanelSummary from "@material-ui/core/ExpansionPanelSummary";
import Typography from "@material-ui/core/Typography";

import ExpandMoreIcon from "@material-ui/icons/ExpandMore";

import Index from "../Index";
import actions from "../../actions";
import data_utils from "../../lib/data_utils";
import logger from "../../lib/logger";

interface IPanels extends WithStyles<typeof styles> {
    location;
    locationParams;
    indexPath;
    index;
    dispatchGetQueries;
}

class Panels extends React.Component<IPanels> {
    public render() {
        const { indexPath, index, classes, locationParams } = this.props;
        logger.info("Panels.render", indexPath, index, locationParams);

        const panels: any[] = [];
        for (let i = 0, len = index.Children.length; i < len; i++) {
            const panel = index.Children[i];
            let subName = "";
            if (
                panel.SubNameParamKeys != null &&
                panel.SubNameParamKeys.length > 0
            ) {
                for (
                    let j = 0, lenj = panel.SubNameParamKeys.length;
                    j < lenj;
                    j++
                ) {
                    const paramKey = panel.SubNameParamKeys[j];
                    const paramData = locationParams[paramKey];
                    if (paramData) {
                        if (j === 0) {
                            subName += ":";
                        }
                        subName += " " + paramKey + "=" + paramData;
                    }
                }
            }
            panels.push(
                <ExpansionPanel
                    key={panel.Name}
                    expanded={panel.Name === indexPath}
                    onChange={() => this.handleChange(i)}
                >
                    <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
                        <Typography variant="subtitle1">
                            {panel.Name}
                            {subName}
                        </Typography>
                    </ExpansionPanelSummary>
                    <ExpansionPanelDetails style={{ padding: 0 }}>
                        <Index parent={index} {...panel} />
                    </ExpansionPanelDetails>
                </ExpansionPanel>
            );
        }

        return <div className={classes.root}>{panels}</div>;
    }

    private handleChange = panelId => {
        const { location, index } = this.props;
        const child = index.Children[panelId];

        let newLocation = null;
        for (const key in location.SubPathMap) {
            const splitedKey = key.split(".");
            for (let i = 0, len = splitedKey.length; i < len; i++) {
                if (splitedKey[i] === child.Name) {
                    newLocation = location.SubPathMap[key];
                }
            }
        }
        if (!newLocation) {
            newLocation = child.Name;
        }

        console.log(
            "DEBUG Panels handleChange",
            location,
            child.Name,
            newLocation
        );
        // location.SubPath = {};
        // location.SubPath[index.Name] = child.Name;
        this.props.dispatchGetQueries(index, newLocation);
    };
}

const styles = (theme: Theme): StyleRules =>
    createStyles({
        root: {
            backgroundColor: theme.palette.background.paper,
            flexGrow: 1,
            width: "100%"
        }
    });

function mapStateToProps(state, ownProps) {
    const { location } = state.service;
    const { index } = ownProps;

    let indexPath;
    if (location.SubPath) {
        indexPath = location.SubPath[index.Name];
    } else if (index.Name === "Root") {
        indexPath = location.Path[0];
    } else {
        for (let i = 0, len = location.Path.length; i < len; i++) {
            if (index.Name === location.Path[i]) {
                indexPath = location.Path[i + 1];
                break;
            }
        }
    }

    const locationData = data_utils.getLocationData();
    let locationParams = {};
    if (locationData.Params) {
        locationParams = locationData.Params;
    }

    return {
        indexPath,
        location,
        locationParams
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
)(withStyles(styles)(Panels));
