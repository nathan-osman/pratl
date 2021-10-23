import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  isAuthenticated: true
};

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    logout(state, action) {
      state.isAuthenticated = false;
    }
  }
});

export const { logout } = authSlice.actions;

export default authSlice.reducer;
