import * as React from "react";
import { connect } from "react-redux";

import Auth from "../apps/auth";
import Login from "../apps/auth/components/Login";

import Service from "../apps/service";

interface IRouter {
    auth: any;
}

class Router extends React.Component<IRouter> {
    public render() {
        const { auth } = this.props;

        if (!auth.user) {
            return (
                <Auth>
                    <Login />
                </Auth>
            );
        }

        return <Service />;
    }
}

function mapStateToProps(state, ownProps) {
    const auth = state.auth;

    return {
        auth
    };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {};
}

export default connect(mapStateToProps, mapDispatchToProps)(Router);
