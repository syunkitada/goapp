import * as React from 'react';

import classNames from 'classnames';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CloseIcon from '@material-ui/icons/Close';
import ErrorIcon from '@material-ui/icons/Error';
import InfoIcon from '@material-ui/icons/Info';
import WarningIcon from '@material-ui/icons/Warning';

import amber from '@material-ui/core/colors/amber';
import green from '@material-ui/core/colors/green';

import IconButton from '@material-ui/core/IconButton';
import Snackbar from '@material-ui/core/Snackbar';
import SnackbarContent from '@material-ui/core/SnackbarContent';

const variantIcon = {
  error: ErrorIcon,
  info: InfoIcon,
  success: CheckCircleIcon,
  warning: WarningIcon,
};

const styles = (theme: Theme): StyleRules =>
  createStyles({
    error: {
      backgroundColor: theme.palette.error.dark,
    },
    icon: {
      fontSize: 20,
    },
    iconVariant: {
      marginRight: theme.spacing.unit,
      opacity: 0.9,
    },
    info: {
      backgroundColor: theme.palette.primary.dark,
    },
    message: {
      alignItems: 'center',
      display: 'flex',
    },
    success: {
      backgroundColor: green[600],
    },
    warning: {
      backgroundColor: amber[700],
    },
  });

interface IMsgSnackbar extends WithStyles<typeof styles> {
  open;
  onClose;
  variant;
  vertical;
  horizontal;
  msg;
}

class MsgSnackbar extends React.Component<IMsgSnackbar> {
  public render() {
    const {
      classes,
      open,
      onClose,
      variant,
      vertical,
      horizontal,
      msg,
    } = this.props;

    const Icon = variantIcon[variant];

    return (
      <Snackbar
        anchorOrigin={{vertical, horizontal}}
        open={open}
        onClose={onClose}
        ContentProps={{
          'aria-describedby': 'message-id',
        }}>
        <SnackbarContent
          className={classes[variant]}
          aria-describedby="client-snackbar"
          message={
            <span id="client-snackbar" className={classes.message}>
              <Icon className={classNames(classes.icon, classes.iconVariant)} />
              {msg}
            </span>
          }
          action={[
            <IconButton
              key="close"
              aria-label="Close"
              color="inherit"
              className={classes.close}
              onClick={onClose}>
              <CloseIcon className={classes.icon} />
            </IconButton>,
          ]}
        />
      </Snackbar>
    );
  }
}

export default withStyles(styles)(MsgSnackbar);
