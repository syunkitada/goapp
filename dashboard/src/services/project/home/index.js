import React, {Component} from 'react';
import Dashboard from '../../../components/Dashboard'
import { connect } from 'react-redux';
import actions from '../../../actions'


class ProjectHome extends Component {
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
          <h2>Project Home</h2>
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
)(ProjectHome)
