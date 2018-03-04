import * as React from 'react';

import FullWidthSection from '../../ui/FullWidthSection';

import styles from './Hero.scss';

const Hero: React.SFC = (): JSX.Element => (
  <FullWidthSection
    backgroundFixed={true}
    backgroundImage={require('../../images/abstract_network.jpg')}>
    <div className={styles.content}>
      <h1 className={styles.title}>
        Explore the Stellar network
      </h1>
      <h2>
        Browse the blockchain, view the orderbook, <br/>
        and analyze trends.
      </h2>
    </div>
  </FullWidthSection>
);

export default Hero;
