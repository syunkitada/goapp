import React, {Component} from 'react';
import { connect } from 'react-redux';

import { withStyles } from '@material-ui/core/styles';
import Dialog from '@material-ui/core/Dialog';

import BasicForm from '../forms/BasicForm'


class FormDialog extends Component {
  render() {
    const { classes, data, open, action, onClose } = this.props
    let title = action.Name + " " + action.DataKind
    let queryKind = action.Name + action.DataKind

		return (
        <div>
          <Dialog
            open={open}
            onClose={onClose}
            aria-labelledby="form-dialog-title"
          >
            <BasicForm onClose={onClose} data={data} index={action} title={title} queryKind={queryKind}
              submitButtonName={action.Name} />
          </Dialog>
        </div>
      );
  }
}

function mapStateToProps(state, ownProps) {
  return {}
}

function mapDispatchToProps(dispatch, ownProps) {
  return {}
}

const styles = theme => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles, {withTheme: true})(FormDialog));
