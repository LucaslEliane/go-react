import React from 'react';
import { connect } from 'react-redux';
import logo from './logo.svg';
import { decrease, increase } from "../store/action/todo";

class Counter extends React.Component {
    render () {
        const { count, handleIncrement, handleDecrement } = this.props;

        return (
            <div className="App">
                <header className="App-header">
                    <img src={logo} className="App-logo" alt="logo" />
                    <p>
                        Edit <code>src/App.js</code> and save to reload.
                    </p>
                    <a className="App-link">
                        { '值为: ' + count }
                    </a>
                    <a className="App-link" onClick={handleIncrement}>
                        Increment
                    </a>
                    <a className="App-link" onClick={handleDecrement}>
                        Decrement
                    </a>
                </header>
            </div>
        );
    }
}

export default connect(
    state => {
        const { todo } = state;
        return {
            count: todo.count,
        };
    },
    dispatch => {
        return {
            handleIncrement: () => increase(dispatch),
            handleDecrement: () => decrease(dispatch),
        };
    },
)(Counter);