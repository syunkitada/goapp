import React, {Component} from 'react';
import Dashboard from '../../components/Dashboard'
import { connect } from 'react-redux';
import actions from '../../actions'
import Index from './components/Index'
import Paper from '@material-ui/core/Paper';


class Service extends Component {
  componentWillMount() {
    const {match, getIndex} = this.props
    this.props.getIndex(match.params)
  }

  render() {
    const {match, auth, service} = this.props

    if (!auth.user) {
      return null
    }

    if (!service.index.Index) {
      return null
    }

    return (
      <Dashboard match={match}>
        <h2>{match.params.service}</h2>
        <Index match={match} />
      </Dashboard>
    );
  }
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const service = state.service

  return {
    auth: auth,
    service: service,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getIndex: (params) => {
      dispatch(actions.service.serviceGetIndex(params));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Service)
