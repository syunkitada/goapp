import React, {Component} from 'react';
import { connect } from 'react-redux';

import Dashboard from '../../components/Dashboard'
import Index from './components/Index'

import actions from '../../actions'


class Service extends Component {
  componentWillMount() {
    console.log("componentWillMount Service")
    this.props.startBackgroundSync()
  }

  componentWillUnmount() {
    console.log("componentWillUnmount Service")
  }

  render() {
    const {match, auth} = this.props

    if (!auth.user) {
      return null
    }

    return (
      <Dashboard match={match}>
        <Index match={match} />
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
