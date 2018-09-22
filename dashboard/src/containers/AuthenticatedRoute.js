import React, {Component} from 'react'
import { connect } from 'react-redux'
import { BrowserRouter, Route, Link, Switch, Redirect, withRouter } from 'react-router-dom';

class AuthenticatedRoute extends Component {
  render() {
    const { component: Component, auth, ...rest } = this.props
    return (
      <Route {...rest}
        render={props =>
          auth.user ? (
            <Component {...props} />
          ) : (
            <Redirect
              to={{
                pathname: "/login",
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
)(AuthenticatedRoute)
