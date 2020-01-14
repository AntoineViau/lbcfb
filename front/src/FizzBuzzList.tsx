import React from "react";

export interface FizzBuzzListProps {
  output: string[];
}

export const FizzBuzzList: React.FC<FizzBuzzListProps> = ({ output }) => {
  return (
    <ul className="items">
      {output && output.length
        ? output.map((str, index) => (
            <div key={index} className="item">
              {str}
            </div>
          ))
        : ""}
    </ul>
  );
};
