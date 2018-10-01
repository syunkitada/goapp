import { combineReducers } from 'redux'

import auth from '../apps/auth/reducers';
import home from '../apps/home/reducers';

export default combineReducers({
  auth,
  home,
});
