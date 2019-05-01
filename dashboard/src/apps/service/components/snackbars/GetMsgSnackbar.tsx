import * as React from 'react';
import {connect} from 'react-redux';

import actions from '../../../../actions';
import MsgSnackbar from './MsgSnackbar';

interface IGetMsgSnackbar {
  onClose;
  open;
  tctx;
}

class GetMsgSnackbar extends React.Component<IGetMsgSnackbar> {
  public render() {
    const {open, tctx} = this.props;

    if (!tctx) {
      return null;
    }

    let variant = 'info';
    let vertical = 'bottom';
    let horizontal = 'left';
    let msg = '';
    if (tctx.StatusCode >= 500) {
      variant = 'error';
      msg = tctx.Err;
    } else if (tctx.StatusCode >= 300) {
      variant = 'warning';
      msg = tctx.Err;
    } else {
      return null;
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
  const {openGetQueriesTctx, getQueriesTctx} = state.service;

  return {
    open: openGetQueriesTctx,
    tctx: getQueriesTctx,
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    onClose: () => {
      dispatch(actions.service.serviceCloseGetQueriesTctx());
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(GetMsgSnackbar);
