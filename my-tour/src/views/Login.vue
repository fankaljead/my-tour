<template>
  <div id="login">
    <van-nav-bar title="登录后更精彩" />
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
import { Toast } from "vant";

Vue.use(Toast);

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
      console.log(values);
      const s = this;
      this.$store.dispatch(LOGIN, {
        ...values,
        call() {
          console.log("uid: " + s.$store.state.user.uid);
          if (s.$store.state.user.uid > 0) {
            s.$toast.success("登录成功");
            s.$router.push("/home");
          } else if (s.$store.state.user.uid == 0) {
            s.$toast.fail("用户不存在");
          } else if (s.$store.state.user.uid == -1) {
            s.$toast.fail("密码错误");
          }
        }
      });

      // this.$store.dispatch(LOGIN, values, () => {
      //   console.log("uid: " + this.$store.state.uid);
      //   if (this.$store.state.uid != 0) {
      //     this.$router.replace("/home");
      //   }
      // });
    }
  }
};
</script>

<style></style>
