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
    const {routes, index} = this.props;
    const route = routes.slice(-1)[0];
    const beforeRoute = routes.slice(-2)[0];

    for (let i = 0, len = index.Panels.length; i < len; i++) {
      const panel = index.Panels[i];
      if (route.match.path === beforeRoute.match.path + panel.Route) {
        this.props.getQueries(routes, panel);
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
    const {routes, index} = this.props;

    const beforeRoute = routes.slice(-2)[0];
    for (let i = 0, len = index.Panels.length; i < len; i++) {
      const panel = index.Panels[i];
      if (expandedPath === beforeRoute.match.path + panel.Route) {
        this.props.getQueries(routes, panel);
        break;
      }
    }

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
    getQueries: (routes, panel) => {
      const route = routes[routes.length - 1];
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

      let view: any = {};
      switch (panel.Kind) {
        case 'Table':
        case 'View':
          view = panel;
          break;
        case 'RouteTabs':
          for (let i = 0, len = panel.Tabs.length; i < len; i++) {
            const tab = panel.Tabs[i];
            if (route.match.params[panel.TabParam] === tab.Name) {
              view = tab;
              break;
            }
          }
          break;
        case 'RoutePanes':
          for (let i = 0, len = panel.Panes.length; i < len; i++) {
            const pane = panel.Panes[i];
            if (route.match.params[panel.PaneParam] === pane.Name) {
              view = pane;
              break;
            }
          }
          break;
        default:
          logger.warning(
            'ExpansionPanels',
            'Unsupported panel.Kind',
            panel.Kind,
          );
          break;
      }
      if (view.DataQueries) {
        dispatch(
          actions.service.serviceGetQueries({
            isSync: view.IsSync,
            params: route.match.params,
            queries: view.DataQueries,
            searchQueries,
          }),
        );
      }
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(ExpansionPanels));
