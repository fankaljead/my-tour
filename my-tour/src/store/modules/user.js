// user module

import {
    request
} from '../../nework/request.js'

import {
    LOGIN,
    REGISTER,
    UPDATE_PERSON_PROFILE
} from '../mutation-types.js'

const state = () => ({
    username: '',
    uid: 0,
    profile_id: 0,
})

// getters
const getters = {}

// actions
const actions = {
    [UPDATE_PERSON_PROFILE](context, payload) {
        request({
            url: "user_info",
            method: 'put',
            params: {
                tel: payload.tel,
                qq: payload.qq,
                email: payload.email,
                wechat: payload.wechat,
                location: payload.location,
                birthday: payload.birthday,
                personalized_signature:payload.personalized_signature,
            }
        }).then(res => {
            console.log("update person profile: ", res)
            context.commit(UPDATE_PERSON_PROFILE, res.data)
            payload.call()
        })
    },
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
    },

    [REGISTER](context, payload) {
        console.log("action register")
        console.log(payload)
        request({
            url: 'user',
            method: 'post',
            params: {
                username: payload.username,
                password: payload.password
            }
        }).then(res => {
            console.log("actions res: ", res.data);
            context.commit(REGISTER, {
                username: payload.username,
                uid: res.data.uid
            })
            payload.call()
        })
    },
}

// mutations
const mutations = {
    [LOGIN](state, payload) {
        state.username = payload.username
        state.uid = payload.uid
        console.log("mutations payload: ", payload);
    },
    [REGISTER](state, payload) {
        state.username = payload.username
        state.uid = payload.uid
    },
    [UPDATE_PERSON_PROFILE](state, payload) {
        state.profile_id = payload
    }
}

export default {
    // namespaced: true,
    state,
    getters,
    actions,
    mutations
}
