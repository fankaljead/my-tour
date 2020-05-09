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
              url:
                'https://cdn0.iconfinder.com/data/icons/flat-round-system/512/archlinux-512.png',
              size: { width: 50, height: 50 },
            }"
          ></bm-marker>
        </baidu-map>
      </van-tab>
      <van-tab title="推荐">
        <TravelNoteBrief
          v-for="travel_note in travel_note_infos.travelNoteInfos"
          v-bind:data="travel_note"
          :key="travel_note.id"
        />

      </van-tab>
      <van-tab title="热榜">
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import Vue from "vue";
import { Tab, Tabs, Search } from "vant";
import { BmMarker } from "vue-baidu-map";

import { GET_ALL_TRAVEL_NOTES } from "../../store/mutation-types.js";

import TravelNoteBrief from "./TravelNoteBrief";

Vue.use(Search);
Vue.use(Tab);
Vue.use(Tabs);

export default {
  name: "Header",
  components: {
    TravelNoteBrief,
    BmMarker,
  },
  mounted() {
    this.$store.dispatch(GET_ALL_TRAVEL_NOTES);
    this.travel_note_infos = this.$store.state.travel_note.all_travel_note_infos;
  },
  data() {
    return {
      travel_note_infos: {},
      markerPoint: { lng: 116.404, lat: 39.915 },
      active: 2,
      data: {
        author_name: "游乐王子",
        title: "摸仙煲一日游",
        content: "摸仙煲一日游摸仙煲一日游摸仙煲一日游摸仙煲一日游摸仙煲一日游",
        cover: "73",
          // "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.5rYEZ-xCStnRSPYhA3_i_QAAAA%26pid%3DApi&f=1",
        create_time: "2020-05-02",
      },
    };
  },
  methods: {
    update(e) {
      this.points = e.target.cornerPoints;
    },
  },
};
</script>

<style>
.map {
  width: 100%;
  height: 400px;
}
</style>
