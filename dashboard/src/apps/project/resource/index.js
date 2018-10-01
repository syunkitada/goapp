import React, {Component} from 'react';
import { connect } from 'react-redux';
import Dashboard from '../../../components/Dashboard'
import actions from '../../../actions'


class ProjectResource extends Component {
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

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

export default connect(
  mapStateToProps,
)(ProjectResource)
