import * as React from 'react';

import RegularContainer from '../../ui/RegularContainer';

import styles from './Footer.scss';

const Footer: React.SFC = () => (
  <div className={styles.root}>
    <RegularContainer>
      <div className={styles.copyright}>
        &copy; 2018
      </div>
    </RegularContainer>
  </div>
);

export default Footer;
