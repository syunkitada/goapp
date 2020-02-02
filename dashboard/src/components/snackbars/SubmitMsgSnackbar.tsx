import * as React from 'react';
import {connect} from 'react-redux';

import actions from '../../actions';
import {toStringFromStatusCode} from '../../lib/codes';
import logger from '../../lib/logger';
import MsgSnackbar from './MsgSnackbar';

interface ISubmitMsgSnackbar {
  open;
  tctx;
  onClose;
}

class SubmitMsgSnackbar extends React.Component<ISubmitMsgSnackbar> {
  public render() {
    const {open, tctx} = this.props;

    const vertical = 'top';
    const horizontal = 'center';
    let variant = 'info';
    let msg = '';
    if (!tctx) {
      logger.warning('tctx is null');
      msg = 'Unknown';
    } else if (tctx.StatusCode >= 150) {
      variant = 'error';
      msg = tctx.Err;
    } else if (tctx.StatusCode >= 100) {
      variant = 'warning';
      msg = tctx.Err;
    } else if (tctx.StatusCode >= 20) {
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

  private handleClose = (event, reason) => {
    // if (reason === 'clickaway') {
    //   return;
    // }

    this.props.onClose();
  };
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

export default connect(mapStateToProps, mapDispatchToProps)(SubmitMsgSnackbar);
