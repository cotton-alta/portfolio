<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
      <div class="article-wrapper" v-for="article in articles" :key="article.id">
      <articleCard :card="article"/>
    </div>
    <!-- /.work-wrapper -->
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import articleCard from "~/components/ui/card.vue"
import axios from "axios"

export default {
  async asyncData({ app }) {
    let articles = await app.$axios.$get("/api/articles")
    if(articles != null) {
      articles = articles.reverse()
    }
    return { articles }
  },
  components: {
    tabTitle,
    articleCard
  },
  mounted() {
    console.log("route", this.$route.path)
  },
  data() {
    return { pageTitle: "BLOG" }
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

.article-wrapper {
  margin: 30px 0px;
}

</style>