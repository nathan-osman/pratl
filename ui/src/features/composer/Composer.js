import './Composer.scss';

const Composer = () => {
  return <div id="composer">
    <div className="controls">
      <input
        type="text"
        autoFocus
        autoComplete="off"
      />
      <button type="button">Send</button>
    </div>
  </div>;
};

export default Composer;
