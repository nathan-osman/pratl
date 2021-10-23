import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import TimeAgo from 'javascript-time-ago';
import en from 'javascript-time-ago/locale/en.json';
import App from './app/App';
import { store } from './app/store';
import './index.scss';

TimeAgo.addDefaultLocale(en);

ReactDOM.render(
  <Provider store={store}>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </Provider>,
  document.getElementById('root')
);
