import './Rooms.scss';

const Rooms = () => {
  return <div id="rooms">
    <div className="room-list">
      <div className="room">
        <div className="name">The First Room</div>
        <div className="time">5 minutes ago</div>
      </div>
      <div className="room active">
        <div className="name">Room for Bob and Jessica</div>
        <div className="time">1 day ago</div>
      </div>
    </div>
  </div>;
};

export default Rooms;
