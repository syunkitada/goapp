import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

// xterm: https://github.com/xtermjs/xterm.js
// xterm addons: https://github.com/xtermjs/xterm.js/tree/4.4.0/addons
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import "xterm/css/xterm.css";

interface IConsoleCore extends WithStyles<typeof styles> {
    webSocket;
}

class ConsoleCore extends React.Component<IConsoleCore> {
    xterm: Terminal;
    fitAddon: FitAddon;
    container: any;

    constructor(props, context) {
        super(props, context);
        this.xterm = new Terminal(props.options);
        this.fitAddon = new FitAddon();
        this.xterm.loadAddon(this.fitAddon);
        this.state = {
            isFocused: false
        };
    }

    componentDidMount() {
        const { webSocket } = this.props;

        this.xterm.open(this.container);
        this.fitAddon.fit();
        webSocket.onmessage = event => {
            const data = JSON.parse(event.data);
            this.xterm.write(atob(data.Bytes));
        };

        this.xterm.onData(data => {
            console.log("DATA", data);
            const body = JSON.stringify({
                Bytes: btoa(data)
            });
            webSocket.send(body);
        });
    }

    shouldComponentUpdate(nextProps, nextState) {
        // TODO
        return false;
    }

    public render() {
        return <div ref={ref => (this.container = ref)} />;
    }
}

function mapStateToProps(state, ownProps) {
    return {};
}

function mapDispatchToProps(dispatch, ownProps) {
    return {};
}

const styles = (theme: Theme): StyleRules =>
    createStyles({
        root: {
            width: "100%"
        }
    });

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(withStyles(styles, { withTheme: true })(ConsoleCore));
