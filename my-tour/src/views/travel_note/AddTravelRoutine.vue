<template>
  <div>
    <van-form @submit="onSubmit">
      <van-field
        v-model="username"
        name="用户名"
        label="用户名"
        placeholder="用户名"
        :rules="[{ required: true, message: '请填写用户名' }]"
      />
      <van-field
        v-model="password"
        type="password"
        name="密码"
        label="密码"
        placeholder="密码"
        :rules="[{ required: true, message: '请填写密码' }]"
      />
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
import { Form } from "vant";

Vue.use(Form);

// 新增路线界面
export default {
  name: "AddTravelRoutine",
  data() {
    return {
      username: "",
      password: ""
    };
  },
  methods: {
    onSubmit(values) {
      console.log("submit", values);
    }
  }
};
</script>

<style></style>
