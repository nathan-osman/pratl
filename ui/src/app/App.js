import './App.scss';
import Composer from '../features/composer/Composer';
import Info from '../features/info/Info';
import Messages from '../features/messages/Messages';
import Rooms from '../features/rooms/Rooms';

function App() {
  return (
    <div id="App">
      <Rooms />
      <Messages />
      <Composer />
      <Info />
    </div>
  );
}

export default App;
