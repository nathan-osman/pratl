import { useDispatch } from 'react-redux';
import SettingsIcon from './settings.svg';
import LogoutIcon from './logout.svg';
import { logout } from '../auth/authSlice';
import './Info.scss';

const Info = () => {

  const dispatch = useDispatch();

  return <div id="info">
    <div className="content">
      TODO
    </div>
    <div className="actions">
      <div className="item">
        <img src={SettingsIcon} className="icon" alt="settings icon" />
        <div className="content">
          <div className="title">Settings</div>
        </div>
      </div>
      <div
        className="item"
        onClick={() => dispatch(logout())}
      >
        <img src={LogoutIcon} className="icon" alt="logout icon" />
        <div className="content">
          <div className="title">Logout</div>
        </div>
      </div>
    </div>
  </div>;
};

export default Info;
