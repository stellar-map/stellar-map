import * as React from 'react';

import styles from './Table.scss';

type Props = {
  orderedHeaders: Array<string>,
  rows: Array<object>,
};

const Table: React.SFC<Props> = (props: Props): JSX.Element => (
  <table>
    <thead>
      <tr className={styles.header}>
        {props.orderedHeaders.map(header => (
          <td className={styles.cell}>{header}</td>
        ))}
      </tr>
    </thead>
    <tbody>
      {props.rows.map(rowData => (
        <tr className={styles.dataRow}>
          {props.orderedHeaders.map(header => (
            <td className={styles.cell}>{rowData[header]}</td>
          ))}
        </tr>
      ))}
    </tbody>
  </table>
);

export default Table;
