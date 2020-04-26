import * as React from "react";
import { connect } from "react-redux";

import Dashboard from "../../components/frames/Dashboard";

import { MuiThemeProvider } from "@material-ui/core/styles";

import actions from "../../actions";
import logger from "../../lib/logger";
import theme_utils from "../../lib/theme_utils";

import components from "../../components";

interface IService {
    auth: any;
    history: any;
    match: any;
    startBackgroundSync: any;
    service: any;
    serviceName: any;
    projectName: any;
    getIndex: any;
}

class Service extends React.Component<IService> {
    public componentWillMount() {
        logger.info("Service", "componentWillMount()");
        this.props.startBackgroundSync();
    }

    public componentWillUnmount() {
        logger.info("Service", "componentWillUnmount()");
        const { getIndex } = this.props;
        getIndex();
    }

    public render() {
        const {
            match,
            history,
            auth,
            service,
            serviceName,
            projectName,
            getIndex
        } = this.props;

        if (!auth.user) {
            return null;
        }

        if (
            service.serviceName !== serviceName ||
            service.projectName !== projectName
        ) {
            getIndex();
            return null;
        }

        let state: any = null;
        if (projectName) {
            state = service.projectServiceMap[projectName][serviceName];
        } else {
            state = service.serviceMap[serviceName];
        }

        let content: any;
        if (state.isFetching) {
            content = <div>Fetching...</div>;
        } else {
            content = components.renderIndex(service.rootIndex);
        }

        return (
            <MuiThemeProvider theme={theme_utils.getTheme(auth.theme)}>
                <Dashboard match={match} history={history}>
                    {content}
                </Dashboard>
            </MuiThemeProvider>
        );
    }
}

function mapStateToProps(state, ownProps) {
    const auth = state.auth;
    const match = ownProps.match;
    const service = state.service;

    return {
        auth,
        match,
        projectName: match.params.project,
        service,
        serviceName: match.params.service
    };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {
        getIndex: () => {
            dispatch(actions.service.serviceGetIndex({ route: ownProps }));
        },
        startBackgroundSync: () => {
            dispatch(actions.service.serviceStartBackgroundSync());
        }
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(Service);
