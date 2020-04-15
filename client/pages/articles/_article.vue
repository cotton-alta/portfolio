<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="article-wrapper">
      <p>{{ Timestamp }}</p>
      <p v-html="Detail"></p>
      <img :src="Image" />
      <twitterButton class="article-twitter" :title="Title"/>
    </div>
    <!-- /.article-wrapper -->
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import twitterButton from "~/components/ui/twitterButton.vue"
import axios from "axios"
import moment from "moment"

export default {
  async asyncData({ app, route }) {
    let urlParameter = encodeURIComponent(route.params.article),
        data = await app.$axios.$get(`/api/articles/${urlParameter}`)
    data[0].Timestamp = moment(data[0].Timestamp).format("LLL")
    return { ...data[0] }
  },
  components: {
    tabTitle,
    twitterButton
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

.article {
  &-wrapper {
    width: 100%;
    line-height: 30px;
    text-align: center;
    margin-bottom: 30px;
    p {
      color: $string-color;
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
  &-twitter {
    margin-top: 30px;
  }
}

</style>