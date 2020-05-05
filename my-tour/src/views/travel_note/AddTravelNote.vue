<template>
  <div>
    This Add Travel Note
    <AddATravelNote />
    <AllTravelRoutine v-bind:routines="this.$store.state.travel_note.routines" />
  </div>
</template>

<script>
import AllTravelRoutine from "../../components/travel_note/AllTravelRoutine";
import { GET_TRAVEL_ROUTINES } from "../../store/mutation-types";
import AddATravelNote from '../../components/travel_note/AddATravelNote'

export default {
  name: "AddTravelNote",
  components: {
    AllTravelRoutine,
    AddATravelNote
  },

  mounted() {
    this.$store.dispatch(GET_TRAVEL_ROUTINES);
  }
};
</script>

<style></style>
