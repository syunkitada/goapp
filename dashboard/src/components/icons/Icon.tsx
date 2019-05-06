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

const styles = (theme: Theme): StyleRules =>
  createStyles({
    error: {
      backgroundColor: theme.palette.error.dark,
    },
    marginLeft: {
      marginLeft: theme.spacing.unit,
    },
    marginRight: {
      marginRight: theme.spacing.unit,
    },
  });

interface IIcon extends WithStyles<typeof styles> {
  kind;
  marginDirection;
}

class Icon extends React.Component<IIcon> {
  public render() {
    const {classes, kind, marginDirection} = this.props;

    let className = '';
    switch (marginDirection) {
      case 'left':
        className = 'marginLeft';
        break;
      case 'right':
        className = 'marginRight';
        break;
    }

    switch (kind) {
      case 'Detail':
        return <DetailsIcon className={classes[className]} />;
      case 'Update':
        return <EditIcon className={classes[className]} />;
      case 'Create':
        return <AddBoxIcon className={classes[className]} />;
      case 'Delete':
        return <DeleteIcon className={classes[className]} />;
      default:
        return <span>IconNotFound</span>;
    }
  }
}

export default withStyles(styles)(Icon);
