<template>
  <div class="homepage">
    <Header />
  </div>
</template>

<script>
import Header from "../components/homepage/Header";

export default {
  name: "HomePage",
  components: {
    Header,
  },
  data() {
    return {
      data: {
        author: "游乐王子",
        title: "摸仙煲一日游",
        intro: "摸仙煲一日游摸仙煲一日游摸仙煲一日游摸仙煲一日游摸仙煲一日游",
        img:
          "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.5rYEZ-xCStnRSPYhA3_i_QAAAA%26pid%3DApi&f=1",
        time: "2020-05-02",
      },
    };
  },
};
</script>

<style>
.homepage {
  margin-bottom: 70px;
}
</style>
