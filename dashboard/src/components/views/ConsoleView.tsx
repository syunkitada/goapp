import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import ConsoleCore from "./ConsoleCore";

import data_utils from "../../lib/data_utils";

interface IConsoleView extends WithStyles<typeof styles> {
    webSocket;
}

class ConsoleView extends React.Component<IConsoleView> {
    public render() {
        const { classes, webSocket } = this.props;
        if (!webSocket) {
            return <div>Loading</div>;
        }
        return (
            <div className={classes.root}>
                <ConsoleCore webSocket={webSocket} />
            </div>
        );
    }
}

function mapStateToProps(state, ownProps) {
    const { index } = ownProps;
    const { Data, WebSocketMap } = data_utils.getServiceState(state);
    let webSocket: any;
    if (index.EnableWebSocket) {
        if (WebSocketMap) {
            webSocket = WebSocketMap[index.WebSocketKey];
        }
    }

    return { data: Data, webSocket };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {};
}

const styles = (theme: Theme): StyleRules =>
    createStyles({
        root: {
            margin: theme.spacing(1)
        }
    });

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(withStyles(styles, { withTheme: true })(ConsoleView));
