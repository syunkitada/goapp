import * as React from 'react';
import {connect} from 'react-redux';

import BasicForm from './BasicForm';

interface IIndexForm {
  data;
  index;
}

class IndexForm extends React.Component<IIndexForm> {
  render() {
    const {data, index} = this.props;

    let rawData = data[index.DataKey];
    if (!rawData) {
      return null;
    }

    let queryKind = index.SubmitAction + index.DataKey;
    return (
      <BasicForm
        data={data}
        index={index}
        rawData={rawData}
        queryKind={queryKind}
        submitButtonName={index.SubmitAction}
      />
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
)(IndexForm);
