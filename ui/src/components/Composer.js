import './Composer.scss';

const Composer = () => {
  return <div id="composer">
    <div className="controls">
      <input
        type="text"
        className="elem input"
        autocomplete="off"
      />
      <button
        type="button"
        className="elem btn"
      >Send</button>
    </div>
  </div>;
};

export default Composer;
