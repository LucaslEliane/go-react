import {COUNTER_DECREASE, COUNTER_INCREASE} from "../action/todo";

export default function todo(state = {
  count: 1,
}, action) {
  switch (action) {
    case COUNTER_INCREASE:
      return {
        count: state.count + 1,
      };
    case COUNTER_DECREASE:
      return {
        count: state.count - 1,
      };
    default:
      return state;
  }
};
