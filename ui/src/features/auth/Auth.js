import { useDispatch, useSelector } from 'react-redux';
import { login, setPassword, setUsername } from './authSlice';
import './Auth.scss';

const Auth = () => {

  const auth = useSelector(state => state.auth);
  const dispatch = useDispatch();

  function handleSubmit(e) {
    e.preventDefault();
    dispatch(
      login({ username: auth.username, password: auth.password })
    );
  }

  return <div id="auth">
    <form
      className="panel"
      onSubmit={handleSubmit}
    >
      <div className="title">Pratl</div>
      {auth.errorMessage ?
        <div className="error">Error: {auth.errorMessage}</div> : null
      }
      <div className="content">
        <input
          type="text"
          placeholder="username"
          value={auth.username}
          onChange={(e) => dispatch(setUsername(e.target.value))}
          disabled={auth.isAuthenticating}
          autoFocus
        />
        <input
          type="password"
          placeholder="password"
          value={auth.password}
          onChange={(e) => dispatch(setPassword(e.target.value))}
          disabled={auth.isAuthenticating}
        />
      </div>
      <button
        type="submit"
        disabled={auth.isAuthenticating}
      >
        {auth.isAuthenticating ? "Please wait..." : "Login"}
      </button>
    </form>
  </div>;
};

export default Auth;
