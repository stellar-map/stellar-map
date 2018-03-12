import * as React from 'react';

import RegularContainer from '../../ui/RegularContainer';

import StellarRocket from '../../images/stellar_rocket.png';
import styles from './Header.scss';

const Header: React.SFC = (): JSX.Element => (
  <header className={styles.header}>
    <RegularContainer>
      <div className={styles.content}>
        <div className={styles.main}>
          <a href='/' className={styles.logoLink}>
            <img src={StellarRocket} className={styles.logo} />
            <span className={styles.name}>StellarMap</span>
          </a>
          <nav>
            <a href='/' className={styles.menuItem}>Blockchain</a>
            <a href='/' className={styles.menuItem}>Stats</a>
          </nav>
        </div>
        <input
          className={styles.searchInput}
          placeholder='Search by transaction, address, ledger'
        />
      </div>
    </RegularContainer>
  </header>
);
export default Header;
