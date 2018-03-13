import * as React from 'react';

import axios from 'axios';
import classNames from 'classnames';

import RegularContainer from '../ui/RegularContainer';
import Table from '../ui/table/Table';
import TableHeader from '../ui/table/TableHeader';

import typography from '../typography.scss';
import styles from './Ledgers.scss';

type State = {
  ledgers: Array<Ledger>,
};

export default class Ledgers extends React.Component<{}, State> {
  state = {
    ledgers: [],
  };

  async componentDidMount(): Promise<void> {
    const ledgers = await axios('https://horizon.stellar.org/ledgers?order=desc&limit=30');
    this.setState })
  }

  render(): JSX.Element {
    return (
      <div className={styles.root}>
        <RegularContainer>
          <h2 className={classNames(styles.title, typography.weightMedium)}>
            All ledgers
          </h2>
        </RegularContainer>
      </div>
    );
  }
}
