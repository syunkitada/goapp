import * as React from 'react';
import {connect} from 'react-redux';
import {Redirect, Route} from 'react-router-dom';

import logger from '../lib/logger';

interface IAuthRoute {
  component;
  auth;
}

class AuthRoute extends React.Component<IAuthRoute> {
  public render() {
    const {component: Component, auth, ...rest} = this.props;
    logger.info('AuthRoute', 'render()');
    return (
      <Route
        {...rest}
        render={props =>
          auth.user ? (
            <Component {...props} />
          ) : (
            <Redirect
              to={{
                pathname: '/login',
                state: {from: props.location},
              }}
            />
          )
        }
      />
    );
  }
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth;

  return {
    auth,
  };
}

export default connect(mapStateToProps)(AuthRoute);
