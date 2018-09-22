import React, {Component} from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import {withRouter} from 'react-router-dom';


export default class Logout extends Component {
  constructor(props) {
    super(props)
  }

  render() {
    const { auth, onClick } = this.props
    console.log("logout debug")
    console.log(auth)

    if (auth.user) {
      return (
        <p>
          Welcom! {auth.user.name}
          <button onClick={onClick}>
            Sign out
          </button>
        </p>
      )
    }

    return (
      <p>You are not logged in.</p>
    )
  }
}

Logout.propTypes = {
  auth: PropTypes.object.isRequired,
  onClick: PropTypes.func.isRequired,
}
