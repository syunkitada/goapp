import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import actions from '../../actions'
import logger from '../../lib/logger'

class Auth extends Component {
  componentWillMount() {
    if (!this.props.auth.isSyncState) {
      logger.info('Auth', 'componentWillMount()')
      this.props.syncState()
    }
  }

  render() {
    const { children } = this.props;

    return (
      <div>
        {children}
      </div>
    )
  }
}

Auth.propTypes = {
  auth: PropTypes.object.isRequired,
  children: PropTypes.object.isRequired,
  syncState: PropTypes.func.isRequired
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const children = ownProps.children

  return {
    auth: auth,
    children: children,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    syncState: () => {
      dispatch(actions.auth.authSyncState());
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Auth)
