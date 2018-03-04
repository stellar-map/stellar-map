import * as React from 'react';
import { Route, Switch } from 'react-router-dom';

import Footer from '../components/main/Footer';
import Header from '../components/main/Header';
import Hero from '../components/main/Hero';
import Home from '../containers/Home';
import RegularContainer from '../ui/RegularContainer';

import styles from './App.scss';

export default class App extends React.Component {
  render() {
    return (
      <div className={styles.root}>
        <Header />
        <Hero />
        <div className={styles.body}>
          <RegularContainer>
            <Switch>
              <Route exact={true} path='/' component={Home} />>
            </Switch>
          </RegularContainer>
        </div>
        <Footer />
      </div>
    );
  }
}
