<template>
  <div>
    <div class="amap-page-container">
      <el-amap
        vid="amap"
        :zoom="zoom"
        :center="center"
        class="amap-travel-note"
      >
        <el-amap-polyline
          :editable="polyline.editable"
          :path="polyline.path"
          :events="polyline.events"
        ></el-amap-polyline>
      </el-amap>
    </div>
    <AllTravelRoutine
      v-bind:routines="this.$store.state.travel_note.routines"
      :call="call"
    />
  </div>
</template>

<script>
import AllTravelRoutine from "../travel_note/AllTravelRoutine";
import { GET_TRAVEL_ROUTINES } from "../../store/mutation-types.js";
import { request } from "../../nework/request.js";

export default {
  name: "TravelMap",
  components: {
    AllTravelRoutine,
  },
  mounted() {
    this.$store.dispatch(GET_TRAVEL_ROUTINES);
  },
  data() {
    return {
      zoom: 12,
      travel_notes_infos: {},
      center: [121.5273285, 31.25515044],
      polyline: {
        path: [
          [121.5389385, 31.21515044],
          [121.5389385, 31.29615044],
          [121.5273285, 31.21515044],
        ],
        events: {
          click(e) {
            alert("click polyline: ", e);
          },
          end: (e) => {
            let newPath = e.target
              .getPath()
              .map((point) => [point.lng, point.lat]);
            console.log(newPath);
          },
        },
        editable: true,
      },
    };
  },
  methods: {
    changeEditable() {
      this.polyline.editable = !this.polyline.editable;
    },
    call(id) {
      request({
        url: "travel_note/getTravelNotesInfoByRoutineId/" + id,
        method: "get",
      }).then((res) => {
        console.log("res all info:", res.data);
        this.travel_notes_info = res.data;
        // this.polyline.path = [];
        // this.center = [
        //   this.travel_notes_info.notes[0].latitude,
        //   this.travel_notes_info.notes[0].longitude,
        // ];
        var path = [];
        // for (let i = 0; i < this.travel_notes_info.notes.length; ++i) {
        //   let t = this.travel_notes_info.notes[i];
        //   let temp = [];
        //   temp.push(t.latitude);
        //   temp.push(t.longitude);
        //   path.push(temp);
        // }
        this.travel_notes_info.notes.forEach((value) => {
          console.log("value:::::", value);
          path.splice(path.length,0,[value.latitude, value.longitude])
          // path.push([value.latitude, value.longitude]);
          console.log("path:::::", path);
        });

        console.log("path:", path);
        this.$set(this.center, 0, path[0][0]);
        this.$set(this.center, 1, path[0][1]);
        this.$set(this.polyline, "path", path);

        console.log("polyline path: ", this.polyline.path);
        console.log("polyline center: ", this.center);
      });
    },
  },
};
</script>

<style>
.amap-map-travel-note {
  height: 300px;
}
</style>
