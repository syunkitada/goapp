import React, {Component} from 'react';
import Dashboard from '../../components/Dashboard'
import { connect } from 'react-redux';
import actions from '../../actions'


class Home extends Component {
  componentWillMount() {
    this.props.syncState()
  }

  render() {
    const {match, auth} = this.props

    if (!auth.user) {
      return null
    }

    return (
      <Dashboard match={match}>
        <div>
          <h2>Home</h2>
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

function mapDispatchToProps(dispatch, ownProps) {
  return {
    syncState: () => {
      dispatch(actions.home.homeSyncState());
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Home)
