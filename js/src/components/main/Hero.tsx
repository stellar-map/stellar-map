import * as React from 'react';

import RegularContainer from '../../ui/RegularContainer';

import styles from './Hero.scss';

const Hero: React.SFC = (): JSX.Element => (
  <div className={styles.hero}>
    <RegularContainer>
      <h1 className={styles.title}>
        Explore the Stellar network
      </h1>
      <h2>
        Browse the blockchain, view the orderbook, <br/>
        and analyze trends.
      </h2>
    </RegularContainer>
  </div>
);

export default Hero;
