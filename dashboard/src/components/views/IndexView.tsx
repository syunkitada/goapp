import * as React from "react";
import { connect } from "react-redux";

import BasicView from "./BasicView";

import data_utils from "../../lib/data_utils";

interface IIndexView {
    render;
    routes;
    data;
    index;
}

class IndexView extends React.Component<IIndexView> {
    public render() {
        const { render, routes, data, index } = this.props;

        console.log("DEBUG render IndexView", data, index);
        const rawData = data[index.DataKey];
        if (!rawData) {
            return null;
        }

        const queryKind = index.SubmitAction + index.DataKey;
        return (
            <BasicView
                data={data}
                index={index}
                render={render}
                routes={routes}
                rawData={rawData}
                queryKind={queryKind}
            />
        );
    }
}

function mapStateToProps(state, ownProps) {
    const data = data_utils.getIndexDataFromState(state, ownProps.index);
    return { data };
}

function mapDispatchToProps(dispatch, ownProps) {
    return {};
}

export default connect(mapStateToProps, mapDispatchToProps)(IndexView);
