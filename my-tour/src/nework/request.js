import axios from 'axios'

export function request(config) {
    const instance = axios.create({
        baseURL: 'https://localhost:8080/v1/',
        timeout: 300,
        withCredentials: true 
    })

    return instance(config)
}
