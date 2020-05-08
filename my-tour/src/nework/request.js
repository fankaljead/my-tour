import axios from 'axios'

export function request(config) {
    const instance = axios.create({
        baseURL: 'http://127.0.0.1:8080/v1/',
        timeout: 3000,
        // headers: {
        //     'Access-Control-Allow-Origin': '*',
        //     'Content-Type': 'application/json',
        //     'Access-Control-Allow-Methods': 'GET, PUT, POST, DELETE, OPTIONS',
        // },

        // withCredentials: true 
    })

    return instance(config)
}


export function fileRequest(file) {
    // const instance = axios.create({
    //     baseURL: 'http://127.0.0.1:8080/v1/',
    //     timeout: 3000,
    //     headers: {
    //         // 'Access-Control-Allow-Origin': '*',
    //         // 'Access-Control-Allow-Methods': 'GET, PUT, POST, DELETE, OPTIONS',
    //         "Content-Type": "multipart/form-data",
    //     },

    //     withCredentials: true
    // })
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "multipart/form-data");

    var formdata = new FormData();
    formdata.append("image_name", file.file);

    var requestOptions = {
        method: 'POST',
        // headers: myHeaders,
        body: formdata,
        redirect: 'follow'
    };

    return fetch("http://localhost:8080/v1/image", requestOptions);
}
