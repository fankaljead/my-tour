import Vue from 'vue'
import App from './App.vue'
// import BaiduMap from 'vue-baidu-map'
import VueAMap from 'vue-amap';

import router from './router'
import store from './store'

Vue.config.productionTip = false
// 引入vue-amap

// Vue.use(BaiduMap, {
//     // ak 是在百度地图开发者平台申请的密钥 详见 http://lbsyun.baidu.com/apiconsole/key */
//     ak: 'dcmnpfSSRrDlnyyWxPgBVVLbszBKPBE5',
// })


Vue.use(VueAMap);
VueAMap.initAMapApiLoader({
    key: 'dc39939bc7d412a77023df8c85f27a92',
    plugin: ['AMap.Autocomplete', 'AMap.PlaceSearch', 'AMap.Scale', 'AMap.OverView', 'AMap.ToolBar', 'AMap.MapType', 'AMap.PolyEditor', 'AMap.CircleEditor'],
    v: '1.4.4'
});

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')
