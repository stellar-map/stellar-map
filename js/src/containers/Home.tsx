import * as React from 'react';

import Table from '../ui/Table';

// import styles from './Home.scss';

export default class Home extends React.Component {
  render(): JSX.Element {
    const orderedHeaders = ['ID', 'Amount', 'Timestamp'];
    const rows = [
      {
        ID: 'slskdjflskdf',
        Amount: '$150.50',
        Timestamp: '500sldkjfsl',
      },
      {
        ID: 'slskdjwerwerf',
        Amount: '$10050.50',
        Timestamp: '500sldkjfsl',
      },
    ];

    return (
      <Table orderedHeaders={orderedHeaders} rows={rows} />
    );
  }
}
