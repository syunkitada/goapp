import { combineReducers } from 'redux'

import auth from '../apps/auth/reducers';
import service from '../apps/service/reducers';

export default combineReducers({
  auth,
  service,
});
