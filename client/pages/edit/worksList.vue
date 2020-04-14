<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="card-wrapper" v-for="work in works" :key="work.id">
      <workCard :card="work"/>
    </div>
  </div>
</template>

<script>
import workCard from "~/components/ui/editCard.vue"
import tabTitle from "~/components/layouts/tabTitle.vue"

export default {
  async asyncData({ app }) {
    let works = await app.$axios.$get("/api/works")
    works = works.reverse()
    return { works }
  },
  components: {
    workCard,
    tabTitle
  },
  data() {
    return { pageTitle: "作品一覧" }
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

.card-wrapper {
  margin-bottom: 30px;
}

</style>