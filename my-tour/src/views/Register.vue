<template>
  <div>
    <van-nav-bar title="用户注册" left-arrow @click-left="onClickLeft" />
    <van-form @submit="onSubmit">
      <van-field
        v-model="username"
        name="username"
        label="用户名"
        placeholder="用户名"
        :rules="[{ required: true, message: '请填写用户名' }]"
      />
      <van-field
        v-model="password"
        type="password"
        name="password"
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

import { Button, Form, Field } from "vant";
import { NavBar } from "vant";

import { REGISTER } from "../store/mutation-types.js";

Vue.use(NavBar);
Vue.use(Form);
Vue.use(Field);
Vue.use(Button);

export default {
  name: "Register",
  data() {
    return {
      username: "",
      password: ""
    };
  },
  methods: {
    onSubmit(values) {
      console.log(values);
      const s = this;
      this.$store.dispatch(REGISTER, {
        ...values,
        call() {
          console.log("uid: " + s.$store.state.user.uid);
          if (s.$store.state.user.uid > 0) {
            s.$toast.success("注册成功！");
            s.$router.replace("/home");
          }
        }
      });
    },
    onClickLeft() {
      this.$router.replace("/");
      // this.$router.pop()
    }
  }
};
</script>

<style></style>
