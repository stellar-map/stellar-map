import * as React from 'react';

import RegularContainer from '../../ui/RegularContainer';

import styles from './Footer.scss';

const Footer: React.SFC = () => (
  <div className={styles.root}>
    <RegularContainer>
      <div className={`${styles.copyright} text-xsmall`}>
        &copy; 2018 StellarMap
      </div>
    </RegularContainer>
  </div>
);

export default Footer;
