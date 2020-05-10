<template>
  <div>
    <Header title="旅行日志" />
    <van-panel :title="data.title" :desc="this.author_name+' / '+data.create_time">
      <van-swipe :autoplay="3000">
        <van-swipe-item v-for="(image, index) in images" :key="index">
          <img
            width="100%"
            height="auto"
            v-lazy="'http://localhost:8080/public/' + image.source"
          />
        </van-swipe-item>
      </van-swipe>
      <van-divider />
      <div>{{data.content}}</div>
    </van-panel>
  </div>
</template>

<script>
import Vue from "vue";
import { Swipe, SwipeItem } from "vant";
import { Lazyload } from "vant";

import Header from "../../components/Header";
// import { BASE_STATIC_IMAGE_URL } from "../../uitl/consts.js";
import { Panel } from 'vant';
import { Divider } from 'vant';

Vue.use(Divider);
Vue.use(Panel);

Vue.use(Lazyload);
Vue.use(Swipe);
Vue.use(SwipeItem);

export default {
  // helll worl
  name: "ShowATravelNote",
  props: {
    images: Array,
    data: {},
    author_name: String
  },
  components: {
    Header,
  },
  mounted() {
    console.log("ddddd======ddddd");
    console.log(this.data);
  },
  data() {
    return {};
  },
};
</script>

<style></style>
