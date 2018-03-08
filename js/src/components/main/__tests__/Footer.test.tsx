import * as React from 'react';

import { shallow } from 'enzyme';

import Footer from '../Footer';

describe('Footer', () => {
  it('renders correctly', () => {
    expect(shallow(
      <Footer />,
    )).toMatchSnapshot();
  });
});
