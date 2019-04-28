import React, {Component} from 'react';
import { connect } from 'react-redux';

import { withStyles } from '@material-ui/core/styles';
import Dialog from '@material-ui/core/Dialog';

import BasicForm from './BasicForm'


class IndexForm extends Component {
  render() {
    const { classes, data, index } = this.props

    let rawData = data[index.DataKey]
    if (!rawData) {
      return null
    }

    let submitName = index.SubmitAction
    let queryKind = index.SubmitAction + index.DataKey
		return <BasicForm data={data} index={index} rawData={rawData} queryKind={queryKind}
      submitButtonName={index.SubmitAction} />
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
)(withStyles(styles, {withTheme: true})(IndexForm));
