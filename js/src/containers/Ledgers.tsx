import * as React from 'react';

import classNames from 'classnames';

import RegularContainer from '../ui/RegularContainer';
import Table from '../ui/table/Table';

import typography from '../typography.scss';
import styles from './Ledgers.scss';

const Ledgers = (): JSX.Element => (
  <div className={styles.root}>
    <RegularContainer>
      <h2 className={classNames(styles.title, typography.weightMedium)}>
        All ledgers
      </h2>
      <Table>
      </Table>
    </RegularContainer>
  </div>
);

export default Ledgers;
