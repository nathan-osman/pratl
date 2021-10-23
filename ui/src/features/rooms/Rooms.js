import classNames from 'classnames';
import { useDispatch, useSelector } from 'react-redux';
import ReactTimeAgo from 'react-time-ago';
import { setActive } from './roomsSlice';
import './Rooms.scss';

const Rooms = () => {

  const rooms = useSelector(state => state.rooms);
  const dispatch = useDispatch();

  return <div id="rooms">
    <div className="room-list">
      {rooms.all.map((r, i) =>
        <div
          className={classNames('item', { active: rooms.activeIndex === i })}
          onClick={() => dispatch(setActive(i))}
        >
          <div className="content">
            <div className="title">{r.name}</div>
            <div className="subtitle"><ReactTimeAgo date={r.creationDate} /></div>
          </div>
        </div>
      )}
    </div>
  </div>;
};

export default Rooms;
