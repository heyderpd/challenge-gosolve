import axios from "axios";

const BASE_URL = process.env.REACT_APP_BASE_URL;

interface Result {
  status: string;
  index: string;
  value: string;
}

const Api = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export async function findIndex(value: string): Promise<Result | null> {
  try {
    const response = await Api.get(`${BASE_URL}/endpoint/${value}`);
    const data = await response.data as Result;
    console.debug(response?.status, response)
    return { ...data, status: response.status.toString() };
  } catch (err) {
    console.error(err)
    return null;
  }
}
