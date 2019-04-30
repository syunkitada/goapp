import * as React from 'react';
import {connect} from 'react-redux';

import actions from '../../../../actions';
import MsgSnackbar from './MsgSnackbar';

interface IRequestErrSnackbar {
  error;
  onClose;
}

class RequestErrSnackbar extends React.Component<IRequestErrSnackbar> {
  render() {
    const {error} = this.props;

    if (!error) {
      return null;
    }

    let variant = 'error';
    let vertical = 'bottom';
    let horizontal = 'left';
    let msg = error.errCode + ': ' + error.err;

    return (
      <MsgSnackbar
        open
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
    error: error,
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
