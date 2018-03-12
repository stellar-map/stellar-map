import * as React from 'react';

import classNames from 'classnames';

import typography from '../../typography.scss';
import styles from './TableRow.scss';

type Props = {
  children: JSX.Element | Array<JSX.Element>,
};

const TableCell: React.SFC<Props> = (props: Props): JSX.Element => (
  <div className={classNames(styles.root, typography.textSmall)}>
    {props.children}
  </div>
);

export default TableCell;
