import { useSelector } from 'react-redux';
import './Messages.scss';

const Messages = () => {

  const messages = useSelector(state => state.messages);

  return <div id="messages">
    {messages.all.map(m =>
      <div className="message left">
        <img className="avatar" src="https://via.placeholder.com/96" alt="User avatar" />
        <div className="content">{m.body}</div>
      </div>
    )}
  </div>;
};

export default Messages;
