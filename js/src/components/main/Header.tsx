import * as React from 'react';

import classNames from 'classnames';

import RegularContainer from '../../ui/RegularContainer';

import StellarRocket from '../../images/stellar_rocket.png';
import styles from './Header.scss';

type Props = {
  shade?: 'light' | 'dark',
};

const Header: React.SFC<Props> = (props: Props): JSX.Element => (
  <header className={classNames({ [styles.root]: true, [styles.dark]: props.shade === 'dark' })}>
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
Header.defaultProps = {
  shade: 'light',
};

export default Header;
