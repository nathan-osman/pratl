import { useDispatch, useSelector } from 'react-redux';
import { login, setPassword, setUsername } from './authSlice';
import './Auth.scss';

const Auth = () => {

  const auth = useSelector(state => state.auth);
  const dispatch = useDispatch();

  // TODO: use submit button to enable "return" for login

  return <div id="auth">
    <div className="panel">
      <div className="title">Pratl</div>
      <div className="content">
        <div>
          <input
            type="text"
            placeholder="username"
            value={auth.username}
            onChange={(e) => dispatch(setUsername(e.target.value))}
            disabled={auth.isAuthenticating}
            autoFocus
          />
        </div>
        <div>
          <input
            type="password"
            placeholder="password"
            value={auth.password}
            onChange={(e) => dispatch(setPassword(e.target.value))}
            disabled={auth.isAuthenticating}
          />
        </div>
      </div>
      <button
        type="button"
        onClick={() => dispatch(login({ username: auth.username, password: auth.password }))}
        disabled={auth.isAuthenticating}
      >
        {auth.isAuthenticating ? "Please wait..." : "Login"}
      </button>
    </div>
  </div>;
};

export default Auth;
