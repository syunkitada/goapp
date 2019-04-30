import * as React from 'react';
import {connect} from 'react-redux';

import actions from '../../../../actions';
import {toStringFromStatusCode} from '../../../../lib/codes';
import MsgSnackbar from './MsgSnackbar';

interface ISubmitMsgSnackbar {
  open;
  tctx;
  onClose;
}

class SubmitMsgSnackbar extends React.Component<ISubmitMsgSnackbar> {
  handleClose = (event, reason) => {
    // if (reason === 'clickaway') {
    //   return;
    // }

    this.props.onClose();
  };

  render() {
    const {open, tctx} = this.props;

    if (!tctx) {
      return null;
    }

    let variant = 'info';
    let vertical = 'top';
    let horizontal = 'center';
    let msg = '';
    if (tctx.StatusCode >= 500) {
      variant = 'error';
      msg = tctx.Err;
    } else if (tctx.StatusCode >= 300) {
      variant = 'warning';
      msg = tctx.Err;
    } else if (tctx.StatusCode > 200) {
      variant = 'success';
      msg = toStringFromStatusCode(tctx.StatusCode);
    } else {
      msg = 'Unknown';
    }

    return (
      <MsgSnackbar
        open={open}
        onClose={this.handleClose}
        vertical={vertical}
        horizontal={horizontal}
        variant={variant}
        msg={msg}
      />
    );
  }
}

function mapStateToProps(state, ownProps) {
  const {openSubmitQueriesTctx, submitQueriesTctx} = state.service;

  return {
    open: openSubmitQueriesTctx,
    tctx: submitQueriesTctx,
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    onClose: () => {
      dispatch(actions.service.serviceCloseSubmitQueriesTctx());
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(SubmitMsgSnackbar);
