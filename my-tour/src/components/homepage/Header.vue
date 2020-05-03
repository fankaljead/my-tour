<template>
  <div class="travel_note_brief_header">
    <van-tabs animated v-model="active" swipeable>
      <van-tab title="我的">
        <baidu-map
          class="map"
          :center="{ lng: 116.404, lat: 39.915 }"
          :zoom="15"
        >
          <bm-marker
            :position="markerPoint"
            :dragging="true"
            animation="BMAP_ANIMATION_BOUNCE"
            :icon="{
              url: 'https://cdn0.iconfinder.com/data/icons/flat-round-system/512/archlinux-512.png',
              size: { width: 50, height: 50 }
            }"
          ></bm-marker>
        </baidu-map>
      </van-tab>
      <van-tab title="推荐">
        <TravelNoteBrief v-bind:data="data" />
        <TravelNoteBrief v-bind:data="data" />
        <TravelNoteBrief v-bind:data="data" />
        <TravelNoteBrief v-bind:data="data" />
        <TravelNoteBrief v-bind:data="data" />
        <TravelNoteBrief v-bind:data="data" />
        <TravelNoteBrief v-bind:data="data" />
      </van-tab>
      <van-tab title="热榜">
        <TravelNoteBrief v-bind:data="data" />
        <TravelNoteBrief v-bind:data="data" />
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import Vue from "vue";
import { Tab, Tabs, Search } from "vant";
import { BmMarker } from "vue-baidu-map";

import TravelNoteBrief from "./TravelNoteBrief";

Vue.use(Search);
Vue.use(Tab);
Vue.use(Tabs);

export default {
  name: "Header",
  components: {
    TravelNoteBrief,
    BmMarker
  },
  data() {
    return {
      markerPoint: { lng: 116.404, lat: 39.915 },
      active: 2,
      data: {
        author: "游乐王子",
        title: "摸仙煲一日游",
        intro: "摸仙煲一日游摸仙煲一日游摸仙煲一日游摸仙煲一日游摸仙煲一日游",
        img:
          "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.5rYEZ-xCStnRSPYhA3_i_QAAAA%26pid%3DApi&f=1",
        time: "2020-05-02"
      }
    };
  },
  methods: {
    update(e) {
      this.points = e.target.cornerPoints;
    }
  }
};
</script>

<style>
.map {
  width: 100%;
  height: 400px;
}
</style>
