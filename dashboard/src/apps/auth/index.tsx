import * as React from 'react';
import {connect} from 'react-redux';

import actions from '../../actions';
import logger from '../../lib/logger';

interface IAuth {
  auth;
  children;
  loginWithToken: () => null;
}

class Auth extends React.Component<IAuth> {
  public componentWillMount() {
    if (!this.props.auth.isSyncState) {
      logger.info('Auth', 'componentWillMount()');
      this.props.loginWithToken();
    }
  }

  public render() {
    const {children} = this.props;

    return <div>{children}</div>;
  }
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth;
  const children = ownProps.children;

  return {auth, children};
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    loginWithToken: () => {
      dispatch(actions.auth.authLoginWithToken());
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Auth);
