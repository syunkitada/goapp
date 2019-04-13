import React, {Component} from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

import actions from '../../../../actions'
import MsgSnackbar from './MsgSnackbar'

class GetMsgSnackbar extends Component {
  handleClose = (event, reason) => {
    // if (reason === 'clickaway') {
    //   return;
    // }

    this.props.onClose();
  };

  render() {
    const { open, tctx } = this.props;

		if (!tctx) {
			return null
		}

		let variant = "info"
		let vertical = "bottom"
		let horizontal = "left"
		let msg = ""
		if (tctx.StatusCode >= 500) {
			variant="error"
			msg = tctx.Err
		} else if (tctx.StatusCode >= 300) {
			variant="warning"
			msg = tctx.Err
		} else {
      return null
		}

    return <MsgSnackbar open={open} onClose={this.handleClose}
      vertical={vertical} horizontal={horizontal} 
      variant={variant} msg={msg} />
  }
}

GetMsgSnackbar.propTypes = {
  className: PropTypes.string,
};

function mapStateToProps(state, ownProps) {
  const { openGetQueriesTctx, getQueriesTctx } = state.service;
	console.log("DEBUG GetMsgSnackbar mapStateToProps", openGetQueriesTctx)

  return {
		open: openGetQueriesTctx,
		tctx: getQueriesTctx,
	}
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    onClose: () => {
      dispatch(actions.service.serviceCloseGetQueriesTctx());
    }
	}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(GetMsgSnackbar)
