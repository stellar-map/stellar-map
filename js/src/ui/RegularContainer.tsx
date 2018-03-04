import * as React from 'react';

import styles from './RegularContainer.scss';

type Props = {
  children: JSX.Element | Array<JSX.Element | boolean | Array<JSX.Element>>,
};

const RegularContainer: React.SFC<Props> = (props: Props): JSX.Element => (
  <div className={styles.root}>
    {props.children}
  </div>
);

export default RegularContainer;
