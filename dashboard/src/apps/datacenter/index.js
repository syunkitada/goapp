import React, {Component} from 'react';
import Dashboard from '../../components/Dashboard'
import { connect } from 'react-redux';
import actions from '../../actions'


class Datacenter extends Component {
  render() {
    const {match, auth} = this.props

    if (!auth.user) {
      return null
    }

    return (
      <Dashboard match={match}>
        <div>
          <h2>Datacenter</h2>
        </div>
      </Dashboard>
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
)(Datacenter)