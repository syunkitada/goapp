import * as React from 'react';
import {connect} from 'react-redux';

import Dialog from '@material-ui/core/Dialog';

import BasicForm from '../forms/BasicForm';

interface IFormDialog {
  data;
  routes;
  selected;
  open;
  action;
  onClose;
}

class FormDialog extends React.Component<IFormDialog> {
  public render() {
    const {data, routes, selected, open, action, onClose} = this.props;
    const title = action.Name + ' ' + action.DataKind;
    const queryKind = action.Name + action.DataKind;

    return (
      <div>
        <Dialog
          open={open}
          onClose={onClose}
          aria-labelledby="form-dialog-title">
          <BasicForm
            onClose={onClose}
            data={data}
            routes={routes}
            selected={selected}
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
