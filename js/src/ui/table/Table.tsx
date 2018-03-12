import * as React from 'react';

import styles from './Table.scss';

type TableProps = {
  children: JSX.Element | Array<JSX.Element | Array<JSX.Element>>,
};

const Table: React.SFC<TableProps> = (props: TableProps): JSX.Element => (
  <div className={styles.root}>
    {props.children}
  </div>
);

export default Table;
