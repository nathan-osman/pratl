import './Messages.scss';

const Messages = () => {
  return <div id="messages">
    <div className="message left">
      <img className="avatar" src="https://via.placeholder.com/96" alt="User avatar" />
      <div className="content">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec rhoncus semper dapibus. Quisque fermentum dignissim sagittis. Vestibulum pretium facilisis elit, at feugiat tortor dictum a. Phasellus nulla ante, consectetur quis efficitur nec, venenatis sed tellus. Nam et risus elit. Morbi semper lacus in massa lobortis viverra. Vestibulum vitae sollicitudin felis. Proin aliquet odio et eros interdum, congue hendrerit neque accumsan. Nullam ac faucibus eros. Fusce tincidunt lacus vel eleifend pretium.Sed id eros eu sem ultrices feugiat. Phasellus fermentum, sapien vitae consequat faucibus, diam sapien venenatis leo, eget efficitur lectus ipsum non risus. Proin tempor laoreet nisl vitae molestie. Proin vulputate metus vitae gravida mollis. Fusce lorem ligula, mollis eu elementum et, condimentum vel mi. In condimentum convallis suscipit. Aenean eu leo justo.
      </div>
    </div>
    <div className="message right">
      <img className="avatar" src="https://via.placeholder.com/96" alt="User avatar" />
      <div className="content">
        Hello, world!
      </div>
    </div>
    <div className="message left">
      <img className="avatar" src="https://via.placeholder.com/96" alt="User avatar" />
      <div className="content">
        A reply to the message!
      </div>
    </div>
  </div>;
};

export default Messages;
