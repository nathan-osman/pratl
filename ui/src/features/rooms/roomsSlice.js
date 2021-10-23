import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  all: [
    {
      name: "The First Room",
      creationDate: "2021-10-22T18:25:43-05:00"
    },
    {
      name: "Room for Bob and Jessica",
      creationDate: "2021-10-21T18:25:43-05:00"
    },
    {
      name: "Test Room",
      creationDate: "2021-10-11T18:25:43-05:00"
    }
  ],
  activeIndex: -1
};

const roomsSlice = createSlice({
  name: 'rooms',
  initialState,
  reducers: {
    setRooms(state, action) {
      state.all = action.payload;
    },
    setActive(state, action) {
      state.activeIndex = action.payload;
    }
  }
});

export const { setRooms, setActive } = roomsSlice.actions;

export default roomsSlice.reducer;
