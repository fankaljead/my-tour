import {
    request
} from '../../nework/request.js'

import {
    ADD_TRAVEL_ROUTINE,
    GET_TRAVEL_ROUTINES,
    DELETE_TRAVEL_ROUTINE,
    ADD_A_TRAVEL_NOTE,
    GET_ALL_TRAVEL_NOTES
} from '../mutation-types.js'

const state = () => ({
    routines: [],
    add_travel_note_id: 0,
    all_travel_note_infos: {},
})

// getters
const getters = {}

// actions
const actions = {

    [ADD_A_TRAVEL_NOTE](context, payload) {
        console.log("add travel routine action: ", payload)
        request({
            url: 'travel_note',
            method: 'post',
            params: {
                title: payload.title,
                content: payload.content,
                publish_time: payload.publish_time,
                routine_id: payload.routine_id,
                image_ids: payload.image_ids,
                cover: payload.cover,
            }
        }).then(res => {
            console.log(res)
            if (res.data > 0) {
                context.commit(ADD_A_TRAVEL_NOTE, {
                    id: res.data
                })
                payload.call()
            }
        })
    },

    [ADD_TRAVEL_ROUTINE](context, payload) {
        console.log("add travel routine action: " + payload)

        request({
            url: 'travel_note/add_travel_routine',
            method: 'post',
            params: {
                title: payload.title,
                description: payload.description
            }
        }).then(res => {
            console.log(context)
            if (res.data > 0) {
                context.dispatch(GET_TRAVEL_ROUTINES)
            }
        })
    },
    [GET_TRAVEL_ROUTINES](context) {
        request({
            url: 'travel_note/get_travel_routines',
            method: 'get',
        }).then(res => {
            console.log("actions res: ", res);
            context.commit(GET_TRAVEL_ROUTINES, res.data)
        })
    },
    [DELETE_TRAVEL_ROUTINE](context) {
        request({
            url: 'travel_note/delete_travel_routine',
            method: 'get',
        }).then(res => {
            if (res.data > 0) {
                context.dispatch(GET_TRAVEL_ROUTINES)
            }
        })
    },
    [GET_ALL_TRAVEL_NOTES](context) {
        request({
            url: 'travel_note/getAllTravelNoteInfo',
            method: 'get'
        }).then(res => {
            console.log("===get all travel note info: ======")
            console.log(res)
            context.commit(GET_ALL_TRAVEL_NOTES, res.data)
        })
    }

}

// mutations
const mutations = {
    // [ADD_TRAVEL_ROUTINE](state, payload) {
    //     // state.routines = payload.routines
    //     console.log("mutations payload: ", state);
    //     console.log("mutations payload: ", payload);
    // },
    [GET_TRAVEL_ROUTINES](state, payload) {
        state.routines = payload.routines
    },
    [ADD_A_TRAVEL_NOTE](state, payload) {
        state.add_travel_note_id = payload.id
    },
    [GET_ALL_TRAVEL_NOTES](state, payload) {
        state.all_travel_note_infos = payload
    }
}

export default {
    // namespaced: true,
    state,
    getters,
    actions,
    mutations
}