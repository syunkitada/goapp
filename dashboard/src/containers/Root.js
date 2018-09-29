import React, { Component } from 'react';
import { Provider} from 'react-redux';
import { BrowserRouter, Route, Link, Switch, Redirect, withRouter } from 'react-router-dom';

import configureStore from '../store/configureStore'

import NotFound from '../components/NotFound'
import Login from './Login'
import Logout from './Logout'
import AuthRoute from './AuthRoute'
import Dashboard from '../components/Dashboard'
import App from './App'

import Home from '../services/home'
import Chat from '../services/chat'
import Datacenter from '../services/datacenter'
import Ticket from '../services/ticket'
import Wiki from '../services/wiki'
import ProjectHome from '../services/project/home'
import ProjectResource from '../services/project/resource'

const store = configureStore()

export default class Root extends Component {
  render() {
    return (
      <Provider store={store}>
        <App>
          <BrowserRouter>
            <Switch>
              <Route path="/login" component={Login} />
              <AuthRoute path="/Home" component={Home} />
              <AuthRoute path="/Chat" component={Chat} />
              <AuthRoute path="/Datacenter" component={Datacenter} />
              <AuthRoute path="/Ticket" component={Ticket} />
              <AuthRoute path="/Wiki" component={Wiki} />
              <AuthRoute path="/Project/:project/Home" component={ProjectHome} />
              <AuthRoute path="/Project/:project/Resource" component={ProjectResource} />
              <Route component={NotFound} />
            </Switch>
          </BrowserRouter>
        </App>
      </Provider>
    )
  }
}
