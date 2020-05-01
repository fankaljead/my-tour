<template>
  <div>
    <van-nav-bar title="登录" />
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

      <!-- <div style="margin: 16px;">
           <router-link to="/register">注册</router-link>
           <router-link to="/forget_password">忘记密码</router-link>
           </div>
      -->

      <!-- 两端对齐 -->
      <van-row type="flex" justify="space-around">
        <van-col span="6">
          <router-link to="/register">注册</router-link>
        </van-col>
        <van-col span="6">
          <router-link to="/forget_password">忘记密码</router-link>
        </van-col>
      </van-row>

      <div style="margin: 16px;">
        <van-button round block type="info" native-type="submit">
          登录
        </van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import Vue from "vue";

import { Button, Form, Field } from "vant";
import { NavBar } from "vant";
import { Col, Row } from "vant";

import { LOGIN } from "../store/mutation-types.js";

Vue.use(Col);
Vue.use(Row);
Vue.use(NavBar);
Vue.use(Form);
Vue.use(Field);
Vue.use(Button);

export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: ""
    };
  },
  methods: {
    onSubmit(values) {
      this.$store.dispatch(LOGIN, values);
    }
  }
};
</script>

<style></style>
