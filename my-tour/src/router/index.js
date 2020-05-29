import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login'
import ForgetPassword from '../views/ForgetPassword'
import Register from '../views/Register'
import HomePage from '../views/HomePage'
import Subscribe from '../views/Subscribe'
import Add from '../views/Add'
import Message from '../views/Message'
import PersonCenter from '../views/PersonCenter'
import AddTravelRoutine from '../views/travel_note/AddTravelRoutine'
import AddTravelNote from '../views/travel_note/AddTravelNote'
import CheckATravelNote from '../views/travel_note/CheckATravelNote'
// import AddPopMenu from '../components/add/AddPopMenu'
import PersonProfile from '../views/PersonProfile'



Vue.use(VueRouter)

const routes = [{
        path: '/',
        name: 'login',
        component: Login
    },
    {
        path: '/home',
        // name: 'Home',
        component: Home,
        children: [{
                path: '',
                component: HomePage
            },
            {
                path: '/subscribe',
                name: 'subscribe',
                component: Subscribe
            },
            {
                path: '/add',
                name: 'add',
                component: Add,
                children: [
                    // {
                    //     path: '',
                    //     name: 'add',
                    //     component: Add
                    // },
                    // {
                    //     path: '/addTravelRoutine',
                    //     name: 'addTravelRoutine',
                    //     component: AddTravelRoutine
                    // }
                ]
            },
            {
                path: '/addTravelRoutine',
                name: 'addTravelRoutine',
                component: AddTravelRoutine
            },
            {
                path: '/checkATravelNote',
                name: 'checkATravelNote',
                component: CheckATravelNote
            },
            {
                path: '/addTravelNote',
                name: 'addTravelNote',
                component: AddTravelNote
            },
            {
                path: '/message',
                name: 'message',
                component: Message
            },
            {
                path: '/personcenter',
                name: 'personcenter',
                component: PersonCenter
            },
            {
                path: '/personProfile',
                name: 'personProfile',
                component: PersonProfile
            },
        ]
    },
    {
        path: '/forget_password',
        name: 'forget_pasword',
        component: ForgetPassword
    },
    {
        path: '/register',
        name: 'register',
        component: Register
    },
    {
        path: '/about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import( /* webpackChunkName: "about" */ '../views/About.vue')
    },
    // 404
    {
        path: '*',
        component: () => import('@/components/404')
    }
]

const router = new VueRouter({
    routes,
    mode: 'history',
})

export default router
