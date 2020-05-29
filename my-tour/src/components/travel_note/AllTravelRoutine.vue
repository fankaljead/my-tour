<template>
  <div>
    <h2 class="van-doc-demo-block__title">所有路线</h2>

    <van-steps
      @click-step="checkRoutine"
      direction="vertical"
      :active="0"
      active-icon="location"
    >
      <van-step v-for="routine in routines" :key="routine.id">
        <van-swipe-cell :before-close="beforeClose">
          <div @click="routineClick(routine.id)">
            <!-- <van-swipe-cell > -->
            <h3>{{ routine.title }}</h3>
            <p>{{ routine.description }}</p>
            <p>{{ routine.create_time }}</p>
          </div>
          <template #right>
            <van-button
              square
              text="删除"
              type="danger"
              class="delete-button"
            />
          </template>
        </van-swipe-cell>
      </van-step>
    </van-steps>
  </div>
</template>

<script>
import Vue from "vue";
import { Step, Steps } from "vant";
import { SwipeCell } from "vant";
// import {DELETE_TRAVEL_ROUTINE} from '../../store/mutation-types'
import { Button } from "vant";
import { Dialog } from "vant";

Vue.use(Button);
Vue.use(SwipeCell);
Vue.use(Step);
Vue.use(Steps);

export default {
  name: "AllTravelRoutine",
  data() {
    return {
      beforeCloese: function () {},
    };
  },
  props: {
    routines: Array,
    call: {
      type: Function,
    },
  },
  methods: {
    checkRoutine(index) {
      console.log(this.routines[index]);
    },
    delete: function () {
      console.log("dddd======lllll");
    },
    beforeClose({ position, instance }) {
      switch (position) {
        case "left":
        case "cell":
        case "outside":
          instance.close();
          break;
        case "right":
          Dialog.confirm({
            message: "确定删除吗？",
          })
            .then(() => {
              instance.close();
            })
            .catch(() => {
              instance.close();
            });
          break;
      }
    },
    routineClick(id) {
      console.log("routine id: ", id);
      this.call(id);
    },
  },
};
</script>

<style>
.van-doc-demo-block__title {
  margin: 0;
  padding: 32px 16px 16px;
  color: rgba(69, 90, 100, 0.6);
  font-weight: normal;
  font-size: 24px;
  line-height: 16px;
}
</style>
