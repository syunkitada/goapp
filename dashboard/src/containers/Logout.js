import { connect } from 'react-redux';
import Logout from '../components/Logout';
import actions from '../actions'

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}


function mapDispatchToProps(dispatch, ownProps) {
  return {
    onClick: () => dispatch(actions.auth.logout({name: "hoge"}))
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Logout)
