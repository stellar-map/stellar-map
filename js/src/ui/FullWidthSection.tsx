import * as React from 'react';

import RegularContainer from './RegularContainer';

import typography from '../typography.scss';
import styles from './FullWidthSection.scss';

type Props = {
  backgroundColor?: string,
  backgroundFixed?: boolean,
  backgroundImage?: any,
  children: JSX.Element | Array<JSX.Element>,
  title?: string,
};

const FullWidthSection: React.SFC<Props> = (props: Props): JSX.Element => (
  <section
    className={styles.root}
    style={{
      backgroundAttachment: props.backgroundFixed ? 'fixed' : 'initial',
      backgroundColor: props.backgroundColor,
      backgroundImage: `url(${props.backgroundImage})`,
    }}>
    <RegularContainer>
      {!!props.title &&
        <h1 className={`${styles.title} ${typography.textCenter}`}>
          {props.title}
        </h1>}
      {props.children}
    </RegularContainer>
  </section>
);
FullWidthSection.defaultProps = {
  backgroundFixed: false,
};

export default FullWidthSection;
