import * as React from "react";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import Paper from "@material-ui/core/Paper";
import Typography from "@material-ui/core/Typography";

const styles = (theme: Theme): StyleRules =>
    createStyles({
        root: {
            backgroundColor: theme.palette.background.paper,
            padding: theme.spacing(1)
        }
    });

interface IMsg extends WithStyles<typeof styles> {
    title;
    msg;
}

class Msg extends React.Component<IMsg> {
    public render() {
        const { classes, title, msg } = this.props;

        return (
            <Paper className={classes.root}>
                <Typography variant="h4">{title}</Typography>
                <Typography variant="body1">{msg}</Typography>
            </Paper>
        );
    }
}

export default withStyles(styles)(Msg);
