import axios from 'axios'

export function request(config) {
    const instance = axios.create({
        baseURL: 'http://127.0.0.1:8080/v1/',
        timeout: 300,
        // headers: {
        //     'Access-Control-Allow-Origin': '*',
        //     'Content-Type': 'application/json',
        //     'Access-Control-Allow-Methods': 'GET, PUT, POST, DELETE, OPTIONS',
        // },

        // withCredentials: true 
    })

    return instance(config)
}
