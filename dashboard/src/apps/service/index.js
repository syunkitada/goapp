import React, {Component} from 'react';
import { connect } from 'react-redux';

import Dashboard from '../../components/Dashboard'
import Index from './components/Index'

import actions from '../../actions'
import logger from '../../lib/logger'


class Service extends Component {
  componentWillMount() {
    logger.info("Service", "componentWillMount()")
    this.props.startBackgroundSync()
  }

  componentWillUnmount() {
    logger.info("Service", "componentWillUnmount()")
  }

  render() {
    const {match, history, auth} = this.props

    if (!auth.user) {
      return null
    }

    return (
      <Dashboard match={match} history={history}>
        <Index match={match} history={history} />
      </Dashboard>
    );
  }
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const match = ownProps.match

  return {
    match: match,
    auth: auth,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    startBackgroundSync: () => {
      dispatch(actions.service.serviceStartBackgroundSync())
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Service)
