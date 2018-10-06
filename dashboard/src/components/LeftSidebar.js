import { connect } from 'react-redux';
import React, {Component} from 'react';
import actions from '../actions'
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListSubheader from '@material-ui/core/ListSubheader';
import DashboardIcon from '@material-ui/icons/Dashboard';
import ChatIcon from '@material-ui/icons/Chat';
import ReceiptIcon from '@material-ui/icons/Receipt';
import HomeIcon from '@material-ui/icons/Home';
import NoteAddIcon from '@material-ui/icons/NoteAdd';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import PeopleIcon from '@material-ui/icons/People';
import BarChartIcon from '@material-ui/icons/BarChart';
import LayersIcon from '@material-ui/icons/Layers';
import AssignmentIcon from '@material-ui/icons/Assignment';
import ViewComfyIcon from '@material-ui/icons/ViewComfy';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import Typography from '@material-ui/core/Typography';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import Collapse from '@material-ui/core/Collapse';
import InboxIcon from '@material-ui/icons/MoveToInbox';
import DraftsIcon from '@material-ui/icons/Drafts';
import SendIcon from '@material-ui/icons/Send';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import StarBorder from '@material-ui/icons/StarBorder';
import { NavLink } from 'react-router-dom';

const styles = theme => ({
  nested: {
    paddingLeft: theme.spacing.unit * 4,
  },
});

class LeftSidebar extends Component {
  state = {
    open: false,
  };

  handleClick = () => {
    this.setState(state => ({open: !state.open}));
  };

  render() {
    const { classes, auth, projectService, match } = this.props;

    if (!auth.user) {
      return null
    }

    var services = [];
    var serviceMap = null
    var projectText = null
    var prefixPath = null
    if (projectService) {
      prefixPath = '/Project/' + projectService.ProjectName + '/'
      projectText = projectService.ProjectName
      serviceMap = projectService.ServiceMap
      serviceMap['Home'] = {}
    } else {
      prefixPath = '/'
      projectText = 'Projects'
      serviceMap = auth.user.Authority.ServiceMap
    }

    // https://material.io/tools/icons/?style=baseline
    var serviceLinks = [
      ["Chat", <ChatIcon />],
      ["Wiki", <ReceiptIcon />],
      ["Ticket", <NoteAddIcon />],
      ["Datacenter", <LayersIcon />],
      ["Home", <DashboardIcon />],
      ["Resource", <ViewComfyIcon />],
    ]

    for (let serviceLink of serviceLinks) {
      if (serviceLink[0] in serviceMap) {
        let path = prefixPath + serviceLink[0]
        services.push(
          <NavLink key={serviceLink[0]} to={path} style={{textDecoration: 'none', color: 'unset'}}>
            <ListItem button selected={match.url === path}>
              <ListItemIcon>
                {serviceLink[1]}
              </ListItemIcon>
              <ListItemText primary={serviceLink[0]} />
            </ListItem>
          </NavLink>
        )
      }
    }

    var projects = [];
    for (let project in auth.user.Authority.ProjectServiceMap) {
      let path = "/Project/" + project + "/Home"
      projects.push(
        <NavLink key={project} to={path} style={{textDecoration: 'none', color: 'unset'}}>
          <List component="div" disablePadding>
            <ListItem button className={classes.nested}>
              <ListItemIcon>
                <DashboardIcon />
              </ListItemIcon>
              <ListItemText inset primary={project} />
            </ListItem>
          </List>
        </NavLink>
      )
    }

    return (
      <div>
        <Divider />
        <List>
          <NavLink to="/Home" style={{textDecoration: 'none', color: 'unset'}}>
            <ListItem button selected={match.url === '/Home'}>
              <ListItemIcon>
                <HomeIcon />
              </ListItemIcon>
              <ListItemText primary="Home" />
            </ListItem>
          </NavLink>

          <ListItem button onClick={this.handleClick}>
            <ListItemIcon>
              <DashboardIcon />
            </ListItemIcon>
            <ListItemText inset primary={projectText} />
            {this.state.open ? <ExpandLess /> : <ExpandMore />}
          </ListItem>
          <Collapse in={this.state.open} timeout="auto" unmountOnExit>
            {projects}
          </Collapse>
        </List>
        <Divider />
        <List>
          {services}
        </List>
        <Divider />
      </div>
    )
  }
}

LeftSidebar.propTypes = {
  classes: PropTypes.object.isRequired,
  auth: PropTypes.object.isRequired,
};

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

export default connect(
  mapStateToProps,
)(withStyles(styles)(LeftSidebar))
