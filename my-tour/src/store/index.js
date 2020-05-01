import Vue from 'vue'
import Vuex from 'vuex'


import {
    request
} from '../nework/request.js'

import {
    LOGIN
} from './mutation-types.js'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        name: '',
        uid: 0,
    },
    mutations: {
        [LOGIN](state, payload) {
            state.name = payload.name
            state.uid = payload.uid
        }
    },
    actions: {
        [LOGIN](context, payload) {
            request({
                url: '/user/login',
                username: payload.username,
                password: payload.username
            }).then(res => {
                console.log("res: ", res);
            })
            context.commit(LOGIN, payload)
        }
    },
    modules: {}
})
