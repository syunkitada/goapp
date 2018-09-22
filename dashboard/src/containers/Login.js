import { connect } from 'react-redux';
import Login from '../components/Login';
import actions from '../actions'

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    onSubmit: (e) => {
      e.preventDefault()
      const name = e.target.name.value.trim()
      const password = e.target.password.value.trim()
      dispatch(actions.auth.login(name, password));
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Login)
