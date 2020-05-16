import * as React from "react";
import { connect } from "react-redux";

import LinearProgress from "@material-ui/core/LinearProgress";
import Paper from "@material-ui/core/Paper";

import Dashboard from "../../components/frames/Dashboard";

import { MuiThemeProvider } from "@material-ui/core/styles";

import actions from "../../actions";
import logger from "../../lib/logger";
import theme_utils from "../../lib/theme_utils";

import Index from "../../components/Index";

interface IService {
    auth: any;
    startBackgroundSync: any;
    service: any;
    serviceName: any;
    projectName: any;
}

class Service extends React.Component<IService> {
    public componentWillMount() {
        logger.info("Service", "componentWillMount()");
        this.props.startBackgroundSync();
    }

    public render() {
        const { auth, service, serviceName, projectName } = this.props;

        let state: any = null;
        if (projectName) {
            state = service.projectServiceMap[projectName][serviceName];
        } else {
            state = service.serviceMap[serviceName];
        }

        if (!state) {
            return (
                <Paper>
                    <LinearProgress />
                </Paper>
            );
        }

        let content: any;
        if (state.isFetching) {
            content = (
                <Paper>
                    <LinearProgress />
                </Paper>
            );
        } else {
            content = <Index {...service.rootIndex} />;
        }

        return (
            <MuiThemeProvider theme={theme_utils.getTheme(auth.theme)}>
                <Dashboard>{content}</Dashboard>
            </MuiThemeProvider>
        );
    }
}

function mapStateToProps(state, ownProps) {
    const auth = state.auth;
    const service = state.service;

    return {
        auth,
        projectName: service.projectName,
        service,
        serviceName: service.serviceName
    };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {
        startBackgroundSync: () => {
            dispatch(actions.service.serviceStartBackgroundSync());
        }
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(Service);
