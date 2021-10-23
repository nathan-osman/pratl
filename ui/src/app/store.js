import { configureStore } from '@reduxjs/toolkit';
import authReducer from '../features/auth/authSlice';
import messagesReducer from '../features/messages/messagesSlice';
import roomsReducer from '../features/rooms/roomsSlice';

export const store = configureStore({
  reducer: {
    auth: authReducer,
    rooms: roomsReducer,
    messages: messagesReducer
  },
});
