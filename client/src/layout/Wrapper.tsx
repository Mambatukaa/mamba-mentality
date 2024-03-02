import React from 'react';
import Header from './Header';
import SideBar from './SideBar';
import Body from './Body';

const Wrapper: React.FC = () => {
  return (
    <div className="grid sm:grid-rows-[1fr_20fr] grid-cols-[1fr_7fr] w-screen h-screen">
      <Header />
      <SideBar />
      <Body />
    </div>
  );
};

export default Wrapper;
