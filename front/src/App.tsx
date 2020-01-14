import React, { useState } from "react";
import "./App.css";
import { FizzBuzzForm } from "./FizzBuzzForm";
import { FizzBuzzList } from "./FizzBuzzList";

export interface FizzBuzzParams {
  string1: string;
  string2: string;
  int1: number;
  int2: number;
  limit: number;
}

const App: React.FC = () => {
  const [output, setOutput] = useState([]);

  const api = (params: FizzBuzzParams) => {
    fetch(
      `https://lbcfb.antoineviau.com/api?string1=${params.string1}&string2=${params.string2}&int1=${params.int1}&int2=${params.int2}&limit=${params.limit}`
    )
      .then(response => response.json())
      .then(data => {
        setOutput(data);
      })
      .catch(e => console.log(e));
  };

  return (
    <div className="App">
      <h1>LeBonCoin FizzBuzz REST API</h1>
      <FizzBuzzForm
        string1="fizz"
        string2="buzz"
        int1={3}
        int2={5}
        limit={20}
        api={api}
      />
      <FizzBuzzList output={output} />
    </div>
  );
};

export default App;
