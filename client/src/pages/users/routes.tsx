import { Route, Routes } from 'react-router-dom';
import Users from '../users/components/Users';

const routes = () => {
  return (
    <Routes>
      <Route path="/users" element={<Users />} />
    </Routes>
  );
};

export default routes;
