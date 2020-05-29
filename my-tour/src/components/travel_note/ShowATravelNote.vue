<template>
  <div>
    <Header title="旅行日志" />
    <van-panel
      :title="data.title"
      :desc="this.author_name + ' / ' + data.create_time"
    >
      <p @click="checkUserProfile(data.author_id)">
        {{ author_name + " / " + data.create_time }}
      </p>
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
      <div>{{ data.content }}</div>
    </van-panel>

    <TravelNoteComment :travel_note_id="data.travel_note_id" />
  </div>
</template>

<script>
import Vue from "vue";
import { Swipe, SwipeItem } from "vant";
import { Lazyload } from "vant";

import Header from "../../components/Header";
// import { BASE_STATIC_IMAGE_URL } from "../../uitl/consts.js";
import TravelNoteComment from "../../components/comment/TravelNoteComment";
import { Panel } from "vant";
import { Divider } from "vant";

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
    author_name: String,
  },
  components: {
    Header,
    TravelNoteComment,
  },
  mounted() {
    console.log("ddddd======ddddd");
    console.log(this.data);
  },
  data() {
    return {};
  },
  methods: {
    checkUserProfile(id) {
      console.log("ccccccliiiick === id:", id);
      this.$router.push({
        name: "personProfile",
        params: { user_id: id },
      });
    },
  },
};
</script>

<style></style>
