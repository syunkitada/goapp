import {applyMiddleware, createStore} from 'redux';
import createSagaMiddleware from 'redux-saga';

import reducers from '../reducers/';
import rootSaga from '../sagas';

function configureStore() {
  const sagaMiddleware = createSagaMiddleware();

  const newStore = createStore(reducers, applyMiddleware(sagaMiddleware));

  sagaMiddleware.run(rootSaga);

  return newStore;
}

const store = configureStore();

export default store;
