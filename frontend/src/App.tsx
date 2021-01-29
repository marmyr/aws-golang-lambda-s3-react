import React from 'react';
import './App.css';
const getenv = require('getenv');

const url = getenv.string('REACT_APP_API');
class App extends React.Component {
  state = {message: '', name: ''};

  callApi() {
    fetch(`${url}/hello?name=${encodeURIComponent(this.state.name)}`, {
        method: 'GET',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        }
      }
    )
      .then((response) => response.json())
      .then((response) => {
        this.setState({message: response.msg})
      });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <label htmlFor="name">Name:</label>
          <input type="input" placeholder="Your name (optional)" id="name" value={this.state.name} onChange={(e) => this.setState({name: e.target.value})} />
          <br/>
          <input type="button" onClick={() => this.callApi()} value="Call API"/>
          <p>
            {this.state.message}
          </p>
        </header>
      </div>
    );
  }
}

export default App;
