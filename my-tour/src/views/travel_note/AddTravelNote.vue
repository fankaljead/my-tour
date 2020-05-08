<template>
  <div>
    <AddATravelNote v-bind:routines="this.$store.state.travel_note.routines"/>
  </div>
</template>

<script>
import { GET_TRAVEL_ROUTINES } from "../../store/mutation-types";
import AddATravelNote from '../../components/travel_note/AddATravelNote'

export default {
  name: "AddTravelNote",
  components: {
    AddATravelNote
  },

  mounted() {
    this.$store.dispatch(GET_TRAVEL_ROUTINES);
  }
};
</script>

<style></style>
