import * as React from 'react';
import { connect } from 'react-redux'
import { Route, Redirect } from 'react-router-dom';

import logger from '../lib/logger';

class AuthRoute extends React.Component {
  render() {
    const { component: React.Component, auth, ...rest } = this.props
    logger.info('AuthRoute', 'render()')
    return (
      <Route {...rest}
        render={props =>
          auth.user ? (
            <React.Component {...props} />
          ) : (
            <Redirect
              to={{
                pathname: '/login',
                state: { from: props.location }
              }}
            />
          )
        }
      />
    );
  }
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

export default connect(
  mapStateToProps,
)(AuthRoute)
