import * as React from 'react';
import {connect} from 'react-redux';
import {NavLink} from 'react-router-dom';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import Collapse from '@material-ui/core/Collapse';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import AssessmentIcon from '@material-ui/icons/Assessment';
import ChatIcon from '@material-ui/icons/Chat';
import CloudIcon from '@material-ui/icons/Cloud';
import CloudQueueIcon from '@material-ui/icons/CloudQueue';
import DashboardIcon from '@material-ui/icons/Dashboard';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import HomeIcon from '@material-ui/icons/Home';
import LayersIcon from '@material-ui/icons/Layers';
import NoteAddIcon from '@material-ui/icons/NoteAdd';
import ReceiptIcon from '@material-ui/icons/Receipt';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    nested: {
      paddingLeft: theme.spacing(4),
    },
  });

interface ILeftSidebar extends WithStyles<typeof styles> {
  history;
  classes;
  auth;
  match;
}

class LeftSidebar extends React.Component<ILeftSidebar> {
  public state = {
    openProjects: false,
  };

  public render() {
    const {classes, auth, match} = this.props;

    if (!auth.user) {
      return null;
    }

    // https://material.io/tools/icons/?style=baseline
    const serviceLinks: any[] = [
      ['Chat', <ChatIcon key={'Chat'} />],
      ['Wiki', <ReceiptIcon key={'Wiki'} />],
      ['Ticket', <NoteAddIcon key={'Ticket'} />],
      ['Datacenter', <LayersIcon key={'Datacenter'} />],
      ['Home.Project', <HomeIcon key={'Home.Project'} />],
      ['Resource.Physical', <CloudIcon key={'Resource.Physical'} />],
      ['Resource.Virtual', <CloudQueueIcon key={'Resource.Virtual'} />],
      ['Monitor', <AssessmentIcon key={'Monitor'} />],
    ];

    const services: any[] = [];
    let serviceMap: any = null;
    let projectText: any = null;
    let prefixPath: any = null;
    if (match.params.project) {
      prefixPath = '/Project/' + match.params.project + '/';
      projectText = match.params.project;
      serviceMap =
        auth.user.authority.ProjectServiceMap[match.params.project].ServiceMap;
    } else {
      prefixPath = '/Service/';
      projectText = 'Projects';
      serviceMap = auth.user.authority.ServiceMap;
    }

    for (const serviceLink of serviceLinks) {
      if (serviceLink[0] in serviceMap) {
        const path = prefixPath + serviceLink[0];
        services.push(
          <NavLink
            key={serviceLink[0]}
            to={path}
            style={{textDecoration: 'none', color: 'unset'}}>
            <ListItem button={true} selected={match.url === path}>
              <ListItemIcon>{serviceLink[1]}</ListItemIcon>
              <ListItemText primary={serviceLink[0]} />
            </ListItem>
          </NavLink>,
        );
      }
    }

    const projects: any[] = [];
    for (const project of Object.keys(auth.user.authority.ProjectServiceMap)) {
      const path = '/Project/' + project + '/Home.Project';
      projects.push(
        <List key={project} disablePadding={true}>
          <ListItem
            button={true}
            className={classes.nested}
            onClick={event => this.handleProjectClick(event, path)}>
            <ListItemIcon>
              <DashboardIcon />
            </ListItemIcon>
            <ListItemText inset={true} primary={project} />
          </ListItem>
        </List>,
      );
    }

    return (
      <div>
        <Divider />
        <List>
          <NavLink
            to="/Service/Home"
            style={{textDecoration: 'none', color: 'unset'}}>
            <ListItem button={true} selected={match.url === '/Service/Home'}>
              <ListItemIcon>
                <HomeIcon />
              </ListItemIcon>
              <ListItemText primary="Home" />
            </ListItem>
          </NavLink>

          <ListItem button={true} onClick={this.handleOpenProjectsClick}>
            <ListItemIcon>
              <DashboardIcon />
            </ListItemIcon>
            <ListItemText inset={true} primary={projectText} />
            {this.state.openProjects ? <ExpandLess /> : <ExpandMore />}
          </ListItem>
          <Collapse
            in={this.state.openProjects}
            timeout="auto"
            unmountOnExit={true}>
            {projects}
          </Collapse>
        </List>
        <Divider />
        <List>{services}</List>
        <Divider />
      </div>
    );
  }

  private handleOpenProjectsClick = () => {
    this.setState(state => ({openProjects: !this.state.openProjects}));
  };

  private handleProjectClick = (event, path) => {
    const {history} = this.props;
    history.push(path);
    this.setState({openProjects: false});
  };
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth;

  return {auth};
}

export default connect(mapStateToProps)(withStyles(styles)(LeftSidebar));
