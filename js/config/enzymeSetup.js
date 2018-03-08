import Enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';

new Adapter();

Enzyme.configure({ adapter: new Adapter() });
