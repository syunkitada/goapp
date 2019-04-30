import * as React from 'react';
import {Redirect} from 'react-router-dom';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import withStyles, {
  WithStyles,
  StyleRules,
} from '@material-ui/core/styles/withStyles';
import createStyles from '@material-ui/core/styles/createStyles';

import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import FormControl from '@material-ui/core/FormControl';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import LockIcon from '@material-ui/icons/LockOutlined';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';

import actions from '../../../actions';
import MsgSnackbar from '../../../components/snackbars/MsgSnackbar';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    layout: {
      width: 'auto',
      display: 'block', // Fix IE11 issue.
      marginLeft: theme.spacing.unit * 3,
      marginRight: theme.spacing.unit * 3,
      [theme.breakpoints.up(400 + theme.spacing.unit * 3 * 2)]: {
        width: 400,
        marginLeft: 'auto',
        marginRight: 'auto',
      },
    },
    paper: {
      marginTop: theme.spacing.unit * 8,
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      padding: `${theme.spacing.unit * 2}px ${theme.spacing.unit * 3}px ${theme
        .spacing.unit * 3}px`,
    },
    avatar: {
      margin: theme.spacing.unit,
      backgroundColor: theme.palette.secondary.main,
    },
    form: {
      width: '100%', // Fix IE11 issue.
      marginTop: theme.spacing.unit,
    },
    submit: {
      marginTop: theme.spacing.unit * 3,
    },
  });

interface ILogin extends WithStyles<typeof styles> {
  auth;
  history;
  onSubmit;
}

class Login extends React.Component<ILogin> {
  handleClose = (event, reason) => {
    console.log('Close');
  };

  render() {
    const {classes, auth, history, onSubmit} = this.props;
    const {from} = history.location.state || {from: {pathname: '/'}};

    if (auth.redirectToReferrer) {
      return <Redirect to={from} />;
    }

    if (auth.user) {
      return <Redirect to={{pathname: '/'}} />;
    }

    if (auth.isFetching) {
      return <div>During authentication</div>;
    }

    let msgHtml: any = null;
    if (auth.error != null && auth.error !== '') {
      let variant = 'error';
      let vertical = 'bottom';
      let horizontal = 'left';

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
              <LockIcon />
            </Avatar>
            <Typography variant="headline">Sign in</Typography>
            <form className={classes.form} onSubmit={onSubmit}>
              <FormControl margin="normal" required fullWidth>
                <InputLabel htmlFor="name">Name</InputLabel>
                <Input id="name" name="name" autoFocus />
              </FormControl>
              <FormControl margin="normal" required fullWidth>
                <InputLabel htmlFor="password">Password</InputLabel>
                <Input
                  name="password"
                  type="password"
                  id="password"
                  autoComplete="current-password"
                />
              </FormControl>
              <Button
                type="submit"
                fullWidth
                variant="raised"
                color="primary"
                className={classes.submit}>
                Sign in
              </Button>
            </form>
          </Paper>
        </main>
      </React.Fragment>
    );
  }
}

// Login.propTypes = {
//   auth: PropTypes.object.isRequired,
//   onSubmit: PropTypes.func.isRequired,
//   history: PropTypes.object.isRequired,
// };

function mapStateToProps(state, ownProps) {
  const auth = state.auth;

  return {
    auth: auth,
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    onSubmit: e => {
      e.preventDefault();
      const name = e.target.name.value.trim();
      const password = e.target.password.value.trim();
      dispatch(actions.auth.authLogin(name, password));
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(Login));
