import { useDispatch } from 'react-redux';
import { login } from './authSlice';
import './Auth.scss';

const Auth = () => {

  const dispatch = useDispatch();

  return <div id="auth">
    <div className="panel">
      <div className="title">Pratl</div>
      <div className="content">
        <div><input type="text" autoFocus placeholder="username" /></div>
        <div><input type="password" placeholder="password" /></div>
      </div>
      <button
        type="button"
        onClick={() => dispatch(login())}
      >
        Login
      </button>
    </div>
  </div>;
};

export default Auth;
