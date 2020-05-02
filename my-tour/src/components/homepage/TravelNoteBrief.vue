<template>
  <div>
    <div class="travel_note_left">
      <img alt="" src="data.img_src" />
    </div>
    <div class="travel_note_right">
      <div>
        <h1>{{ data.title }}</h1>
        <p>
          {{ data.intro }}
        </p>
        <p>
          <span>{{ data.author }}</span>
          <span>
            {{ data.time }}
          </span>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "TravelNoteBrief",
  data(props) {
    return {
      data: props.data
    };
  }
};
</script>

<style>
.travel_note_left {
  width: 20%;
  background-color: #12f32f;
  height: 100px;
  float: left;
}
.travel_note_right {
  width: 80%;
  background-color: #fff321;
  height: 100px;
  float: right;
}
</style>
