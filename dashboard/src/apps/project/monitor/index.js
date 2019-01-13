import React, {Component} from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import Dashboard from '../../../components/Dashboard'
import actions from '../../../actions'


class ProjectMonitor extends Component {
  componentWillMount() {
    const {match, syncState} = this.props
    syncState(match.params.project)
  }

  render() {
    const {match, auth, monitor} = this.props
    console.log("DEBUG Monitor")
    console.log(monitor)

    if (!auth.user) {
      return null
    }

    const projectService = auth.user.Authority.ProjectServiceMap[match.params.project]

    if (!monitor.monitor) {
      console.log("!monitor.monitor")
      return (
        <Dashboard projectService={projectService} match={match}>
          <div>
            <h2>Monitor</h2>
          </div>
        </Dashboard>
      );
    } else {
      console.log("success monitor.monitor")
      var indexHtml = ""
      for (let index in monitor.monitor.IndexMap) {
        indexHtml += index
      }
      return (
        <Dashboard projectService={projectService} match={match}>
          <div>
            <h2>Monitor</h2>
            {indexHtml}
          </div>
        </Dashboard>
      );
    }
  }
}

ProjectMonitor.propTypes = {
  auth: PropTypes.object.isRequired,
  monitor: PropTypes.object.isRequired,
  syncState: PropTypes.func.isRequired
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth
  const monitor = state.monitor

  return {
    auth: auth,
    monitor: monitor,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    syncState: (projectName) => {
      dispatch(actions.monitor.monitorSyncState(projectName));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(ProjectMonitor)
