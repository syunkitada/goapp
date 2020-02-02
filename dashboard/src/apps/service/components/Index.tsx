import * as React from "react";
import { connect } from "react-redux";

import IndexForm from "../../../components/forms/IndexForm";
import SearchForm from "../../../components/forms/SearchForm";
import RoutePanels from "../../../components/panels/RoutePanels";
import Panes from "../../../components/panes/Panes";
import GetMsgSnackbar from "../../../components/snackbars/GetMsgSnackbar";
import RequestErrSnackbar from "../../../components/snackbars/RequestErrSnackbar";
import SubmitMsgSnackbar from "../../../components/snackbars/SubmitMsgSnackbar";
import IndexTable from "../../../components/tables/IndexTable";
import Tabs from "../../../components/tabs/Tabs";
import IndexView from "../../../components/views/IndexView";

import actions from "../../../actions";
import logger from "../../../lib/logger";

function renderIndex(routes, data, index) {
  if (!index) {
    return <div>Not Found</div>;
  }
  logger.info("Index", "renderIndex:", index.Kind, index.Name, routes);
  switch (index.Kind) {
    case "Msg":
      return <div>{index.Name}</div>;
    case "RoutePanels":
      return (
        <RoutePanels
          render={renderIndex}
          routes={routes}
          data={data}
          index={index}
        />
      );
    case "RouteTabs":
      return (
        <Tabs render={renderIndex} routes={routes} data={data} index={index} />
      );
    case "RoutePanes":
      return (
        <Panes render={renderIndex} routes={routes} data={data} index={index} />
      );
    case "Table":
      return (
        <IndexTable
          render={renderIndex}
          routes={routes}
          index={index}
          data={data}
        />
      );
    case "View":
      return (
        <IndexView
          render={renderIndex}
          routes={routes}
          index={index}
          data={data}
        />
      );
    case "SearchForm":
      return <SearchForm routes={routes} index={index} data={data} />;
    case "Form":
      return <IndexForm routes={routes} index={index} data={data} />;
    default:
      return <div>Unsupported Kind: {index.Kind}</div>;
  }
}

interface IIndex {
  match;
  service;
  serviceName;
  projectName;
  getIndex;
}

class Index extends React.Component<IIndex> {
  public state = {
    openAlertSnackbar: true,
    traceMsgMap: {}
  };

  public componentWillMount() {
    logger.info("Index", "componentWillMount()");
    const { getIndex } = this.props;
    getIndex();
  }

  public render() {
    const { service, serviceName, projectName, getIndex } = this.props;
    logger.info("Index", "render", projectName, serviceName);

    if (
      service.serviceName !== serviceName ||
      service.projectName !== projectName
    ) {
      getIndex();
      return null;
    }

    let state: any = null;
    if (projectName) {
      state = service.projectServiceMap[projectName][serviceName];
    } else {
      state = service.serviceMap[serviceName];
    }

    if (state.isFetching) {
      return <div>Fetching...</div>;
    }

    const routes = [this.props];
    let html: any = null;
    if (state.Index) {
      html = renderIndex(routes, state.Data, state.Index.View);
    }

    return (
      <div>
        {html}
        <RequestErrSnackbar />
        <GetMsgSnackbar />
        <SubmitMsgSnackbar />
      </div>
    );
  }
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth;
  const service = state.service;
  const { match } = ownProps;

  return {
    auth,
    match,
    projectName: match.params.project,
    service,
    serviceName: match.params.service
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getIndex: () => {
      dispatch(actions.service.serviceGetIndex({ route: ownProps }));
    }
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(Index);
