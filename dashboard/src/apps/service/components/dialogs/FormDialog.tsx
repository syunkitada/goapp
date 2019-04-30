import * as React from 'react';
import {connect} from 'react-redux';

import Dialog from '@material-ui/core/Dialog';

import BasicForm from '../forms/BasicForm';

interface IFormDialog {
  data;
  open;
  action;
  onClose;
}

class FormDialog extends React.Component<IFormDialog> {
  render() {
    const {data, open, action, onClose} = this.props;
    let title = action.Name + ' ' + action.DataKind;
    let queryKind = action.Name + action.DataKind;

    return (
      <div>
        <Dialog
          open={open}
          onClose={onClose}
          aria-labelledby="form-dialog-title">
          <BasicForm
            onClose={onClose}
            data={data}
            index={action}
            title={title}
            queryKind={queryKind}
            submitButtonName={action.Name}
          />
        </Dialog>
      </div>
    );
  }
}

function mapStateToProps(state, ownProps) {
  return {};
}

function mapDispatchToProps(dispatch, ownProps) {
  return {};
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(FormDialog);
