// image module

import {
    fileRequest
} from '../../nework/request.js'

import {
    UPLOAD_A_IMAGE
} from '../mutation-types.js'

const state = () => ({
    upload_image_ids: ''
})

// getters
const getters = {}

// actions
const actions = {
    [UPLOAD_A_IMAGE](context, payload) {
        console.log(payload)
        const file = payload.file
        fileRequest(file)
            .then((response) => {
                return response.text();
            })
            .then((res) => {
                console.log("result:", res);
                file.status = 'done';
                file.message = '上传成功';
                context.commit(UPLOAD_A_IMAGE, {
                    id: res,
                    call() {
                        payload.call()
                    }
                })
            })
            .catch((error) => console.log("error", error));
    },
}

// mutations
const mutations = {
    [UPLOAD_A_IMAGE](state, payload) {
        state.upload_image_ids = state.upload_image_id == "" ? payload.id : state.upload_image_ids + ":" + payload.id
        payload.call()
        console.log("mutations payload: ", payload);
    },
}

export default {
    // namespaced: true,
    state,
    getters,
    actions,
    mutations
}
