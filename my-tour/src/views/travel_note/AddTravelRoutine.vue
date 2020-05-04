<template>
  <div class="alltravelroutine">
    <Header title="旅行路线"  />
      <van-form @submit="onSubmit">
        <van-field
          v-model="title"
          name="title"
          label="名称"
          placeholder="旅行路线名称"
          :rules="[{ required: true, message: '请填写旅行路线名称' }]"
        />
        <van-field
          v-model="description"
          name="description"
          label="介绍"
          placeholder="旅行路线介绍"
          :rules="[{ required: true, message: '请填写旅行路线介绍' }]"
        />
        <div style="margin: 16px;">
          <van-button round block type="info" native-type="submit">
            新增路线
          </van-button>
        </div>
      </van-form>
      <AllTravelRoutine />
  </div>
</template>

<script>
import Vue from "vue";
import { Form } from "vant";
import Header from '../../components/Header'
import AllTravelRoutine from '../../components/travel_note/AllTravelRoutine'

Vue.use(Form);

// 新增路线界面
export default {
  name: "AddTravelRoutine",
  components: {
    Header,
    AllTravelRoutine
  },
  data() {
    return {
      show: true,
      title: "",
      description: ""
    };
  },
  methods: {
    onSubmit(values) {
      console.log("submit", values);
    },
    back() {
      this.$router.go(-2);
    }
  }
};
</script>

<style>
 .alltravelroutine {
   margin-bottom: 200px;
 }
</style>
