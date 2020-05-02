<template>
  <div>
    <div class="travel_note_left">
      <!-- <img alt="" v-bind:src="data.img" /> -->
      <van-image
        width="15rem"
        height="15rem"
        fit="contain"
        v-bind:src="data.img"
      >
      </van-image>
    </div>
    <div class="travel_note_right">
      <div>
        <h2>{{ data.title }}</h2>
        <p>
          {{ data.intro }}
        </p>
        <p>
          <span>{{ data.author }}</span>
          <span>
            {{ data.time }}
          </span>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue';
import { Image as VanImage } from 'vant';
import { Card } from 'vant';

Vue.use(Card);

Vue.use(VanImage);

export default {
  name: "TravelNoteBrief",
  props: {
    data: {
      author: String,
      title: String,
      intro: String,
      img: String,
      time: String
    }
  }
};
</script>

<style>
.travel_note_left {
  width: 20%;
  /* background-color: #12f32f; */
  height: 100px;
  float: left;
}
.travel_note_left img {
  width: 100px;
  height: 100px;
}
.travel_note_right {
  width: 80%;
  background-color: #fff321;
  height: 100px;
  float: right;
}
</style>
