import * as React from 'react';
import { Route, Switch } from 'react-router-dom';

import Footer from '../components/main/Footer';
import Header from '../components/main/Header';
import Home from '../containers/Home';

import styles from './App.scss';

export default class App extends React.Component {
  render() {
    return (
      <div className={styles.root}>
        <Header />
        <Switch>
          <Route exact={true} path='/' component={Home} />>
        </Switch>
        <Footer />
      </div>
    );
  }
}
