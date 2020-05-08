<template>
  <div>
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
        label="城市"
        :value="value"
        placeholder="选择城市"
        @click="showPlacePicker = true"
      />
      <van-popup v-model="showPlacePicker" round position="bottom">
        <van-picker
          show-toolbar
          :columns="placeColumns"
          @cancel="onPlaceCancel"
          @confirm="onPlaceConfirm"
        />
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

Vue.use(Popup);
import { Uploader } from "vant";
import { UPLOAD_A_IMAGE } from "../../store/mutation-types.js";

Vue.use(Picker);
Vue.use(Uploader);
Vue.use(Form);
Vue.use(DatetimePicker);

export default {
  name: "AddATravelNote",
  components: {
    Header,
  },
  methods: {
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
      this.routineValue = value;
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
      this.time = value.toJSON();
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
    return {
      center: {
        lng: 121.480237,
        lat: 31.236305,
        of: "inner",
      },
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
    };
  },
};
</script>

<style>
.amap-page-container {
  height: 400px;
}
</style>
