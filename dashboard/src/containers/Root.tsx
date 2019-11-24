import * as React from 'react';
import {Provider} from 'react-redux';
import {BrowserRouter, Route, Switch} from 'react-router-dom';

import store from '../store';

import {MuiThemeProvider} from '@material-ui/core/styles';
import NotFound from '../components/NotFound';
import AuthRoute from './AuthRoute';

import Auth from '../apps/auth';
import Login from '../apps/auth/components/Login';
import Service from '../apps/service';

import darkTheme from '../components/themes/darkTheme';
import lightTheme from '../components/themes/lightTheme';

export default class Root extends React.Component {
  public render() {
    const themeName: string = 'light';
    let theme: any = null;
    if (themeName === 'dark') {
      theme = darkTheme;
    } else {
      theme = lightTheme;
    }
    return (
      <Provider store={store}>
        <MuiThemeProvider theme={theme}>
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
        </MuiThemeProvider>
      </Provider>
    );
  }
}
