import axios from "axios";

const BASE_URL = process.env.REACT_APP_BASE_URL;

interface Result {
  status: string;
  index: string;
  value: string;
}

export async function findIndex(value: string): Promise<Result | null> {
  try {
    const response = await axios.get(`${BASE_URL}/endpoint/${value}`);
    const data = await response.data as Result;
    console.debug(response.status, data)
    return { ...data, status: response.status.toString() };
  } catch (err) {
    console.error(err)
    return null;
  }
}
