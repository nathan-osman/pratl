import axios from 'axios';

const instance = axios.create({
  baseURL: process.env.NODE_ENV === 'development' ? 'http://localhost' : ''
});

export const post = async function (url, data) {
  try {
    const response = await instance.post(url, data);
    return Promise.resolve(response.data);
  } catch (e) {
    let errorMessage = e.message;
    if (typeof e.response !== 'undefined' && 'error' in e.response.data) {
      errorMessage = e.response.data.error;
    }
    return Promise.reject(errorMessage);
  }
};

export default instance;
