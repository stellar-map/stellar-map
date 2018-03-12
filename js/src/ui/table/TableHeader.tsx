import * as React from 'react';

import TableCell from './TableCell';

import typography from '../../typography.scss';
import styles from './TableHeader.scss';

type Props = {
  columns: Array<{
    key: string,
    label: string,
  }>,
  columnToSize: { [column: string]: TableCellSizes },
};

const TableHeader: React.SFC<Props> = (props: Props): JSX.Element => (
  <div className={`${styles.root} ${typography.weightSemibold} ${typography.textXsmall}`}>
    {props.columns.map((columnInfo: {
      key: string,
      label: string,
    }) => (
      <TableCell
        column={columnInfo.key}
        key={`header-${columnInfo.key}`}
        size={props.columnToSize[columnInfo.key]}>
        {columnInfo.label}
      </TableCell>
    ))}
  </div>
);

export default TableHeader;
