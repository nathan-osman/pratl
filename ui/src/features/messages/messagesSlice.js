import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  all: [
    {
      member: {
        user: {
          username: "Bob"
        },
        is_owner: false,
        is_admin: false
      },
      body: "This is a test message.",
      creation_date: "2021-10-22T18:25:43-05:00",
      star_count: 1,
      is_edited: false
    },
    {
      member: {
        user: {
          username: "Sam"
        }
      },
      body: "This is a reply.",
      creation_date: "2021-10-22T18:25:43-05:00",
      star_count: 0,
      is_edited: false
    }
  ]
};

const messagesSlice = createSlice({
  name: 'messages',
  initialState,
  reducers: {
    setMessages(state, action) {
      state.all = action.payload;
    }
  }
});

export const { setMessages } = messagesSlice.actions;

export default messagesSlice.reducer;
