<template>
  <div>
    <ShowATravelNote :images="images" :data="data" />
  </div>
</template>

<script>
import ShowATravelNote from "../../components/travel_note/ShowATravelNote";
import { request } from "../../nework/request.js";

export default {
  name: "CheckATravelNote",
  components: {
    ShowATravelNote,
  },
  data() {
    return {
      travel_note_id: "",
      image_ids: "",
      images: [],
      data: {},
    };
  },
  mounted() {
    this.travel_note_id = this.$route.params.travel_note_id;
    request({
      url: "travel_note/" + this.travel_note_id,
      method: "get",
    })
      .then((res) => {
        console.log("checkATravelNote===");
        console.log(res);
        this.image_ids = res.data.image_ids;
        this.data = res.data;
      })
      .then(() => {
        request({
          url: "image/getImagesByIds",
          method: "get",
          params: {
            image_ids: this.image_ids,
          },
        }).then((res) => {
          console.log("checkATravelNote images ===");
          console.log(res);
          this.images = res.data;
        });
      });
  },
  methods: {},
};
</script>

<style></style>
