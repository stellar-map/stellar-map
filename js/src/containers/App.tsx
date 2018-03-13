import * as React from 'react';
import { Route, Switch } from 'react-router-dom';

import Footer from '../components/main/Footer';
import Header from '../components/main/Header';
import Home from './Home';
import Ledgers from './Ledgers';

import styles from './App.scss';

export default class App extends React.Component {
  renderDarkHeader = (): JSX.Element => {
    return <Header shade='dark' />;
  };

  render() {
    return (
      <div className={styles.root}>
        <Switch>
          <Route exact={true} path='/' component={Header} />
          <Route path='/(.+)' render={this.renderDarkHeader} />
        </Switch>
        <div className={styles.body}>
          <Switch>
            <Route exact={true} path='/' component={Home} />
            <Route path='/ledgers' component={Ledgers} />
          </Switch>
        </div>
        <Footer />
      </div>
    );
  }
}
