import React, { useState } from "react";
import { FizzBuzzParams } from "./App";

export interface FizzBuzzFormProps extends FizzBuzzParams {
  api: (params: FizzBuzzParams) => void;
}

export const FizzBuzzForm: React.FC<FizzBuzzFormProps> = ({
  string1,
  string2,
  int1,
  int2,
  limit,
  api
}) => {
  const [params, setParams] = useState({ string1, string2, int1, int2, limit });

  const handleChange = (field: string, event: any) => {
    setParams({ ...params, [field]: event.target.value });
  };

  const handleSubmit = () => {
    api(params as FizzBuzzParams);
  };

  return (
    <form>
      <input
        className="form-control"
        type="text"
        name="string1"
        placeholder="String1"
        onChange={e => handleChange("string1", e)}
        value={params.string1}
        required
      />
      <input
        className="form-control"
        type="text"
        name="string2"
        placeholder="String2"
        onChange={e => handleChange("string2", e)}
        value={params.string2}
        required
      />
      <input
        className="form-control"
        type="number"
        min="0"
        name="int1"
        placeholder="Int1"
        onChange={e => handleChange("int1", e)}
        value={params.int1}
        required
      />
      <input
        className="form-control"
        type="number"
        min="0"
        name="int2"
        placeholder="Int2"
        onChange={e => handleChange("int2", e)}
        value={params.int2}
        required
      />
      <input
        className="form-control"
        type="number"
        min="0"
        name="limit"
        placeholder="Limit"
        onChange={e => handleChange("limit", e)}
        value={params.limit}
        required
      />
      <button
        className="btn btn-primary btn-lg"
        type="button"
        onClick={handleSubmit}
      >
        FizzBuzz
      </button>
    </form>
  );
};
