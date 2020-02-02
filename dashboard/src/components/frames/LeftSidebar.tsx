import * as React from "react";
import { connect } from "react-redux";
import { NavLink } from "react-router-dom";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
  StyleRules,
  WithStyles
} from "@material-ui/core/styles/withStyles";

import Collapse from "@material-ui/core/Collapse";
import Divider from "@material-ui/core/Divider";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";

import DashboardIcon from "@material-ui/icons/Dashboard";
import ExpandLess from "@material-ui/icons/ExpandLess";
import ExpandMore from "@material-ui/icons/ExpandMore";
import HomeIcon from "@material-ui/icons/Home";

import Icon from "../icons/Icon";

const styles = (theme: Theme): StyleRules =>
  createStyles({
    nested: {
      paddingLeft: theme.spacing(1)
    }
  });

interface ILeftSidebar extends WithStyles<typeof styles> {
  classes;
  services;
  selectedServiceIndex;
  auth;
  history;
  match;
}

class LeftSidebar extends React.Component<ILeftSidebar> {
  public state = {
    openProjects: false
  };

  public render() {
    const { services, selectedServiceIndex, auth, match, classes } = this.props;

    if (services) {
      const serviceHtmls: any[] = [];
      for (let i = 0, len = services.length; i < len; i++) {
        const service = services[i];
        const name = service.Name;
        let to = "/" + name;
        if (i === 0) {
          to = "/";
        }
        serviceHtmls.push(
          <NavLink
            key={name}
            to={to}
            style={{ textDecoration: "none", color: "unset" }}
          >
            <ListItem
              button={true}
              dense={true}
              selected={i === selectedServiceIndex}
            >
              <ListItemIcon style={{ minWidth: 30 }}>
                <Icon name={service.Icon} />
              </ListItemIcon>
              <ListItemText primary={name} />
            </ListItem>
          </NavLink>
        );
      }

      return (
        <div>
          <List dense={true}>{serviceHtmls}</List>
          <Divider />
        </div>
      );
    } else {
      if (!auth.user) {
        return null;
      }

      const serviceHtmls: any[] = [];
      let serviceMap: any = null;
      let projectText: any = null;
      let prefixPath: any = null;
      if (match.params.project) {
        prefixPath = "/Project/" + match.params.project + "/";
        projectText = match.params.project;
        serviceMap =
          auth.user.authority.ProjectServiceMap[match.params.project]
            .ServiceMap;
      } else {
        prefixPath = "/Service/";
        projectText = "Projects";
        serviceMap = auth.user.authority.ServiceMap;
      }

      const tmpServices = Object.keys(serviceMap);
      tmpServices.sort();
      for (const serviceName of tmpServices) {
        const path = prefixPath + serviceName;
        serviceHtmls.push(
          <NavLink
            key={serviceName}
            to={path}
            style={{ textDecoration: "none", color: "unset" }}
          >
            <ListItem button={true} dense={true} selected={match.url === path}>
              <ListItemIcon style={{ minWidth: 30 }}>
                <Icon name={serviceName} />
              </ListItemIcon>
              <ListItemText primary={serviceName} />
            </ListItem>
          </NavLink>
        );
      }

      const projects: any[] = [];
      const tmpProjects = Object.keys(auth.user.authority.ProjectServiceMap);
      tmpProjects.sort();
      for (const project of tmpProjects) {
        const path = "/Project/" + project + "/HomeProject";
        projects.push(
          <List key={project} disablePadding={true} dense={true}>
            <ListItem
              button={true}
              dense={true}
              className={classes.nested}
              onClick={event => this.handleProjectClick(event, path)}
            >
              <ListItemIcon style={{ minWidth: 30 }}>
                <DashboardIcon />
              </ListItemIcon>
              <ListItemText inset={true} primary={project} />
            </ListItem>
          </List>
        );
      }

      return (
        <div>
          <Divider />
          <List dense={true}>
            <NavLink
              to="/Service/Home"
              style={{ textDecoration: "none", color: "unset" }}
            >
              <ListItem
                button={true}
                dense={true}
                selected={match.url === "/Service/Home"}
              >
                <ListItemIcon style={{ minWidth: 30 }}>
                  <HomeIcon />
                </ListItemIcon>
                <ListItemText primary="Home" />
              </ListItem>
            </NavLink>

            <ListItem
              button={true}
              dense={true}
              onClick={this.handleOpenProjectsClick}
            >
              <ListItemIcon style={{ minWidth: 30 }}>
                <DashboardIcon />
              </ListItemIcon>
              <ListItemText inset={true} primary={projectText} />
              {this.state.openProjects ? <ExpandLess /> : <ExpandMore />}
            </ListItem>
            <Collapse
              in={this.state.openProjects}
              timeout="auto"
              unmountOnExit={true}
            >
              {projects}
            </Collapse>
          </List>
          <Divider />
          <List dense={true}>{serviceHtmls}</List>
          <Divider />
        </div>
      );
    }
  }

  private handleOpenProjectsClick = () => {
    this.setState(state => ({ openProjects: !this.state.openProjects }));
  };

  private handleProjectClick = (event, path) => {
    const { history } = this.props;
    history.push(path);
    this.setState({ openProjects: false });
  };
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth;

  return { auth };
}

export default connect(mapStateToProps)(withStyles(styles)(LeftSidebar));
