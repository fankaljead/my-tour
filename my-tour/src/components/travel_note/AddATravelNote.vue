<template>
  <div>
    <!-- <van-uploader v-model="fileList" multiple /> -->
    <van-uploader v-model="fileList" capture="camera" :afterRead="afterRead" />
  </div>
</template>

<script>
import Vue from "vue";
import { Uploader } from "vant";
// import { fileRequest } from "../../nework/request.js";
// import { request } from "../../nework/request.js";
// import axios from "axios";

import { UPLOAD_A_IMAGE } from "../../store/mutation-types.js";

Vue.use(Uploader);

export default {
  name: "AddATravelNote",
  methods: {
    afterRead(file) {
      console.log(file);
      file.status = "uploading";
      file.message = "上传中...";
      this.$store.dispatch(UPLOAD_A_IMAGE, {
        file,
        call() {
          file.status = 'done';
          file.message = '上传成功';
        }
      });
      // fileRequest(file)
      //   .then((response) => {
      //     return response.text();
      //   })
      //   .then((res) => {
      //     console.log("result:", res);
      //     file.status = 'done';
      //     file.message = '上传成功';
      //   })
      //   .catch((error) => console.log("error", error));
    },
  },
  data() {
    return {
      fileList: [
        // {url: ""}
      ],
    };
  },
};
</script>

<style></style>
