import { Route, Routes, BrowserRouter } from 'react-router-dom';
import Product from '../products/components/Products';

const routes = () => {
  return (
    <Routes>
      <Route path="/" element={<Product />} />
      <Route path="/products" element={<Product />} />
    </Routes>
  );
};

export default routes;
