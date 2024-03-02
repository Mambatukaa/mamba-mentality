import { Route, Routes } from 'react-router-dom';
import Product from '../products/components/Products';

const routes = () => {
  return (
    <Routes>
      <Route path="/" element={<h1>Hello world</h1>} />
      <Route path="/products" element={<Product />} />
    </Routes>
  );
};

export default routes;
