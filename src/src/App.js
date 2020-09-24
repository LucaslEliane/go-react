import React from 'react';
import { Provider, connect } from 'react-redux';
import { renderToString } from 'react-dom/server';
import store from './store';
import logo from './logo.svg';
import './App.css';

function App(props) {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          { props.count }
        </a>
      </header>
    </div>
  );
}

const mapStateToProps = state => {
  const { todo } = state;
  return {
    count: todo.count,
  };
};

const mapDispatchToProps = () => {
  return {};
};

const Target = connect(mapStateToProps, mapDispatchToProps)(App);

function Container() {
  return (
    <Provider store={store}>
      <Target />
    </Provider>
  );
}

export function render() {
  return {
    html: renderToString(
      <Provider store={store}>
        <Target />
      </Provider>
    ),
    state: JSON.stringify(store.getState()),
  };
}

export default Container;
