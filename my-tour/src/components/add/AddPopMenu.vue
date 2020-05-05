<template>
  <div>
    <van-action-sheet
      cancel-text="取消"
      @cancel="onCancel"
      v-model="show"
      :actions="actions"
      @select="onSelect"
    />
  </div>
</template>

<script>
import Vue from "vue";
import { ActionSheet } from "vant";
import { Toast } from "vant";

Vue.use(ActionSheet);

// 弹出新增的菜单
export default {
  name: "AddPopMenu",
  data() {
    return {
      show: true,
      actions: [{ name: "新增地点" }, { name: "新增路线" }]
    };
  },

  methods: {
    onSelect(item, index) {
      // 默认情况下点击选项时不会自动收起
      // 可以通过 close-on-click-action 属性开启自动收起
      this.show = false;
      console.log(index);
      switch (index) {
        // 新增地点
        case 0:
          this.$router.replace("/addTravelNote");
          break;
        // 新增路线
        case 1:
          this.$router.replace("/addTravelRoutine")
          break;
        default:
      }
      Toast(item.name);
    },
    onCancel() {
      // 后退一步记录，等同于 history.back()
      this.$router.go(-1);
    }
  }
};
</script>

<style>
.content {
  padding: 16px 16px 160px;
}
</style>
