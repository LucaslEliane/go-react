export const COUNTER_INCREASE = 'COUNTER_INCREASE';
export const COUNTER_DECREASE = 'COUNTER_DECREASE';

export const increase = dispatch => {
    dispatch(COUNTER_INCREASE);
};

export const decrease = dispatch => {
    dispatch(COUNTER_DECREASE);
};
