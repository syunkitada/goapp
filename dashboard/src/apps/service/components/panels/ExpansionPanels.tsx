import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import Typography from '@material-ui/core/Typography';

import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import actions from '../../../../actions';
import logger from '../../../../lib/logger';

interface IExpansionPanels extends WithStyles<typeof styles> {
  render;
  routes;
  data;
  index;
  getQueries;
}

class ExpansionPanels extends React.Component<IExpansionPanels> {
  public state = {
    expanded: null,
    expandedUrl: null,
  };

  public componentWillMount() {
    const {routes, index, data} = this.props;
    const route = routes.slice(-1)[0];
    const beforeRoute = routes.slice(-2)[0];
    console.log('DEBUG: ExpansionPanels.componentWillMount');

    const location = route.location;
    const queryStr = decodeURIComponent(location.search);
    let searchQueries = {};
    try {
      const value = queryStr.match(new RegExp('[?&]q=({.*?})(&|$|#)'));
      if (value) {
        searchQueries = JSON.parse(value[1]);
      }
    } catch (e) {
      console.log('Ignored failed parse', queryStr);
    }

    for (let i = 0, len = index.Panels.length; i < len; i++) {
      const panel = index.Panels[i];
      if (
        panel.ExpectedDataKeys &&
        panel.GetQueries &&
        route.match.path === beforeRoute.match.path + panel.Route
      ) {
        let isInit = false;
        for (
          let j = 0, keysLen = panel.ExpectedDataKeys.length;
          j < keysLen;
          j++
        ) {
          if (!data[panel.ExpectedDataKeys[j]]) {
            isInit = true;
            break;
          }
        }
        if (isInit && panel.GetQueries) {
          this.props.getQueries(
            panel.GetQueries,
            searchQueries,
            panel.IsSync,
            route.match.params,
          );
        }
        break;
      }
    }
  }

  public render() {
    const {classes, render, routes, data, index} = this.props;
    const {expanded, expandedUrl} = this.state;

    logger.info('ExpansionPanels', 'render()', routes);

    const route = routes.slice(-1)[0];
    const beforeRoute = routes.slice(-2)[0];
    let expandedPath: any = null;

    if (expanded === null) {
      expandedPath = route.match.path;
    } else {
      expandedPath = expanded;
    }

    if (
      expandedUrl !== null &&
      expandedPath !== route.match.path &&
      expandedUrl !== route.match.url
    ) {
      expandedPath = route.match.path;
    }

    const panels: any[] = [];
    for (let i = 0, len = index.Panels.length; i < len; i++) {
      const panel = index.Panels[i];
      panels.push(
        <ExpansionPanel
          key={i}
          expanded={expandedPath === beforeRoute.match.path + panel.Route}
          onChange={() =>
            this.handleChange(
              beforeRoute.match.path + panel.Route,
              route.match.url,
            )
          }>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <Typography variant="subtitle1">
              {panel.Name} {route.match.params[panel.Subname]}
            </Typography>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails style={{padding: 0}}>
            {render(routes, data, panel)}
          </ExpansionPanelDetails>
        </ExpansionPanel>,
      );
    }

    return <div className={classes.root}>{panels}</div>;
  }

  private handleChange = (expandedPath, expandedUrl) => {
    console.log('TODO GetData for Panel on handleChangeExpansionPanels');
    this.setState({
      expanded: expandedPath,
      expandedUrl,
    });
  };
}

const styles = (theme: Theme): StyleRules =>
  createStyles({
    root: {
      backgroundColor: theme.palette.background.paper,
      flexGrow: 1,
      width: '100%',
    },
  });

function mapStateToProps(state, ownProps) {
  const auth = state.auth;

  return {auth};
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getQueries: (queries, searchQueries, isSync, params) => {
      dispatch(
        actions.service.serviceGetQueries({
          isSync,
          params,
          queries,
          searchQueries,
        }),
      );
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(ExpansionPanels));
