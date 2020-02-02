import * as React from 'react';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import AddBoxIcon from '@material-ui/icons/AddBox';
import BookmarkBorderIcon from '@material-ui/icons/BookmarkBorder';
import BookmarksIcon from '@material-ui/icons/Bookmarks';
import ChatIcon from '@material-ui/icons/Chat';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CheckCircleOutlineIcon from '@material-ui/icons/CheckCircleOutline';
import CloudIcon from '@material-ui/icons/Cloud';
import CloudCircleIcon from '@material-ui/icons/CloudCircle';
import CloudQueueIcon from '@material-ui/icons/CloudQueue';
import DeleteIcon from '@material-ui/icons/Delete';
import DetailsIcon from '@material-ui/icons/Details';
import EditIcon from '@material-ui/icons/Edit';
import ErrorOutlineIcon from '@material-ui/icons/ErrorOutline';
import HelpOutlineIcon from '@material-ui/icons/HelpOutline';
import HighlightOffOutlinedIcon from '@material-ui/icons/HighlightOffOutlined';
import HomeIcon from '@material-ui/icons/Home';
import InfoIcon from '@material-ui/icons/Info';
import NotificationsOffIcon from '@material-ui/icons/NotificationsOff';
import ReceiptIcon from '@material-ui/icons/Receipt';
import SaveIcon from '@material-ui/icons/Save';
import SettingsIcon from '@material-ui/icons/Settings';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    error: {
      backgroundColor: theme.palette.error.dark,
    },
    marginDefault: {},
    marginLeft: {
      marginLeft: theme.spacing(1),
    },
    marginRight: {
      marginRight: theme.spacing(1),
    },
  });

interface IIcon extends WithStyles<typeof styles> {
  name;
  marginDirection?;
  onClick?;
  key?;
  style?;
}

class Icon extends React.Component<IIcon> {
  public render() {
    const {classes, name, marginDirection, ...props} = this.props;

    let className = '';
    switch (marginDirection) {
      case 'left':
        className = 'marginLeft';
        break;
      case 'right':
        className = 'marginRight';
        break;
      default:
        className = 'marginDefault';
        break;
    }

    switch (name) {
      case 'Check':
        return (
          <CheckCircleIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'Uncheck':
        return (
          <CheckCircleOutlineIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'Detail':
        return (
          <DetailsIcon key={name} className={classes[className]} {...props} />
        );
      case 'Check':
        return (
          <DetailsIcon key={name} className={classes[className]} {...props} />
        );
      case 'Update':
        return (
          <EditIcon key={name} className={classes[className]} {...props} />
        );
      case 'Save':
        return (
          <SaveIcon key={name} className={classes[className]} {...props} />
        );
      case 'Add':
        return (
          <AddBoxIcon key={name} className={classes[className]} {...props} />
        );
      case 'Create':
        return (
          <AddBoxIcon key={name} className={classes[className]} {...props} />
        );
      case 'Delete':
        return (
          <DeleteIcon key={name} className={classes[className]} {...props} />
        );
      case 'Success':
        return (
          <CheckCircleOutlineIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'Silenced':
        return (
          <NotificationsOffIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'Warning':
        return (
          <ErrorOutlineIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'Critical':
        return (
          <HighlightOffOutlinedIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'BookmarkBorder':
        return (
          <BookmarkBorderIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'Bookmarks':
        return (
          <BookmarksIcon key={name} className={classes[className]} {...props} />
        );
      case 'Info':
        return (
          <InfoIcon key={name} className={classes[className]} {...props} />
        );
      case 'Home':
        return (
          <HomeIcon key={name} className={classes[className]} {...props} />
        );
      case 'Settings':
        return (
          <SettingsIcon key={name} className={classes[className]} {...props} />
        );
      case 'HomeProject':
        return (
          <HomeIcon key={name} className={classes[className]} {...props} />
        );
      case 'Chat':
        return (
          <ChatIcon key={name} className={classes[className]} {...props} />
        );
      case 'Wiki':
        return (
          <ReceiptIcon key={name} className={classes[className]} {...props} />
        );
      case 'Ticket':
        return (
          <ReceiptIcon key={name} className={classes[className]} {...props} />
        );
      case 'ResourcePhysicalAdmin':
        return (
          <CloudIcon key={name} className={classes[className]} {...props} />
        );
      case 'ResourceVirtualAdmin':
        return (
          <CloudQueueIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'ResourcePhysical':
        return (
          <CloudIcon key={name} className={classes[className]} {...props} />
        );
      case 'ResourceVirtual':
        return (
          <CloudQueueIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      case 'ResourceMonitor':
        return (
          <CloudCircleIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
      default:
        return (
          <HelpOutlineIcon
            key={name}
            className={classes[className]}
            {...props}
          />
        );
    }
  }
}

export default withStyles(styles)(Icon);
