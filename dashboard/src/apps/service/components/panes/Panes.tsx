import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import Typography from '@material-ui/core/Typography';

import actions from '../../../../actions';
import logger from '../../../../lib/logger';

const styles = (theme: Theme): StyleRules =>
  createStyles({
    root: {
      backgroundColor: theme.palette.background.paper,
      flexGrow: 1,
      width: '100%',
    },
  });

interface IPanes extends WithStyles<typeof styles> {
  classes;
  render;
  routes;
  data;
  index;
  getQueries;
}

class Panes extends React.Component<IPanes> {
  public state = {
    tabRoute: null,
  };

  public componentWillMount() {
    const {routes, index, data} = this.props;
    const route = routes.slice(-1)[0];
    const beforeRoute = routes.slice(-2)[0];
    console.log('DEBUG: Panes.componentWillMount');

    for (let i = 0, len = index.Panes.length; i < len; i++) {
      const pane = index.Panes[i];
      console.log('DEBUG pane', pane);
      console.log(route.match.path, beforeRoute.match.path, pane.Route);

      console.log(
        route.match.params[pane.RouteParamKey],
        pane.RouteParamKey,
        pane.RouteParamValue,
      );

      if (
        pane.RouteParamKey &&
        pane.ExpectedDataKeys &&
        pane.GetQueries &&
        route.match.params[pane.RouteParamKey] === pane.RouteParamValue
      ) {
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
        if (isInit && pane.GetQueries) {
          this.props.getQueries(
            pane.GetQueries,
            pane.IsSync,
            route.match.params,
          );
        }
        break;
      } else if (
        pane.ExpectedDataKeys &&
        pane.GetQueries &&
        route.match.path === beforeRoute.match.path + pane.Route
      ) {
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
        if (isInit && pane.GetQueries) {
          this.props.getQueries(
            pane.GetQueries,
            pane.IsSync,
            route.match.params,
          );
        }
        break;
      }
    }
  }

  public render() {
    const {classes, render, routes, data, index} = this.props;
    logger.info('Panes', 'render()', routes);

    const route = routes[routes.length - 1];

    let tabContainer: any = null;
    for (let i = 0, len = index.Panes.length; i < len; i++) {
      const pane = index.Panes[i];
      if (route.match.params[index.PaneParam] === pane.Name) {
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
  return {};
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getQueries: (queries, isSync, params) => {
      dispatch(
        actions.service.serviceGetQueries({
          isSync,
          params,
          queries,
        }),
      );
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(Panes));
