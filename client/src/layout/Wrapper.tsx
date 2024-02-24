import React from 'react';
import Body from './Body';
import SideBar from './SideBar';

const Wrapper: React.FC = () => {
  return (
    <div className="flex relative h-dvh w-dvw items-center justify-center overflow-hidden rounded-md">
      <SideBar />

      <Body />
    </div>
  );
};

export default Wrapper;
