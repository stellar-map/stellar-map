import * as React from 'react';

import Card from '../../ui/Card';
import FullWidthSection from '../../ui/FullWidthSection';
import Table from '../../ui/table/Table';
import TableCell from '../../ui/table/TableCell';
import TableHeader from '../../ui/table/TableHeader';
import TableRow from '../../ui/table/TableRow';

import typography from '../../typography.scss';
import styles from './BlockchainSection.scss';

const COLUMN_TO_SIZE: { [header: string]: TableCellSizes } = {
  sequence: 'size1',
  transactions: 'size4',
  operations: 'size2',
  closed_at: 'size8',
};

const BlockchainSection: React.SFC = (): JSX.Element => {
  // TODO: Replace with real data
  const orderedHeaders = [
    {
      key: 'sequence',
      label: 'Sequence',
    },
    {
      key: 'transactions',
      label: 'Transactions',
    },
    {
      key: 'operations',
      label: 'Operations',
    },
    {
      key: 'closed_at',
      label: 'Closed At',
    },
  ];
  const rows = [
    {
      sequence: '16612351',
      transactions: 2,
      operations: 2,
      closed_at: '5/5/18 3:20PM PST',
    },
    {
      sequence: '16612350',
      transactions: 5,
      operations: 12,
      closed_at: '5/5/18 3:18PM PST',
    },
    {
      sequence: '16612349',
      transactions: 1,
      operations: 20,
      closed_at: '5/5/18 3:10PM PST',
    },
  ];

  return (
    <FullWidthSection title='Examine blockchain details'>
      <div className={styles.content}>
        <div className={styles.cardContainer}>
          <Card>
            <div className={styles.cardHeading}>
              <h2 className={`${typography.weightMedium}`}>Latest ledgers</h2>
              <a href='/'>See all &rarr;</a>
            </div>
            <Table>
              <TableHeader columns={orderedHeaders} columnToSize={COLUMN_TO_SIZE} />
              {rows.map((rowData, i) => (
                <TableRow key={i}>
                  {orderedHeaders.map((header) => (
                    <TableCell
                      column={header.key}
                      key={`${header}-${rowData[header.key]}`}
                      size={COLUMN_TO_SIZE[header.key]}>
                      {rowData[header.key]}
                    </TableCell>
                  ))}
                </TableRow>
              ))}
            </Table>
          </Card>
        </div>
        <div className={styles.cardContainer}>
          <Card>
            <div className={styles.cardHeading}>
              <h2 className={`${typography.weightMedium}`}>Latest transactions</h2>
              <a href='/'>See all &rarr;</a>
            </div>
            <Table>
              {/* <TableHeader columnNames={orderedHeaders} />
              <tbody>
                {rows.map((rowData, i) => (
                  <tr className={`${styles.tableRow} ${typography.textSmall}`} key={i}>
                    {orderedHeaders.map(header => (
                      <TableCell key={`${header}-${rowData[header]}`}>{rowData[header]}</TableCell>
                    ))}
                  </tr>
                ))}
              </tbody> */}
            </Table>
          </Card>
        </div>
      </div>
    </FullWidthSection>
  );
};

export default BlockchainSection;
