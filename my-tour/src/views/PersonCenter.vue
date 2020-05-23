<template>
  <div>
    <van-nav-bar title="个人中心" />
    <div class="person_center_bar">
      <div class="person_center_bar_head">
        <van-image
          width="100"
          height="100"
          src="https://img.yzcdn.cn/vant/cat.jpeg"
          round
        />
      </div>
      <div class="person_center_bar_titl">
        <h2>{{ $store.state.user.username }}</h2>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import { NavBar } from "vant";
import { Image as VanImage } from "vant";

Vue.use(VanImage);
Vue.use(NavBar);
export default {
  name: "PersonCenter",
};
</script>

<style>
.person_center_bar {
  margin: 20px;
}
.person_center_bar_head {
  width: 30%;
  float: left;
}
.person_center_bar_title {
  width: 30%;
  float: right;
  margin-left: 30px;
}
</style>
