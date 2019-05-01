import * as React from 'react';
import {connect} from 'react-redux';

import actions from '../../actions';
import logger from '../../lib/logger';

interface IAuth {
  auth;
  children;
  syncState: () => null;
}

class Auth extends React.Component<IAuth> {
  public componentWillMount() {
    if (!this.props.auth.isSyncState) {
      logger.info(['Auth', 'componentWillMount()']);
      this.props.syncState();
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
    syncState: () => {
      dispatch(actions.auth.authSyncState());
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Auth);
