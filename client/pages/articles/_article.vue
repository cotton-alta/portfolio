<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="article-wrapper">
      <p>{{ Detail }}</p>
      <img :src="Image" />
    </div>
    <!-- /.article-wrapper -->
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
    return { ...data[0] }
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
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  min-height: calc(100vh - 50px);
  @include mq {
    min-height: calc(100vh - 80px);
  }
}

.article-wrapper {
  width: 100%;
  line-height: 30px;
  text-align: center;
  p {
    // font-weight: bolder;
    color: $menu-background;
    margin: 0 auto;
    width: 90%;
    text-align: left;
    margin-bottom: 20px;
  }
  img {
    width: 80%;
    border-radius: 3px;
  }
}

</style>