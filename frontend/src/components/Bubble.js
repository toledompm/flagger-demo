import React from 'react';

const Bubble = ({ color }) => {
  console.log(color);
  const style = {
    width: '50px',
    height: '50px',
    borderRadius: '50%',
    backgroundColor: color,
  };

  return <div style={style}></div>;
};

export default Bubble;