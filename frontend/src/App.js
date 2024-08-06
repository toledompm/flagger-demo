import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Bubble from './components/Bubble';

const apiAddress = 'http://localhost:8000';

const allowedColors = ['green', 'blue', 'yellow', 'purple', 'orange', 'pink', 'brown', 'black', 'white'];

function App() {
  const [bubbles, setBubbles] = useState([]);

  useEffect(() => {
    const interval = setInterval(async () => {
      let bubbleColor = '';
      try {
        const response = await axios.get(`${apiAddress}/color`);
        if (response.status === 200 && allowedColors.includes(response.data.color)) {
          bubbleColor = response.data.color;
        } else {
          bubbleColor = 'red';
        }
      } catch (error) {
        console.error(error);
        bubbleColor = 'red';
      }

      setBubbles([...bubbles, bubbleColor]);
    }, 1000); // Sends request every 1 seconds

    return () => clearInterval(interval); // Cleanup on unmount
  }, [bubbles]);

  return (
    <div className="App" style={{ display: 'flex', flexWrap: 'wrap' }}>
      {bubbles.map((bubble, index) => (
        <React.Fragment key={index}>
          <Bubble color={bubble} />
          {(index + 1) % 10 === 0 && <br />}
        </React.Fragment>
      ))}
    </div>
  );
}

export default App;