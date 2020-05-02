<template>
  <div>
    <router-view />

    <van-tabbar route v-model="active">
      <van-tabbar-item icon="home-o" to="/home">首页</van-tabbar-item>
      <van-tabbar-item icon="like-o" to="/subscribe">关注</van-tabbar-item>
      <van-tabbar-item icon="add-o" to="/add">新增</van-tabbar-item>
      <van-tabbar-item icon="friends-o" to="message">消息</van-tabbar-item>
      <van-tabbar-item icon="setting-o" to="personcenter">个人中心</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
import Vue from "vue";
import { Tabbar, TabbarItem } from "vant";

Vue.use(Tabbar);
Vue.use(TabbarItem);

export default {
  name: "Nav",
  data() {
    return {
      active: 0
    };
  }
};
</script>

<style></style>
