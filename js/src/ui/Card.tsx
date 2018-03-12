import * as React from 'react';

import styles from './Card.scss';

type Props = {
  children: JSX.Element | Array<JSX.Element>,
};

const Card: React.SFC<Props> = (props: Props): JSX.Element => (
  <div className={styles.root}>
    {props.children}
  </div>
);

export default Card;
