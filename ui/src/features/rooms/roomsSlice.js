import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  all: [
    {
      name: "The First Room",
      creationDate: "45 minutes"
    },
    {
      name: "Room for Bob and Jessica",
      creationDate: "1 day"
    },
    {
      name: "Test Room",
      creationDate: "2 weeks"
    }
  ],
  activeIndex: -1
};

const roomsSlice = createSlice({
  name: 'rooms',
  initialState: initialState,
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
