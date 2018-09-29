import { connect } from 'react-redux';
import LeftSidebar from '../components/LeftSidebar';
import actions from '../actions'

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

export default connect(
  mapStateToProps,
)(LeftSidebar)
