import * as React from 'react';

import BlockchainSection from '../components/home/BlockchainSection';
import Hero from '../components/home/Hero';

const Home = (): JSX.Element => (
  <div>
    <Hero />
    <BlockchainSection />
  </div>
);

export default Home;
