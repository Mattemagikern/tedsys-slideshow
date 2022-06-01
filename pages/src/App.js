import logo from './logo.svg';
import './App.css';
import Upload from './components/upload.js';
import List from './components/list.js';

function App() {
  return (
    <div className="App">
	<div className="split list left">
	<div className="centered">
	  	<List />
	</div>
	</div>
	<div className="split upload right">
	<div className="centered">
		<Upload /> 
	</div>
	</div>
    </div>
  );
}

export default App;
