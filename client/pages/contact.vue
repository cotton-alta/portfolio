<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import axios from "axios"

export default {
  async asyncData({ app }) {
    let works = await app.$axios.$put("/api/works")
    return { works }
  },
  components: {
    tabTitle
  },
  data() {
    return { pageTitle: "CONTACT" }
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