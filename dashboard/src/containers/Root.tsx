import * as React from 'react';
import {Provider} from 'react-redux';
import {BrowserRouter, Route, Switch} from 'react-router-dom';

import store from '../store';

import NotFound from '../components/NotFound';
import AuthRoute from './AuthRoute';

import Auth from '../apps/auth';
import Login from '../apps/auth/components/Login';
import Service from '../apps/service';

export default class Root extends React.Component {
  public render() {
    return (
      <Provider store={store}>
        <Auth>
          <BrowserRouter>
            <Switch>
              <Route path="/login" component={Login} />
              <AuthRoute path="/Service/:service" component={Service} />
              <AuthRoute
                path="/Project/:project/:service"
                component={Service}
              />
              <Route component={NotFound} />
            </Switch>
          </BrowserRouter>
        </Auth>
      </Provider>
    );
  }
}
