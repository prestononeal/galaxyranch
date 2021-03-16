import logo from './logo.svg';
import './App.css';
import MomentsComponent from './MomentsComponent';

function App() {
  return (
    <div className="App">
      <header className="App-header">
      <img src={logo} className="App-logo" alt="logo" />
        <p>
          NFT Ranch
        </p>
        <div>
          NBA Top Shot moments:
          <MomentsComponent />
        </div>
      </header>
    </div>
  );
}

export default App;
