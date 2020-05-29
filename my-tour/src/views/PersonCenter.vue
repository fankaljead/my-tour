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

      <div class="person_center_profile">
        <van-form @submit="onSubmit">
          <van-field
            v-model="tel"
            type="tel"
            name="电话"
            label="电话"
            placeholder="电话"
            :rules="[{ required: true, message: '请填写电话' }]"
          />
          <van-field
            v-model="password"
            type="password"
            name="密码"
            label="密码"
            placeholder="密码"
            :rules="[{ required: true, message: '请填写密码' }]"
          />

          <van-field
            readonly
            clickable
            label="选择时间"
            :value="birthday.toLocaleDateString()"
            placeholder="选择时间"
            @click="show = true"
          />
          <van-action-sheet v-model="show">
            <van-datetime-picker
              v-model="time"
              type="date"
              title="选择生日"
              @confirm="confirmBirthday"
              @cancel="cancelBirthday"
              :min-date="minDate"
              :max-date="maxDate"
            />
          </van-action-sheet>

          <van-field
            v-model="email"
            type="text"
            name="邮箱"
            label="邮箱"
            placeholder="邮箱"
            :rules="[{ required: true, message: '请填写邮箱' }]"
          />

          <van-field
            v-model="qq"
            type="text"
            name="QQ"
            label="QQ"
            placeholder="QQ"
            :rules="[{ required: true, message: '请填写QQ' }]"
          />

          <van-field
            v-model="wechat"
            type="text"
            name="微信"
            label="微信"
            placeholder="微信"
            :rules="[{ required: true, message: '请填写微信' }]"
          />

          <van-field
            v-model="personalized_signature"
            type="text"
            name="个性标签"
            label="个性标签"
            placeholder="个性标签"
            :rules="[{ required: true, message: '个性标签' }]"
          />
          <van-field
            v-model="location"
            type="text"
            name="地点"
            label="地点"
            placeholder="地点"
            :rules="[{ required: true, message: '请填写地点' }]"
          />
          <div style="margin: 16px;">
            <van-button round block type="info" native-type="submit">
              提交
            </van-button>
          </div>
        </van-form>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import { NavBar } from "vant";
import { Image as VanImage } from "vant";
import { Form } from "vant";
import { DatetimePicker } from "vant";
import { ActionSheet } from "vant";
import { UPDATE_PERSON_PROFILE } from "../store/mutation-types.js";
import { request } from "../nework/request.js";

Vue.use(ActionSheet);
Vue.use(DatetimePicker);
Vue.use(Form);
Vue.use(VanImage);
Vue.use(NavBar);
export default {
  name: "PersonCenter",
  data() {
    return {
      username: "",
      password: "",
      tel: "",
      qq: "",
      email: "",
      location: "",
      personalized_signature: "",
      wechat: "",
      backgroundImage: "",
      minDate: new Date(1920, 0, 1),
      maxDate: new Date(2025, 10, 1),
      birthday: new Date(),
      show: false,
      time: "",
    };
  },
  mounted() {
    this.getUserInfo();
  },
  methods: {
    onSubmit(values) {
      console.log("submit", values);
      this.$store.dispatch(UPDATE_PERSON_PROFILE, {
        tel: this.tel,
        qq: this.qq,
        email: this.email,
        wechat: this.wechat,
        birthday: this.birthday,
        personalized_signature: this.personalized_signature,
        location: this.location,
        call: this.getUserInfo(),
      });
    },
    getUserInfo() {
      request({
        url: "user_info/" + this.$store.state.user.uid,
        method: "get",
      }).then((res) => {
        console.log("user info res: ", res);
        this.location = res.data.location;
        this.tel = res.data.tel;
        this.wechat = res.data.wechat;
        this.qq = res.data.qq;
        this.email = res.data.email;
        // this.birthday = res.data.birthday;
        this.personalized_signature = res.data.personalized_signature;
      });
    },
    confirmBirthday() {
      this.show = false;
      this.birthday = this.time;
    },
    cancelBirthday() {
      this.show = false;
    },
  },
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
