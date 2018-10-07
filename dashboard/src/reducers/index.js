import { combineReducers } from 'redux'

import auth from '../apps/auth/reducers';
import home from '../apps/home/reducers';
import resource from '../apps/project/resource/reducers';

export default combineReducers({
  auth,
  home,
  resource,
});
