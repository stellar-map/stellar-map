import * as React from 'react';

import classNames from 'classnames';

import styles from './TableCell.scss';

type Props = {
  children: JSX.Element | Array<JSX.Element> | string,
  column: string,
  size?: TableCellSizes,
};

const TableCell: React.SFC<Props> = (props: Props): JSX.Element => (
  <div className={classNames(styles.root, styles[props.size!])}>
    {props.children}
  </div>
);
TableCell.defaultProps = {
  size: 'size1',
};

export default TableCell;
