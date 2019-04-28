import { connect } from 'react-redux';
import React, {Component} from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import DashboardIcon from '@material-ui/icons/Dashboard';
import ChatIcon from '@material-ui/icons/Chat';
import ReceiptIcon from '@material-ui/icons/Receipt';
import HomeIcon from '@material-ui/icons/Home';
import NoteAddIcon from '@material-ui/icons/NoteAdd';
import LayersIcon from '@material-ui/icons/Layers';
import CloudQueueIcon from '@material-ui/icons/CloudQueue';
import CloudIcon from '@material-ui/icons/Cloud';
import AssessmentIcon from '@material-ui/icons/Assessment';
import Collapse from '@material-ui/core/Collapse';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import { NavLink } from 'react-router-dom';

const styles = theme => ({
  nested: {
    paddingLeft: theme.spacing.unit * 4,
  },
});

class LeftSidebar extends Component {
  state = {
    openProjects: false,
  };

  handleOpenProjectsClick = () => {
    this.setState(state => ({openProjects: !state.openProjects}));
  };

  handleProjectClick = (event, path) => {
    const { history } = this.props;
    history.push(path)
    this.setState({ openProjects: false });
  };

  render() {
    const { classes, auth, match } = this.props;

    if (!auth.user) {
      return null
    }

    // https://material.io/tools/icons/?style=baseline
    var serviceLinks = [
      ["Chat", <ChatIcon />],
      ["Wiki", <ReceiptIcon />],
      ["Ticket", <NoteAddIcon />],
      ["Datacenter", <LayersIcon />],
      ["Home.Project", <HomeIcon />],
      ["Resource.Physical", <CloudIcon />],
      ["Resource.Virtual", <CloudQueueIcon />],
      ["Monitor", <AssessmentIcon />],
    ]

    var services = [];
    var serviceMap = null
    var projectText = null
    var prefixPath = null
    if (match.params.project) {
      prefixPath = '/Project/' + match.params.project + '/'
      projectText = match.params.project
      serviceMap = auth.user.Authority.ProjectServiceMap[match.params.project].ServiceMap
    } else {
      prefixPath = '/Service/'
      projectText = 'Projects'
      serviceMap = auth.user.Authority.ServiceMap
    }

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
      let path = "/Project/" + project + "/Home.Project"
      projects.push(
        <List key={project} component="div" disablePadding>
          <ListItem button className={classes.nested} onClick={event => this.handleProjectClick(event, path)}>
            <ListItemIcon>
              <DashboardIcon />
            </ListItemIcon>
            <ListItemText inset primary={project} />
          </ListItem>
        </List>
      )
    }

    return (
      <div>
        <Divider />
        <List>
          <NavLink to="/Service/Home" style={{textDecoration: 'none', color: 'unset'}}>
            <ListItem button selected={match.url === '/Service/Home'}>
              <ListItemIcon>
                <HomeIcon />
              </ListItemIcon>
              <ListItemText primary="Home" />
            </ListItem>
          </NavLink>

          <ListItem button onClick={this.handleOpenProjectsClick}>
            <ListItemIcon>
              <DashboardIcon />
            </ListItemIcon>
            <ListItemText inset primary={projectText} />
            {this.state.openProjects ? <ExpandLess /> : <ExpandMore />}
          </ListItem>
          <Collapse in={this.state.openProjects} timeout="auto" unmountOnExit>
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
  match: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired,
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
