<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="work-wrapper" v-for="work in works" :key="work.id">
      <workCard :card="work"/>
    </div>
    <!-- /.work-wrapper -->
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import workCard from "~/components/ui/card.vue"
import axios from "axios"

export default {
  async asyncData({ app }) {
    let works = await app.$axios.$get("/api/works")
    if(works != null) {
      works = works.reverse()
    }
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
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  min-height: calc(100vh - 50px);
  @include mq {
    min-height: calc(100vh - 80px);
  }
}

.work-wrapper {
  margin: 30px 0px;
}

</style>