<template>
  <div class="travel_note_brief">
    <van-divider
      :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 16px' }"
    />

    <van-panel
      v-bind:icon="img_src"
      v-bind:title="data.title"
      @click="travelNoteBriefClick(data.id, data.author_name)"
      v-bind:desc="data.content.substring(0, 30) + '...'"
    >
      <div>{{ data.author_name }} | {{ data.create_time }}</div>
    </van-panel>

    <!-- <van-divider
         :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 16px' }"
         /> -->
  </div>
</template>

<script>
import Vue from "vue";
import { Panel } from "vant";
import { Divider } from "vant";
import { request } from "../../nework/request.js";
import { BASE_STATIC_IMAGE_URL } from "../../uitl/consts.js";

Vue.use(Divider);
Vue.use(Panel);

export default {
  name: "TravelNoteBrief",
  data() {
    return {
      img_src: "",
    };
  },
  methods: {
    travelNoteBriefClick(travel_note_id, author_name) {
      console.log("==========travelNoteBriefClick");
      console.log(travel_note_id);
      this.$router.push({
        name: "checkATravelNote",
        params: { travel_note_id: travel_note_id, author_name: author_name },
      });
    },
  },
  mounted() {
    request({
      // url: "image/getImageById/"+this.data.cover,
      url: "image/" + this.data.cover,
      method: "get",
    }).then((res) => {
      console.log("image res====");
      console.log(res);
      this.img_src = BASE_STATIC_IMAGE_URL + res.data.source;
    });
  },
  props: {
    data: {
      author: String,
      title: String,
      intro: String,
      cover: String,
      time: String,
    },
  },
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
.travel_note_brief {
  margin-bottom: 40px;
}
</style>
