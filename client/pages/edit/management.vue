<template>
  <div class="container">
    <div class="card-wrapper" v-for="link in links" :key="link.id">
      <editButton :link="link"/>
    </div>
  </div>
</template>

<script>
import editButton from "~/components/ui/button.vue"

export default {
  components: { editButton },
  data() {
    return {
      links: [
        { title: "作品投稿", href: "/edit/createWork" },
        { title: "記事投稿", href: "/edit/createArticle" },
        { title: "作品管理", href: "/edit/worksList" },
        { title: "記事管理", href: "/edit/articlesList" },
        { title: "ログアウト", href: "/api/logout" },
      ]
    }
  }
}
</script>

<style lang="scss" scoped>

.container {
  display: block;
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  padding-top: 100px;
  min-height: calc(100vh - 50px);
  @include mq {
    min-height: calc(100vh - 80px);
  }
}

.card-wrapper {
  text-align: center;
  margin: 30px auto;
  width: 200px;
}

</style>