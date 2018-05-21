import { fetchRecommends } from '../services/api';

export const M_SET = 'set';
export const M_START_LOAD = 'startLoad';
export const A_FETCH_POIS = 'fetchPois';

export const state = () => ({
  list: [],
  loading: false,
});

export const mutations = {
  [M_SET](state, list) {
    state.list = list;
    state.loading = false;
  },
  [M_START_LOAD](state) {
    state.list = [];
    state.loading = true;
  },
};

export const actions = {
  [A_FETCH_POIS]({ commit, rootState }) {
    commit(M_START_LOAD);
    fetchRecommends(rootState.condition).then(res => {
      // console.log(JSON.stringify(res, null, 2));
      commit(M_SET, res.recommends || []);
    });
  },
};
