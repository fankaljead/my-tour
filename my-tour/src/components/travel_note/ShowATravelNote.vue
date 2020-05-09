<template>
  <div>
    <Header title="旅行日志" />

    <h1>{{}}</h1>

    <van-swipe :autoplay="3000">
      <van-swipe-item v-for="(image, index) in images" :key="index">
        <img
          width="100%"
          height="auto"
          v-lazy="'http://localhost:8080/public/' + image.source"
        />
      </van-swipe-item>
    </van-swipe>
  </div>
</template>

<script>
import Vue from "vue";
import { Swipe, SwipeItem } from "vant";
import { Lazyload } from "vant";

import Header from "../../components/Header";
// import { BASE_STATIC_IMAGE_URL } from "../../uitl/consts.js";

Vue.use(Lazyload);
Vue.use(Swipe);
Vue.use(SwipeItem);

export default {
  name: "ShowATravelNote",
  props: {
    images: Array,
    data: {},
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
