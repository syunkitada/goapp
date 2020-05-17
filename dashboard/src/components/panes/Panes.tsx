import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import Typography from "@material-ui/core/Typography";

import Index from "../Index";

import actions from "../../actions";
import logger from "../../lib/logger";

const styles = (theme: Theme): StyleRules =>
    createStyles({
        root: {
            backgroundColor: theme.palette.background.paper,
            flexGrow: 1,
            width: "100%"
        }
    });

interface IPanes extends WithStyles<typeof styles> {
    classes;
    index;
    indexPath;
    location;
    dispatchGetQueries;
}

class Panes extends React.Component<IPanes> {
    public render() {
        const { classes, indexPath, index } = this.props;
        if (!index.Children) {
            logger.error("Invalid index", index);
            return null;
        }

        let pane: any = null;
        for (let i = 0, len = index.Children.length; i < len; i++) {
            const child = index.Children[i];
            if (child.Name === indexPath) {
                pane = (
                    <Typography component="div">
                        <Index {...child} />
                    </Typography>
                );
                break;
            }
        }

        return <div className={classes.root}>{pane}</div>;
    }
}

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
)(withStyles(styles)(Panes));
