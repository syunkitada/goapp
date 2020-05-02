import * as React from "react";
import { connect } from "react-redux";

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

import actions from "../../actions";

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
    projectName;
    serviceName;
    dispatchServiceGetIndex;
}

class LeftSidebar extends React.Component<ILeftSidebar> {
    public state = {
        openProjects: false
    };

    public render() {
        const {
            services,
            selectedServiceIndex,
            auth,
            projectName,
            serviceName,
            classes
        } = this.props;

        if (services) {
            const serviceHtmls: any[] = [];
            for (let i = 0, len = services.length; i < len; i++) {
                const service = services[i];
                const name = service.Name;
                serviceHtmls.push(
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
                );
            }

            return (
                <div>
                    <List dense={true}>{serviceHtmls}</List>
                    <Divider />
                </div>
            );
        } else {
            const serviceHtmls: any[] = [];
            let serviceMap: any = null;
            let projectText: any = null;
            if (projectName) {
                projectText = projectName;
                serviceMap =
                    auth.user.authority.ProjectServiceMap[projectName]
                        .ServiceMap;
            } else {
                projectText = "Projects";
                serviceMap = auth.user.authority.ServiceMap;
            }

            const tmpServices = Object.keys(serviceMap);
            tmpServices.sort();
            for (const tmpServiceName of tmpServices) {
                serviceHtmls.push(
                    <ListItem
                        key={tmpServiceName}
                        button={true}
                        dense={true}
                        selected={tmpServiceName === serviceName}
                        onClick={this.handleClickService.bind(
                            this,
                            tmpServiceName
                        )}
                    >
                        <ListItemIcon style={{ minWidth: 30 }}>
                            <Icon name={tmpServiceName} />
                        </ListItemIcon>
                        <ListItemText primary={tmpServiceName} />
                    </ListItem>
                );
            }

            const projects: any[] = [];
            const tmpProjects = Object.keys(
                auth.user.authority.ProjectServiceMap
            );
            tmpProjects.sort();
            for (const tmpProjectName of tmpProjects) {
                projects.push(
                    <List
                        key={tmpProjectName}
                        disablePadding={true}
                        dense={true}
                    >
                        <ListItem
                            button={true}
                            dense={true}
                            className={classes.nested}
                            onClick={event =>
                                this.handleProjectClick(event, tmpProjectName)
                            }
                        >
                            <ListItemIcon style={{ minWidth: 30 }}>
                                <DashboardIcon />
                            </ListItemIcon>
                            <ListItemText
                                inset={true}
                                primary={tmpProjectName}
                            />
                        </ListItem>
                    </List>
                );
            }

            return (
                <div>
                    <Divider />
                    <List dense={true}>
                        <ListItem
                            button={true}
                            dense={true}
                            selected={serviceName === "/Service/Home"}
                        >
                            <ListItemIcon style={{ minWidth: 30 }}>
                                <HomeIcon />
                            </ListItemIcon>
                            <ListItemText primary="Home" />
                        </ListItem>

                        <ListItem
                            button={true}
                            dense={true}
                            onClick={this.handleOpenProjectsClick}
                        >
                            <ListItemIcon style={{ minWidth: 30 }}>
                                <DashboardIcon />
                            </ListItemIcon>
                            <ListItemText inset={true} primary={projectText} />
                            {this.state.openProjects ? (
                                <ExpandLess />
                            ) : (
                                <ExpandMore />
                            )}
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
        this.setState({ openProjects: false });
    };

    private handleClickService = (serviceName, event) => {
        this.props.dispatchServiceGetIndex(this.props.projectName, serviceName);
    };
}

function mapStateToProps(state, ownProps) {
    const auth = state.auth;
    const { projectName, serviceName } = state.service;

    return { auth, projectName, serviceName };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {
        dispatchServiceGetIndex: (projectName, serviceName) => {
            dispatch(
                actions.service.serviceGetIndex({ projectName, serviceName })
            );
        }
    };
}

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(withStyles(styles)(LeftSidebar));
