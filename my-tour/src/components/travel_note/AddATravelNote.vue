<template>
  <div class="add_a_travel_note">
    <!-- <van-uploader v-model="fileList" multiple /> -->

    <Header title="新增地点" />

    <!-- ====== 选择Map start ====== -->
    <!-- <van-field
         readonly
         clickable
         label="选择Map"
         :value="value"
         placeholder="选择Map"
         @click="showMapPicker = true"
         /> -->
    <!-- <van-popup v-model="showMapPicker" round position="bottom">
         <template>
         <baidu-map
         @touchstart="touchStart"
         @dragstart="dragStart"
         @touchmove="touchMove"
         @touchend="touchEnd"
         @moving="syncCenterAndZoom"
         @moveend="syncCenterAndZoom"
         @zoomend="syncCenterAndZoom"
         :scroll-wheel-zoom="true"
         class="map"
         style="height: 500px;"
         :center="center"
         >
         <bm-city-list anchor="BMAP_ANCHOR_TOP_LEFT"></bm-city-list>
         <bm-geolocation
         anchor="BMAP_ANCHOR_BOTTOM_RIGHT"
         :showAddressBar="showAddressBar"
         :autoLocation="autoLocation"
         ></bm-geolocation>
         </baidu-map>
         </template>
         </van-popup> -->
    <!-- ====== 选择Map end ====== -->

    <van-form @submit="onFormSubmit">
      <!-- ===== 输入标题 start ===== -->
      <van-field
        v-model="title"
        name="title"
        label="标题"
        placeholder="请输入标题"
        :rules="[{ required: true, message: '请输入标题' }]"
      />
      <!--===== 输入标题 end =====-->

      <!-- ====== 选择时间 start ====== -->
      <van-field
        readonly
        clickable
        label="选择时间"
        :value="time"
        placeholder="选择时间"
        @click="showTimePicker = true"
      />
      <van-popup v-model="showTimePicker" round position="bottom">
        <van-datetime-picker
          v-model="currentDate"
          type="datetime"
          title="选择时间"
          :min-date="minDate"
          :max-date="maxDate"
          visible-item-count="5"
          @cancel="showTimePicker = false"
          @confirm="onTimeConfirm"
        />
      </van-popup>
      <!-- ====== 选择时间 end ====== -->

      <!-- ====== 选择路线 start ====== -->
      <van-field
        readonly
        clickable
        label="选择路线"
        :value="routineValue"
        placeholder="选择路线"
        @click="showRoutinePicker = true"
      />
      <van-popup v-model="showRoutinePicker" round position="bottom">
        <van-picker
          title="选择路线"
          show-toolbar
          visible-item-count="6"
          :columns="routineColumns"
          @cancel="onRoutineCancel"
          @change="onRoutineChange"
          @confirm="onRoutineConfirm"
        />
      </van-popup>
      <!-- ====== 选择路线 end ====== -->

      <!-- ====== 选择城市 start ====== -->
      <van-field
        readonly
        clickable
        label="地点"
        :value="
          position.address == undefined ? '' : position.name + position.address
        "
        placeholder="选择地点"
        @click="showPlacePicker = true"
      />
      <van-popup v-model="showPlacePicker" round position="bottom">
        <div class="amap-page-container">
          <el-amap-search-box
            class="search-box"
            :search-option="searchOption"
            :on-search-result="onSearchResult"
          ></el-amap-search-box>
          <el-amap
            vid="amapDemo"
            :center="mapCenter"
            :zoom="12"
            class="amap-demo"
          >
            <el-amap-marker
              v-for="(marker, index) in markers"
              :position="marker"
              :key="'lalallal' + index"
            ></el-amap-marker>
          </el-amap>
        </div>
        <!-- <template>
             <baidu-map
             @ready="mapReady"
             @click="alert('pc click')"
             class="map"
             style="height: 500px;"
             center="北京"
             >
             <bm-city-list anchor="BMAP_ANCHOR_TOP_LEFT"></bm-city-list>
             </baidu-map>
             </template> -->
        <!-- <van-picker
             show-toolbar
             :columns="placeColumns"
             @cancel="onPlaceCancel"
             @confirm="onPlaceConfirm"
             /> -->
      </van-popup>
      <!-- ====== 选择城市 end ====== -->

      <van-field
        v-model="content"
        name="content"
        label="内容"
        placeholder="请输入内容"
        show-word-limit
        rows="5"
        autosize
        maxlength="500"
        type="textarea"
        :rules="[{ required: true, message: '请输入内容' }]"
      />

      <van-field name="uploader" label="上传照片">
        <template #input>
          <van-uploader
            v-model="fileList"
            capture="camera"
            :afterRead="afterRead"
          />
        </template>
      </van-field>
      <div style="margin: 16px;">
        <van-button round block type="info" native-type="submit">
          提交
        </van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import Vue from "vue";
import Header from "../../components/Header";
import { Form } from "vant";
import { DatetimePicker } from "vant";
import { Picker } from "vant";
import { Popup } from "vant";

// import { BmCityList } from "vue-baidu-map";
import { AMapManager } from "vue-amap";
let amapManager = new AMapManager();
Vue.use(Popup);
import { Uploader } from "vant";
import {
  UPLOAD_A_IMAGE,
  ADD_A_TRAVEL_NOTE,
} from "../../store/mutation-types.js";

Vue.use(Picker);
Vue.use(Uploader);
Vue.use(Form);
Vue.use(DatetimePicker);

export default {
  name: "AddATravelNote",
  components: {
    Header,
    // BmCityList,
  },
  methods: {
    getMapKey() {
      return "map lalal" + this.mapKey++;
    },
    addMarker: function () {
      let lng = 121.5 + Math.round(Math.random() * 1000) / 10000;
      let lat = 31.197646 + Math.round(Math.random() * 500) / 10000;
      this.markers.push([lng, lat]);
    },
    onSearchResult(pois) {
      console.log("position: ", pois);
      this.position = pois[0];

      let latSum = 0;
      let lngSum = 0;
      if (pois.length > 0) {
        pois.forEach((poi) => {
          let { lng, lat } = poi;
          lngSum += lng;
          latSum += lat;
          this.markers.push([poi.lng, poi.lat]);
        });
        let center = {
          lng: lngSum / pois.length,
          lat: latSum / pois.length,
        };
        this.mapCenter = [center.lng, center.lat];
      }
    },
    getMap() {
      // amap vue component
      console.log(amapManager._componentMap);
      // gaode map instance
      console.log(amapManager._map);
    },
    // mapClick(e) {
    //   if (e.overlay && e.overlay.customClickHandler_) {
    //     e.overlay.customClickHandler_.call(e.overlay, e);
    //   }
    // },
    afterRead(file) {
      console.log(file);
      file.status = "uploading";
      file.message = "上传中...";
      this.$store.dispatch(UPLOAD_A_IMAGE, {
        file,
        call() {
          file.status = "done";
          file.message = "上传成功";
        },
      });
    },
    onFormSubmit(values) {
      console.log("submit", values);
      console.log("ids: ", this.$store.state.image.upload_image_ids);
      const that = this;
      this.$store.dispatch(ADD_A_TRAVEL_NOTE, {
        title: values.title,
        content: values.content,
        publish_time: values.time,
        routine_id: this.routine_id,
        latitude: this.position.lat,
        longitude: this.position.lng,
        cover: this.$store.state.image.upload_image_ids
          .toString()
          .split(":")[0],
        image_ids: this.$store.state.image.upload_image_ids,
        call() {
          that.title = "";
          that.content = "";
          that.time = "";
          that.routine_id = 0;
          that.routineValue = "";
          that.image_ids = "";
          that.fileList = [];
          that.position = {};
          that.$toast.success("新增地点成功!");
        },
      });
    },
    onPlaceConfirm(value, index) {
      this.$toast(`当前值：${value}, 当前索引：${index}`);
      this.placeValue = value;
      this.showPlacePicker = false;
    },
    onPlaceChange(picker, value, index) {
      this.$toast(`当前值：${value}, 当前索引：${index}`);
    },
    onPlaceCancel() {
      this.$toast("取消");
      this.showPlacePicker = false;
    },

    onRoutineConfirm(value, index) {
      console.log("routine select value: ", this.routines[index]);
      this.routineValue = this.routines[index].title;
      this.routine_id = this.routines[index].id;
      console.log("idd=============");
      console.log(this.routines[index].id);
      this.showRoutinePicker = false;
    },
    onRoutineChange(picker, value, index) {
      this.$toast(`当前值：${value}, 当前索引：${index}`);
    },
    onRoutineCancel() {
      this.$toast("取消");
      this.showRoutinePicker = false;
    },
    onTimeConfirm(value) {
      this.time = value.toLocaleString();
      this.showTimePicker = false;
      console.log("selet time: ", this.time);
    },
    touchStart(type, target, point, pixel) {
      console.log("type:", type);
      console.log("target:", target);
      console.log("point:", point);
      console.log("pixel:", pixel);
    },
    touchMove(type, target, point, pixel) {
      console.log("type:", type);
      console.log("target:", target);
      console.log("point:", point);
      console.log("pixel:", pixel);
    },
    touchEnd(type, target, point, pixel) {
      console.log("type:", type);
      console.log("target:", target);
      console.log("point:", point);
      console.log("pixel:", pixel);
    },
    dragStart(type, target, point, pixel) {
      console.log("type:", type);
      console.log("target:", target);
      console.log("point:", point);
      console.log("pixel:", pixel);
    },
    syncCenterAndZoom(e) {
      const { lng, lat } = e.target.getCenter();
      this.center.lng = lng;
      this.center.lat = lat;
      this.zoom = e.target.getZoom();
    },
    // mapReady({map }) {
    //   // 解决移动端点击事件不生效问题
    //   let obj = {};
    //   map.addEventListener("touchstart", (e) => {
    //     obj.e = e.changedTouches ? e.changedTouches[0] : e;
    //     obj.target = e.target;
    //     obj.time = Date.now();
    //     obj.X = obj.e.pageX;
    //     obj.Y = obj.e.pageY;
    //   });
    //   map.addEventListener("touchend", (e) => {
    //     obj.e = e.changedTouches ? e.changedTouches[0] : e;
    //     if (
    //       obj.target === e.target &&
    //       // 大于 750 可看成长按了
    //       Date.now() - obj.time < 750 &&
    //       // 应用勾股定理判断，如果 touchstart 的点到 touchend 的点小于 15，就可当成地图被点击了
    //       Math.sqrt(
    //         Math.pow(obj.X - obj.e.pageX, 2) + Math.pow(obj.Y - obj.e.pageY, 2)
    //       ) < 15
    //     ) {
    //       // 这里是地图点击需要执行的操作
    //       alert("移动端点击了地图");
    //     }
    //   });
    // },
  },

  props: {
    routines: Array,
  },
  mounted() {
    this.routineColumns = [];
    for (let i of this.routines) {
      this.routineColumns.push(i.title);
    }
  },
  data() {
    let self = this;
    return {
      mapCenter: [121.59996, 31.197646],
      amapManager,
      mapKey: 1,
      zoom: 12,
      center: [121.59996, 31.197646],
      searchOption: {
        city: "重庆",
        citylimit: false,
      },
      position: {},
      markers: [
        [121.59996, 31.197646],
        [121.40018, 31.197622],
        [121.69991, 31.207649],
      ],
      events: {
        init: (o) => {
          console.log(o.getCenter());
          console.log(this.$refs.map.$$getInstance());
          o.getCity((result) => {
            console.log(result);
          });
        },
        moveend: () => {},
        zoomchange: () => {},
        click: (e) => {
          alert("map clicked:", e);
        },
      },
      loaded: false,
      plugin: [
        "ToolBar",
        {
          pName: "MapType",
          defaultType: 0,
          events: {
            init(o) {
              console.log(o);
            },
          },
        },
        {
          pName: "Geolocation",
          events: {
            init(o) {
              // o 是高德地图定位插件实例
              o.getCurrentPosition((status, result) => {
                if (result && result.position) {
                  self.lng = result.position.lng;
                  self.lat = result.position.lat;
                  self.center = [self.lng, self.lat];
                  self.loaded = true;
                  self.$nextTick();
                }
              });
            },
          },
        },
      ],
      // center: {
      //   lng: 121.480237,
      //   lat: 31.236305,
      //   of: "inner",
      // },
      fileList: [
        // {url: ""}
      ],
      title: "",
      content: "",
      minDate: new Date(1990, 0, 1),
      maxDate: new Date(2090, 10, 1),
      currentDate: new Date(),
      time: "",
      routineColumns: ["杭州", "宁波", "温州", "嘉兴", "湖州", "New Yourk"],
      placeColumns: [
        {
          text: "浙江",
          children: [
            {
              text: "杭州",
              children: [{ text: "西湖区" }, { text: "余杭区" }],
            },
            {
              text: "温州",
              children: [{ text: "鹿城区" }, { text: "瓯海区" }],
            },
          ],
        },
        {
          text: "福建",
          children: [
            {
              text: "福州",
              children: [{ text: "鼓楼区" }, { text: "台江区" }],
            },
            {
              text: "厦门",
              children: [{ text: "思明区" }, { text: "海沧区" }],
            },
          ],
        },
      ],
      showPicker: false,
      showTimePicker: false,
      showPlacePicker: false,
      showRoutinePicker: false,
      showMapPicker: false,
      showAddressBar: true,
      autoLocation: true,
      value: "",
      placeValue: "",
      routineValue: "",
      routine_id: 0,
    };
  },
};
</script>

<style>
.amap-page-container {
  height: 400px;
}
.add_a_travel_note {
  margin-bottom: 70px;
}
.amap-demo {
  height: 300px;
}
.search-box {
  position: absolute;
  top: 25px;
  left: 20px;
  width: 80%;
}
</style>
