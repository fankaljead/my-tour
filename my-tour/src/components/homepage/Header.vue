<template>
  <div class="travel_note_brief_header">
    <van-tabs animated v-model="active" swipeable>
      <van-tab title="我的地点">
        <TravelNoteBrief
          v-for="travel_note in travel_note_infos.travelNoteInfos"
          v-bind:data="travel_note"
          :key="travel_note.id"
        />
      </van-tab>
      <van-tab title="我的路线">
        <TravelMap />
        <!-- <template>
             <baidu-map
             class="map"
             :center="{ lng: 116.404, lat: 39.915 }"
             :zoom="15"
             :scroll-wheel-zoom="true"
             >
             <bm-polyline
             :path="polylinePath"
             stroke-color="blue"
             :stroke-opacity="0.5"
             :stroke-weight="2"
             :editing="true"
             @lineupdate="updatePolylinePath"
             ></bm-polyline>
             </baidu-map>
             </template> -->
      </van-tab>
      <van-tab title="发现">
        <TravelNoteBrief
          v-for="travel_note in all_travel_note_infos.travelNoteInfos"
          v-bind:data="travel_note"
          :key="travel_note.id"
        />
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import Vue from "vue";
import { Tab, Tabs, Search } from "vant";
// import { BmPolyline } from "vue-baidu-map";

import TravelMap from './TravelMap'
import { GET_ALL_TRAVEL_NOTES } from "../../store/mutation-types.js";
import { request } from "../../nework/request.js";

import TravelNoteBrief from "./TravelNoteBrief";

Vue.use(Search);
Vue.use(Tab);
Vue.use(Tabs);

export default {
  name: "Header",
  components: {
    TravelNoteBrief,
    // BmPolyline,
    TravelMap
  },
  mounted() {
    this.$store.dispatch(GET_ALL_TRAVEL_NOTES);
    this.travel_note_infos = this.$store.state.travel_note.all_travel_note_infos;
    request({
      url: "travel_note/getAllUserTravelNoteInfo",
      method: "get",
    }).then((res) => {
      this.all_travel_note_infos = res.data;
    });
  },
  data() {
    return {
      travel_note_infos: {},
      all_travel_note_infos: {},
      polylinePath: [
        { lng: 116.404, lat: 39.915 },
        { lng: 116.405, lat: 39.92 },
        { lng: 116.423493, lat: 39.907445 },
      ],
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
    updatePolylinePath(e) {
      this.polylinePath = e.target.getPath();
    },
    addPolylinePoint() {
      this.polylinePath.push({ lng: 116.404, lat: 39.915 });
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
