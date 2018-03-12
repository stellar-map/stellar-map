import * as React from 'react';

import RegularContainer from '../../ui/RegularContainer';

import typography from '../../typography.scss';
import styles from './Footer.scss';

const Footer: React.SFC = () => (
  <section className={styles.root}>
    <RegularContainer>
      <div className={`${styles.copyright} ${typography.textXsmall}`}>
        &copy; 2018 StellarMap
      </div>
    </RegularContainer>
  </section>
);

export default Footer;
