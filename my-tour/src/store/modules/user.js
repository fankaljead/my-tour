// user module

import {
    request
} from '../../nework/request.js'

import {
    LOGIN
} from '../mutation-types.js'

const state = () => ({
    username: '',
    uid: 0,
})

// getters
const getters = {}

// actions
const actions = {
    [LOGIN](context, payload) {
        console.log(payload)
        request({
            url: 'user/login',
            params: {
                username: payload.username,
                password: payload.password
            }
        }).then(res => {
            console.log("actions res: ", res.data);
            context.commit(LOGIN, {
                username: payload.username,
                uid: res.data
            })
            payload.call()
        })
    }
}

// mutations
const mutations = {
    [LOGIN](state, payload) {
        state.username = payload.username
        state.uid = payload.uid
        console.log("mutations payload: ", payload);
    }
}

export default {
    // namespaced: true,
    state,
    getters,
    actions,
    mutations
}
