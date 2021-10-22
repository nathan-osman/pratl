import { configureStore } from '@reduxjs/toolkit';
import messagesReducer from '../features/messages/messagesSlice';
import roomsReducer from '../features/rooms/roomsSlice';

export const store = configureStore({
  reducer: {
    rooms: roomsReducer,
    messages: messagesReducer
  },
});
