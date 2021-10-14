import './App.scss';
import Composer from './components/Composer';
import Info from './components/Info';
import Messages from './components/Messages';
import Rooms from './components/Rooms';

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
