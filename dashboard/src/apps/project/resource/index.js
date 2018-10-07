import React, {Component} from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import Dashboard from '../../../components/Dashboard'
import actions from '../../../actions'


class ProjectResource extends Component {
  componentWillMount() {
    const {match, syncState} = this.props
    console.log("ProjectResource willMount")
    console.log(match)
    syncState(match.params.project)
  }

  render() {
    const {match, auth} = this.props
    console.log("DEBUG")
    console.log(match.url)
    console.log(match.params.project)

    if (!auth.user) {
      return null
    }

    const projectService = auth.user.Authority.ProjectServiceMap[match.params.project]

    return (
      <Dashboard projectService={projectService} match={match}>
        <div>
          <h2>Resource</h2>
        </div>
      </Dashboard>
    );
  }
}

ProjectResource.propTypes = {
  auth: PropTypes.object.isRequired,
  syncState: PropTypes.func.isRequired
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    syncState: (projectName) => {
      dispatch(actions.resource.resourceSyncState(projectName));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(ProjectResource)
