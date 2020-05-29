<template>
  <div>
    <van-nav-bar title="我的关注" />
    <TravelNoteBrief
      v-for="travel_note in travel_note_infos.travelNoteInfos"
      v-bind:data="travel_note"
      :key="travel_note.id"
    />
  </div>
</template>

<script>
import Vue from "vue";
import { NavBar } from "vant";
import { request } from "../nework/request.js";
import TravelNoteBrief from "../components/homepage/TravelNoteBrief.vue";

Vue.use(NavBar);
export default {
  name: "Subscribe",
  components: {
    TravelNoteBrief,
  },
  data() {
    return {
      travel_note_infos: {},
    };
  },
  mounted() {
    request({
      url: "/travel_note/getAllSubscribeTravelNoteInfo",
      method: "get",
    }).then((res) => {
      console.log("subscribe res: ", res);
      this.travel_note_infos = res.data;
    });
  },
};
</script>

<style></style>
