import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
    StyleRules,
    WithStyles
} from "@material-ui/core/styles/withStyles";

import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import FormControl from "@material-ui/core/FormControl";
import Input from "@material-ui/core/Input";
import InputLabel from "@material-ui/core/InputLabel";
import Paper from "@material-ui/core/Paper";
import Typography from "@material-ui/core/Typography";

import LockOutlinedIcon from "@material-ui/icons/LockOutlined";

import MsgSnackbar from "../../../components/snackbars/MsgSnackbar";

import actions from "../../../actions";
import logger from "../../../lib/logger";

const styles = (theme: Theme): StyleRules =>
    createStyles({
        avatar: {
            backgroundColor: theme.palette.secondary.main,
            margin: theme.spacing(1)
        },
        form: {
            marginTop: theme.spacing(1),
            width: "100%" // Fix IE11 issue.
        },
        layout: {
            display: "block", // Fix IE11 issue.
            marginLeft: theme.spacing(3),
            marginRight: theme.spacing(3),
            width: "auto",
            [theme.breakpoints.up(400 + theme.spacing(6))]: {
                marginLeft: "auto",
                marginRight: "auto",
                width: 400
            }
        },
        paper: {
            alignItems: "center",
            display: "flex",
            flexDirection: "column",
            marginTop: theme.spacing(8),
            padding: `${theme.spacing(2)}px ${theme.spacing(
                3
            )}px ${theme.spacing(3)}px`
        },
        submit: {
            marginTop: theme.spacing(3)
        }
    });

interface ILogin extends WithStyles<typeof styles> {
    auth;
    onSubmit;
}

class Login extends React.Component<ILogin> {
    public render() {
        const { classes, auth, onSubmit } = this.props;
        logger.info("Login", "render");

        if (auth.isFetching) {
            return <div>During authentication</div>;
        }

        let msgHtml: any = null;
        if (auth.error != null && auth.error !== "") {
            const variant = "error";
            const vertical = "bottom";
            const horizontal = "left";

            msgHtml = (
                <MsgSnackbar
                    open={true}
                    onClose={this.handleClose}
                    vertical={vertical}
                    horizontal={horizontal}
                    variant={variant}
                    msg={auth.error}
                />
            );
        }

        return (
            <React.Fragment>
                <CssBaseline />
                {msgHtml}
                <main className={classes.layout}>
                    <Paper className={classes.paper}>
                        <Avatar className={classes.avatar}>
                            <LockOutlinedIcon />
                        </Avatar>
                        <Typography variant="h4">Sign in</Typography>
                        <form className={classes.form} onSubmit={onSubmit}>
                            <FormControl
                                margin="normal"
                                required={true}
                                fullWidth={true}
                            >
                                <InputLabel htmlFor="username">Name</InputLabel>
                                <Input
                                    id="username"
                                    name="username"
                                    autoFocus={true}
                                />
                            </FormControl>
                            <FormControl
                                margin="normal"
                                required={true}
                                fullWidth={true}
                            >
                                <InputLabel htmlFor="password">
                                    Password
                                </InputLabel>
                                <Input
                                    name="password"
                                    type="password"
                                    id="password"
                                    autoComplete="current-password"
                                />
                            </FormControl>
                            <Button
                                type="submit"
                                fullWidth={true}
                                variant="contained"
                                color="primary"
                                className={classes.submit}
                            >
                                Sign in
                            </Button>
                        </form>
                    </Paper>
                </main>
            </React.Fragment>
        );
    }

    private handleClose = (event, reason) => {
        return;
    };
}

function mapStateToProps(state, ownProps) {
    const auth = state.auth;

    return { auth };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {
        onSubmit: e => {
            e.preventDefault();
            const username = e.target.username.value.trim();
            const password = e.target.password.value.trim();
            dispatch(actions.auth.authLogin({ username, password }));
        }
    };
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(withStyles(styles)(Login));
