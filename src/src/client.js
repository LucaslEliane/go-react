import React from 'react';
import { render, hydrate } from 'react-dom';
import { createStore } from 'redux';
import { Provider } from 'react-redux';
import Counter from './pages/Counter';
import reducer from './store/reducer';

function load(state, isPreload) {
    const store = createStore(reducer, state);

    (isPreload ? hydrate : render)(
        <Provider store={store}>
            <Counter />
        </Provider>,
        document.getElementById('root'),
    );
}

load(window.__PRELOADED_STATE__, !!window.__PRELOADED_STATE__);