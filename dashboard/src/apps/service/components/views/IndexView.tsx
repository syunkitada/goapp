import * as React from 'react';
import {connect} from 'react-redux';

import BasicView from './BasicView';

interface IIndexView {
  routes;
  data;
  index;
}

class IndexView extends React.Component<IIndexView> {
  public render() {
    const {routes, data, index} = this.props;

    console.log('DEBUG render IndexView', data, index);
    const rawData = data[index.DataKey];
    if (!rawData) {
      return null;
    }

    const queryKind = index.SubmitAction + index.DataKey;
    return (
      <BasicView
        data={data}
        index={index}
        routes={routes}
        rawData={rawData}
        queryKind={queryKind}
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
)(IndexView);
