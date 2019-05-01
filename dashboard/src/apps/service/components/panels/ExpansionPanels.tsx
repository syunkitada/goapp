import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import withStyles, {
  WithStyles,
  StyleRules,
} from '@material-ui/core/styles/withStyles';
import createStyles from '@material-ui/core/styles/createStyles';

import Typography from '@material-ui/core/Typography';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
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
    let route = routes.slice(-1)[0];
    let beforeRoute = routes.slice(-2)[0];

    for (let i = 0, len = index.Panels.length; i < len; i++) {
      let panel = index.Panels[i];
      if (
        panel.ExpectedDataKeys &&
        panel.GetQueries &&
        route.match.path === beforeRoute.match.path + panel.Route
      ) {
        let isInit = false;
        for (let j = 0, len = panel.ExpectedDataKeys.length; j < len; j++) {
          if (!data[panel.ExpectedDataKeys[j]]) {
            isInit = true;
            break;
          }
        }
        if (isInit && panel.GetQueries) {
          this.props.getQueries(
            panel.GetQueries,
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
    let {expanded, expandedUrl} = this.state;

    logger.info(['ExpansionPanels', 'render()', routes]);

    let route = routes.slice(-1)[0];
    let beforeRoute = routes.slice(-2)[0];

    if (expanded === null) {
      expanded = route.match.path;
    }

    if (
      expandedUrl !== null &&
      expanded !== route.match.path &&
      expandedUrl !== route.match.url
    ) {
      expanded = route.match.path;
    }

    return (
      <div className={classes.root}>
        {index.Panels.map(p => (
          <ExpansionPanel
            key={p.Name}
            expanded={(() => {
              return expanded === beforeRoute.match.path + p.Route;
            })()}
            onChange={() =>
              this.handleChange(
                beforeRoute.match.path + p.Route,
                route.match.url,
              )
            }>
            <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
              <Typography variant="title">
                {p.Name} {route.match.params[p.Subname]}
              </Typography>
            </ExpansionPanelSummary>
            <ExpansionPanelDetails style={{padding: 0}}>
              {render(routes, data, p)}
            </ExpansionPanelDetails>
          </ExpansionPanel>
        ))}
      </div>
    );
  }

  private handleChange = (expandedPath, expandedUrl) => {
    this.setState({
      expanded: expandedPath,
      expandedUrl: expandedUrl,
    });
  };
}

const styles = (theme: Theme): StyleRules =>
  createStyles({
    root: {
      flexGrow: 1,
      width: '100%',
      backgroundColor: theme.palette.background.paper,
    },
  });

function mapStateToProps(state, ownProps) {
  const auth = state.auth;

  return {
    auth: auth,
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    getQueries: (queries, isSync, params) => {
      dispatch(
        actions.service.serviceGetQueries({
          queries: queries,
          isSync: isSync,
          params: params,
        }),
      );
    },
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles)(ExpansionPanels));
