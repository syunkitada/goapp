import * as React from 'react';
import {connect} from 'react-redux';
import {NavLink} from 'react-router-dom';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import Icon from '../icons/Icon';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    nested: {
      paddingLeft: theme.spacing(1),
    },
  });

interface ILeftSidebar extends WithStyles<typeof styles> {
  classes;
  services;
  selectedServiceIndex;
}

class LeftSidebar extends React.Component<ILeftSidebar> {
  public render() {
    const {services, selectedServiceIndex} = this.props;

    const serviceHtmls: any[] = [];
    for (let i = 0, len = services.length; i < len; i++) {
      const service = services[i];
      const name = service.Name;
      let to = '/' + name;
      if (i === 0) {
        to = '/';
      }
      serviceHtmls.push(
        <NavLink
          key={name}
          to={to}
          style={{textDecoration: 'none', color: 'unset'}}>
          <ListItem
            button={true}
            dense={true}
            selected={i === selectedServiceIndex}>
            <ListItemIcon style={{minWidth: 30}}>
              <Icon name={service.Icon} />
            </ListItemIcon>
            <ListItemText primary={name} />
          </ListItem>
        </NavLink>,
      );
    }

    return (
      <div>
        <List dense={true}>{serviceHtmls}</List>
        <Divider />
      </div>
    );
  }
}

function mapStateToProps(state, ownProps) {
  return {};
}

export default connect(mapStateToProps)(withStyles(styles)(LeftSidebar));
