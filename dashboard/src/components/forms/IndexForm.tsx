import * as React from 'react';
import {connect} from 'react-redux';

import BasicForm from './BasicForm';

interface IIndexForm {
  routes;
  data;
  index;
}

class IndexForm extends React.Component<IIndexForm> {
  public render() {
    const {routes, data, index} = this.props;

    const rawData = data[index.DataKey];
    if (!rawData) {
      return null;
    }

    const queryKind = index.SubmitAction;
    return (
      <BasicForm
        data={data}
        index={index}
        routes={routes}
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
