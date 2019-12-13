import * as React from 'react';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import AddBoxIcon from '@material-ui/icons/AddBox';
import DeleteIcon from '@material-ui/icons/Delete';
import DetailsIcon from '@material-ui/icons/Details';
import EditIcon from '@material-ui/icons/Edit';
import ErrorOutlineIcon from '@material-ui/icons/ErrorOutline';
import HelpOutlineIcon from '@material-ui/icons/HelpOutline';
import HighlightOffOutlinedIcon from '@material-ui/icons/HighlightOffOutlined';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    error: {
      backgroundColor: theme.palette.error.dark,
    },
    marginLeft: {
      marginLeft: theme.spacing(1),
    },
    marginRight: {
      marginRight: theme.spacing(1),
    },
  });

interface IIcon extends WithStyles<typeof styles> {
  name;
  marginDirection;
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
    }

    switch (name) {
      case 'Detail':
        return <DetailsIcon className={classes[className]} {...props} />;
      case 'Update':
        return <EditIcon className={classes[className]} {...props} />;
      case 'Create':
        return <AddBoxIcon className={classes[className]} {...props} />;
      case 'Delete':
        return <DeleteIcon className={classes[className]} {...props} />;
      case 'Warning':
        return <ErrorOutlineIcon className={classes[className]} {...props} />;
      case 'Critical':
        return (
          <HighlightOffOutlinedIcon className={classes[className]} {...props} />
        );
      default:
        return <HelpOutlineIcon className={classes[className]} {...props} />;
    }
  }
}

export default withStyles(styles)(Icon);
