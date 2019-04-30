import * as React from 'react';
import {connect} from 'react-redux';
import {Route} from 'react-router-dom';

import ExpansionPanels from './ExpansionPanels';
import logger from '../../../../lib/logger';

interface IRoutePanels {
  render;
  routes;
  data;
  index;
}

class RoutePanels extends React.Component<IRoutePanels> {
  render() {
    const {render, routes, data, index} = this.props;

    let beforeRoute = routes.slice(-1)[0];
    logger.info(['RoutePanels', 'render()', beforeRoute]);
    console.log(index.Panels);

    return (
      <div>
        {index.Panels.map(v => (
          <Route
            exact
            path={beforeRoute.match.path + v.Route}
            key={v.Name}
            render={props => {
              console.log(v.Route);
              const newRoutes = routes.slice(0);
              newRoutes.push(props);
              return (
                <ExpansionPanels
                  render={render}
                  routes={newRoutes}
                  index={index}
                  data={data}
                  root={v}
                />
              );
            }}
          />
        ))}
      </div>
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
)(RoutePanels);
