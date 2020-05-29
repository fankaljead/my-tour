<template>
  <div>
    <van-nav-bar
      left-text="返回"
      left-arrow
      @click-left="backClick"
      title="用户详情"
    />
    <div class="person_profile_bar">
      <div class="person_center_bar_titl"></div>
      <div class="person_profile_bar_head">
        <van-form>
          <van-field
            v-model="username"
            name="用户名"
            label="用户名"
            placeholder="用户名"
            :rules="[{ required: true, message: '请填写用户名' }]"
          />
        </van-form>

        <van-image
          width="100"
          height="100"
          src="https://img.yzcdn.cn/vant/cat.jpeg"
          round
        />
      </div>
      <van-button
        type="isSubscribe?'primary':'info'"
        :text="isSubscribe ? '订阅' : '取消订阅'"
        @click="subscribeClick"
      />
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import { NavBar } from "vant";
import { Button } from "vant";
import { Form } from "vant";

Vue.use(Form);
Vue.use(Button);
import { request } from "../nework/request.js";

Vue.use(NavBar);
export default {
  name: "PersonProfile",
  data() {
    return {
      username: "",
      isSubscribe: false,
    };
  },
  computed: {
    getUsername: function () {
      return this.username;
    },
  },
  props: {
    user_id: Number,
    data: {},
  },
  mounted() {
    this.getUserData();
  },
  methods: {
    getUserData() {
      request({
        url: "user_info/" + this.$route.params.user_id,
        method: "get",
      }).then((res) => {
        console.log("user info profile: ", res.data);
        this.data = res.data;
        request({
          url: "user/" + this.data.user_id,
          method: "get",
        }).then((ress) => {
          console.log("user pass info: ", ress.data);
          this.username = ress.data.username;
        });
      });
    },
    backClick() {
      this.$router.go(-1);
    },
    subscribeClick() {
      this.isSubscribe = !this.isSubscribe;
      if (this.isSubscribe) {
        request({
          url: "fan/followUser",
          method: "post",
          params: {
            follower_id: this.user_id,
          },
        }).then((res) => {
          if (res.data > 0) {
            this.$toast("关注成功");
          }
        });
      } else {
        request({
          url: "fan/unfollowUser",
          method: "post",
          params: {
            follower_id: this.user_id,
          },
        }).then((res) => {
          if (res.data > 0) {
            this.$toast("取消关注成功");
          }
        });
      }
    },
  },
};
</script>

<style></style>
