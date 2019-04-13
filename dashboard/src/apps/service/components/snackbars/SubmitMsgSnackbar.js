import React, {Component} from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

import actions from '../../../../actions'
import MsgSnackbar from './MsgSnackbar'

class SubmitMsgSnackbar extends Component {
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
		let vertical = "top"
		let horizontal = "center"
		let msg = ""
		if (tctx.StatusCode >= 500) {
				variant="error"
				msg = tctx.Err
		} else if (tctx.StatusCode >= 300) {
				variant="warning"
				msg = tctx.Err
		} else if (tctx.StatusCode > 200) {
				variant="success"
				msg = tctx.Err
		}

    return <MsgSnackbar open={open} onClose={this.handleClose}
      vertical={vertical} horizontal={horizontal} 
      variant={variant} msg={msg} />
  }
}

SubmitMsgSnackbar.propTypes = {
  className: PropTypes.string,
  message: PropTypes.node,
  onClose: PropTypes.func,
};

function mapStateToProps(state, ownProps) {
  const { openSubmitQueriesTctx, submitQueriesTctx } = state.service;
	console.log(state.service)

  return {
		open: openSubmitQueriesTctx,
		tctx: submitQueriesTctx,
	}
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    onClose: () => {
      dispatch(actions.service.serviceCloseSubmitQueriesTctx());
    }
	}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(SubmitMsgSnackbar)
