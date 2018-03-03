import * as React from 'react';

import RegularContainer from '../../ui/RegularContainer';

import logo from '../../images/logo.svg';
import styles from './Header.scss';

const Header: React.SFC = (): JSX.Element => (
  <header className={styles.header}>
    <RegularContainer>
      <div className={styles.content}>
        <div className={styles.main}>
          <a href='/' className={styles.logoLink}>
            <img src={logo} className={styles.logo} alt="logo" />
            <h1 className={styles.title}>TITLE HERE</h1>
          </a>
          <nav>
            <a href='/' className={styles.menuItem}>Home</a>
            <a href='/' className={styles.menuItem}>Blockchain</a>
            <a href='/' className={styles.menuItem}>Stats</a>
          </nav>
        </div>
        <input
          className={styles.searchInput}
          placeholder="Search by transaction, address, block"
        />
      </div>
    </RegularContainer>
  </header>
);
export default Header;
