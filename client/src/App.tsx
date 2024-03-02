import './style.css';
import Wrapper from './layout/Wrapper';
import { BrowserRouter } from 'react-router-dom';

const App = () => {
  return (
    <BrowserRouter>
      <Wrapper />
    </BrowserRouter>
  );
};

export default App;
