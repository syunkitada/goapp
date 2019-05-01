import * as React from 'react';
import {connect} from 'react-redux';

import actions from '../../../../actions';
import MsgSnackbar from './MsgSnackbar';

interface IRequestErrSnackbar {
  error;
  onClose;
}

class RequestErrSnackbar extends React.Component<IRequestErrSnackbar> {
  public render() {
    const {error} = this.props;

    if (!error) {
      return null;
    }

    const variant = 'error';
    const vertical = 'bottom';
    const horizontal = 'left';
    const msg = error.errCode + ': ' + error.err;

    return (
      <MsgSnackbar
        open={true}
        onClose={this.props.onClose}
        vertical={vertical}
        horizontal={horizontal}
        variant={variant}
        msg={msg}
      />
    );
  }
}

function mapStateToProps(state, ownProps) {
  const {error} = state.service;

  return {
    error,
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    onClose: () => {
      dispatch(actions.service.serviceCloseErr());
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(RequestErrSnackbar);
