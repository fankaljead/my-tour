<template>
  <div>
    <van-tabs v-model="active" swipeable>
      <van-tab v-for="index in 4" :title="'选项 ' + index">
        内容 {{ index }}
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import Vue from 'vue';
import { Tab, Tabs } from 'vant';

Vue.use(Tab);
Vue.use(Tabs);

export default {
  name: 'ShowAllTravelNotesMap'
}
</script>

<style>

</style>
