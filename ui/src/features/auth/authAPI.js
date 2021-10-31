import { post } from "../../common/axios";

export async function loginUser(username, password) {
  return await post('/auth/login', { username, password });
};
