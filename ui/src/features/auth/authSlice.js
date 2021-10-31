import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { loginUser } from "./authAPI";

export const login = createAsyncThunk(
  'auth/login',
  async (arg) => {
    return await loginUser(arg.username, arg.password);
  }
);

const initialState = {
  isAuthenticating: false,
  isAuthenticated: false,
  username: '',
  password: ''
};

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    logout(state, action) {
      state.isAuthenticated = false;
      // TODO: invalidate local credentials
    },
    setUsername(state, action) {
      state.username = action.payload;
    },
    setPassword(state, action) {
      state.password = action.payload;
    }
  },
  extraReducers: (builder) => {
    builder.addCase(login.pending, (state, action) => {
      state.isAuthenticating = true;
    });
    builder.addCase(login.fulfilled, (state, action) => {
      //state.isAuthenticating = false;
      //state.isAuthenticated = true;
    });
    builder.addCase(login.rejected, (state, action) => {
      state.isAuthenticating = false;
      // TODO: error message
    });
  }
});

export const { logout, setUsername, setPassword } = authSlice.actions;

export default authSlice.reducer;
