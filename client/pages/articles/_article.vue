<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    {{ data }}
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import axios from "axios"

export default {
  async asyncData({ app, route }) {
    let urlParameter = encodeURIComponent(route.params.article),
        data = await app.$axios.$get(`/api/articles/${urlParameter}`)
    return { data }
  },
  components: {
    tabTitle
  },
  data() {
    return { pageTitle: this.$route.params.article }
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