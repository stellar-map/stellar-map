import * as React from 'react';

import RegularContainer from '../../ui/RegularContainer';

import styles from './Hero.scss';

const Hero: React.SFC = (): JSX.Element => (
  <div className={styles.hero}>
    <RegularContainer>
      <h1 className={styles.title}>Stellar Blockchain Explorer</h1>
      <h2>1,000,234 transactions</h2>
    </RegularContainer>
  </div>
);

export default Hero;
