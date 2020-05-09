<template>
  <div>
    <!-- <h1>{{ image.title }}</h1> -->

    <van-swipe :autoplay="3000">
      <van-swipe-item v-for="(image, index) in images" :key="index">
        <img width="100%" height="auto" v-lazy="image" />
      </van-swipe-item>
    </van-swipe>
  </div>
</template>

<script>
import Vue from "vue";
import { Swipe, SwipeItem } from "vant";
import { Lazyload } from "vant";

Vue.use(Lazyload);
Vue.use(Swipe);
Vue.use(SwipeItem);

export default {
  name: "ShowATravelNote",
  props: {
    image: {
      title: String,
      content: String,
      id: String,
      image_ids: String,
      create_time: String,
    },
  },
  data() {
    return {
      images: [
        "https://img.yzcdn.cn/vant/apple-1.jpg",
        "https://img.yzcdn.cn/vant/apple-2.jpg",
      ],
    };
  },
};
</script>

<style></style>
