import { useState } from 'react';
import { findIndex } from '../api';

type InputEvent = React.ChangeEvent<HTMLInputElement>;

export function Search() {
  const [number, setNumber] = useState('');
  const [index, setIndex] = useState('');
  const [value, setValue] = useState('');
  const [status, setStatus] = useState('');

  const numberChange = (event: InputEvent) => {
    setNumber(event?.target?.value ?? '');
  };

  const searchClick = async () => {
    if (!number) {
      return
    }
    const response = await findIndex(number);
    setStatus(response?.status ?? '');
    setIndex(response?.index ?? '');
    setValue(response?.value ?? '');
  };

  return (
    <div className="App">
      <p>
        <input type="text" placeholder="type number to find" onChange={numberChange} value={number} />
        <button type="button" onClick={searchClick}>Search</button>
      </p>
      <div>
        <span>Result:</span>
        <p>
          <label>api status:</label>
          <span>{status}</span>
        </p>
        <p>
          <label>Index found:</label>
          <span>{index}</span>
        </p>
        <p>
          <label>value found:</label>
          <span>{value}</span>
        </p>
      </div>
    </div>
  );
}
