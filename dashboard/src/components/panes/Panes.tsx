import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
  StyleRules,
  WithStyles
} from "@material-ui/core/styles/withStyles";

import Typography from "@material-ui/core/Typography";

import actions from "../../actions";
import logger from "../../lib/logger";

const styles = (theme: Theme): StyleRules =>
  createStyles({
    root: {
      backgroundColor: theme.palette.background.paper,
      flexGrow: 1,
      width: "100%"
    }
  });

interface IPanes extends WithStyles<typeof styles> {
  auth;
  classes;
  render;
  routes;
  data;
  index;
  getQueries;
}

class Panes extends React.Component<IPanes> {
  public state = {
    tabRoute: null
  };

  public componentWillMount() {
    const { routes, index, data } = this.props;
    const route = routes.slice(-1)[0];
    const beforeRoute = routes.slice(-2)[0];
    console.log("TODO DEBUG: Panes.componentWillMount");

    const location = route.location;
    const queryStr = decodeURIComponent(location.search);
    let searchQueries = {};
    try {
      const value = queryStr.match(new RegExp("[?&]q=({.*?})(&|$|#)"));
      if (value) {
        searchQueries = JSON.parse(value[1]);
      }
    } catch (e) {
      console.log("Ignored failed parse", queryStr);
    }

    for (let i = 0, len = index.Panes.length; i < len; i++) {
      const pane = index.Panes[i];
      console.log("DEBUG TODO pane", pane);
      console.log(
        "DEBUG TODO route1",
        route.match.path,
        beforeRoute.match.path,
        pane.Route
      );

      console.log(
        "DEBUG TODO route2",
        route.match.params[pane.RouteParamKey],
        pane.RouteParamKey,
        pane.RouteParamValue
      );

      if (
        pane.RouteParamKey &&
        pane.ExpectedDataKeys &&
        pane.DataQueries &&
        route.match.params[pane.RouteParamKey] === pane.RouteParamValue
      ) {
        console.log("DEBUG TODO HOGE");
        let isInit = false;
        for (
          let j = 0, keysLen = pane.ExpectedDataKeys.length;
          j < keysLen;
          j++
        ) {
          if (!data[pane.ExpectedDataKeys[j]]) {
            isInit = true;
            break;
          }
        }
        if (isInit && pane.DataQueries) {
          console.log("DEBUG isInit", isInit, pane.DataQueries);
          this.props.getQueries(pane, this.state, route, searchQueries);
        }
        break;
      } else if (
        pane.ExpectedDataKeys &&
        pane.DataQueries &&
        route.match.path === beforeRoute.match.path + pane.Route
      ) {
        console.log("DEBUG TODO HOGE2");
        let isInit = false;
        for (
          let j = 0, keysLen = pane.ExpectedDataKeys.length;
          j < keysLen;
          j++
        ) {
          if (!data[pane.ExpectedDataKeys[j]]) {
            isInit = true;
            break;
          }
        }
        if (isInit && pane.DataQueries) {
          this.props.getQueries(pane, this.state, route, searchQueries);
        }
        break;
      }
    }
  }

  public render() {
    const { classes, render, routes, data, index } = this.props;
    logger.info("Panes", "render()", routes);

    const route = routes[routes.length - 1];

    let tabContainer: any = null;
    for (let i = 0, len = index.Panes.length; i < len; i++) {
      const pane = index.Panes[i];
      console.log("TODO Panes", index.PaneParam, route.match.params, pane.Name);
      if (route.match.params[index.PaneParam] === pane.Name) {
        console.log("TODO Panes2");
        tabContainer = (
          <Typography component="div">{render(routes, data, pane)}</Typography>
        );
        break;
      }
    }

    return <div className={classes.root}>{tabContainer}</div>;
  }
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth;
  return { auth };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getQueries: (index, route, searchQueries) => {
      console.log("DEBUG TODO getQueries panes", index, route, searchQueries);
      dispatch(
        actions.service.serviceGetQueries({
          index,
          route,
          searchQueries
        })
      );
    }
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(Panes));