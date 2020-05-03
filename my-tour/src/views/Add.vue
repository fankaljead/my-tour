<template>
  <div>
    <AddPopMenu /> 
  </div>
</template>

<script>

import AddPopMenu from '../components/add/AddPopMenu'

export default {
  name: 'Add',
  components: {
    AddPopMenu
  },
  data() {
    return {
      show: true,
    }
  }
}
</script>

<style>

</style>
