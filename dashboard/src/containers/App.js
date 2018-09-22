import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import actions from '../actions'

class App extends Component {
  componentWillMount() {
    if (!this.props.auth.isSyncState) {
      console.log("Debug App WillMount")
      this.props.syncState()
    }
  }

  render() {
    const { auth, children } = this.props;

    return (
      <div>
        {children}
      </div>
    )
  }
}

App.propTypes = {
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
      dispatch(actions.auth.syncState());
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
