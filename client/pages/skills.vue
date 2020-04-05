<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"

export default {
  components: {
    tabTitle
  },
  data() {
    return { pageTitle: "SKILLS" }
  }
}
</script>

<style lang="scss" scoped>

.container {
  display: block;
  min-height: calc(100vh - 80px);
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
}
</style>