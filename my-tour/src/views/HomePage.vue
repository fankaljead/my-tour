<template>
  <div>
    This Home Page
    <TravelNoteBrief v-bind:data="data" />
  </div>
</template>

<script>

import TravelNoteBrief from "../components/homepage/TravelNoteBrief"

export default {
  name: 'HomePage',
  components: {
    TravelNoteBrief
  },
  data() {
    return {
      data: {
        author: "游乐王子",
        title: "摸仙煲一日游",
        intro: "摸仙煲一日游摸仙煲一日游摸仙煲一日游摸仙煲一日游摸仙煲一日游",
        img: 'https://imgsa.baidu.com/forum/w%3D580/sign=b99e7da20355b3199cf9827d73a88286/a6895a540923dd54401396cbd809b3de9d82484e.jpg',
        time: "2020-05-02",
      }
    }
  }
}
</script>

<style>

</style>
