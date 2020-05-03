<template>
  <div class="travel_note_brief">
    <van-panel
      v-bind:icon="data.img"
      v-bind:title="data.title"
      v-bind:desc="data.intro"
    >
      <div>{{data.author}} | {{ data.time }}</div>
    </van-panel>
  </div>
</template>

<script>
import Vue from "vue";
import { Panel } from "vant";

Vue.use(Panel);

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
/* .travel_note_left {
    width: 20%;
    background-color: #12f32f;
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
    } */
.travel_note_brief .van-icon__image {
  width: 5em;
  height: 5em;
  object-fit: contain;
}

.travel_note_brief .van-cell {
  color: #323233;
  font-size: 18px;
  line-height: 24px;
}
.travel_note_brief .van-panel__content {
  float: right;
  margin-right: 30px;
}
</style>
