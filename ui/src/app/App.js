import { useSelector } from 'react-redux';
import Composer from '../features/composer/Composer';
import Info from '../features/info/Info';
import Messages from '../features/messages/Messages';
import Rooms from '../features/rooms/Rooms';
import Auth from '../features/auth/Auth';
import './App.scss';

function App() {

  const auth = useSelector(state => state.auth);

  return (
    <div id="App">
      {auth.isAuthenticated ?
        <div className="page">
          <Rooms />
          <Messages />
          <Composer />
          <Info />
        </div> :
        <Auth />
      }
    </div>
  );
}

export default App;
