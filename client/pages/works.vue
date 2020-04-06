<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div v-for="work in works" :key="work.id">
      <workCard :work="work"/>
    </div>
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import workCard from "~/components/ui/card.vue"
import axios from "axios"

export default {
  async asyncData({ app }) {
    let works = JSON.stringify(app.$axios.$get("/api"))
    return { works }
  },
  components: {
    tabTitle,
    workCard
  },
  data() {
    return { pageTitle: "WORKS" }
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