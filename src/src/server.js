import React from 'react';
import { renderToString } from 'react-dom/server';
import createStore from 'redux';
import { Provider } from 'react-redux';
import Counter from './pages/Counter';
import reducer from './store/reducer';

export function render(preloadState = undefined) {
    const store = createStore(reducer, preloadState ? JSON.parse(preloadState) : undefined);

    return {
        html: renderToString(
            <Provider store={store}>
                <Counter />
            </Provider>
        ),
        state: JSON.stringify(store.getState())
    }
}