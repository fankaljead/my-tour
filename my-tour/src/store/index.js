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
        username: '',
        uid: 0,
    },
    mutations: {
        [LOGIN](state, payload) {
            state.username = payload.username
            state.uid = payload.uid
            console.log("mutations payload: ", payload);
        }
    },
    actions: {
        [LOGIN](context, payload) {
            console.log(payload)
            request({
                url: 'user/login',
                username: payload.username,
                password: payload.username
            }).then(res => {
                console.log("actions res: ", res.data);
                context.commit(LOGIN, {
                    username: payload.username,
                    uid: res.data
                })
                payload.call()
            })
        }
    },
    modules: {}
})
